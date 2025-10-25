package syntax

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/grammar"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type Syntax struct {
	buffer *textbuf.TextBuf

	grammar grammar.Grammar
	parser  *treeSitter.Parser
	tree    *treeSitter.Tree

	spans chan span
	span  span
	idx   int

	log     *os.File
	text    []byte
	started time.Time
}

type span struct {
	startIdx int
	endIdx   int
	name     string
}

func New(buffer *textbuf.TextBuf) *Syntax {
	s := Syntax{
		buffer: buffer,
	}

	s.initLogger()

	return &s
}

func (s *Syntax) SetLanguage(grm grammar.Grammar) {
	if s.tree != nil {
		s.tree.Close()
	}

	if s.parser != nil {
		s.parser.Close()
	}

	s.grammar = grm
	if grm == nil {
		return
	}

	s.parser = treeSitter.NewParser()

	err := s.parser.SetLanguage(grm.Lang())
	if err != nil {
		panic(err)
	}
}

func (s *Syntax) Delete(change textbuf.Change) {
	if s.tree == nil {
		return
	}

	e := treeSitter.InputEdit{
		StartByte:  uint(change.Start.Idx),
		OldEndByte: uint(change.End.Idx),

		StartPosition:  treeSitter.NewPoint(uint(change.Start.Ln), uint(change.Start.ColIdx)),
		OldEndPosition: treeSitter.NewPoint(uint(change.End.Ln), uint(change.End.ColIdx)),
	}

	e.NewEndByte = e.StartByte
	e.NewEndPosition = e.StartPosition

	s.tree.Edit(&e)

	fmt.Fprintf(s.log, "delete: change %+v\n", change)
	fmt.Fprintf(s.log, "delete: e %+v\n", e)
}

func (s *Syntax) Insert(change textbuf.Change) {
	if s.tree == nil {
		return
	}

	e := treeSitter.InputEdit{
		StartByte:  uint(change.Start.Idx),
		NewEndByte: uint(change.End.Idx),

		StartPosition:  treeSitter.NewPoint(uint(change.Start.Ln), uint(change.Start.ColIdx)),
		NewEndPosition: treeSitter.NewPoint(uint(change.End.Ln), uint(change.End.ColIdx)),
	}

	e.OldEndByte = e.StartByte
	e.OldEndPosition = e.StartPosition

	s.tree.Edit(&e)

	fmt.Fprintf(s.log, "insert: change %+v\n", change)
	fmt.Fprintf(s.log, "insert: e %+v\n", e)
}

func (s *Syntax) Highlight(startLn, endLn int) {
	if s.grammar == nil {
		return
	}

	s.started = time.Now()

	fmt.Fprintln(s.log, "highlight: started")

	startPos, _ := s.buffer.Pos(startLn, 0)
	endPos := s.buffer.EndPos(endLn, 0)

	s.spans = make(chan span, 2024)
	s.span = span{startIdx: -1, endIdx: -1}
	s.idx = startPos.Idx

	go s.highlight(startPos, endPos)
}

func (s *Syntax) NextSpan(l int) string {
	if s.grammar == nil {
		return ""
	}

	var name string

	if s.idx >= s.span.endIdx {
		if spn, ok := <-s.spans; ok {
			s.span = spn
		}
	}

	if s.idx >= s.span.startIdx && s.idx < s.span.endIdx {
		name = s.span.name
	}

	s.idx += l

	return name
}

func (s *Syntax) highlight(startPos textbuf.Pos, endPos textbuf.Pos) {
	query := s.grammar.Query()

	s.parse(startPos, endPos)
	s.prepareText(startPos, endPos)

	qc := treeSitter.NewQueryCursor()
	qc.SetByteRange(uint(startPos.Idx), uint(endPos.Idx))
	defer qc.Close()

	var spn span

	capts := qc.Captures(query, s.tree.RootNode(), s.text)

	match, captIdx := capts.Next()
	if match != nil {
		capt := match.Captures[captIdx]
		spn = span{int(capt.Node.StartByte()), int(capt.Node.EndByte()), query.CaptureNames()[capt.Index]}
	}

	for ; match != nil; match, captIdx = capts.Next() {
		capt := match.Captures[captIdx]
		name := query.CaptureNames()[capt.Index]
		startIdx := int(capt.Node.StartByte())
		endIdx := int(capt.Node.EndByte())

		//fmt.Fprintf(s.log, "highlight: %v:%v %s (%s)\n", capt.Node.StartPosition(), capt.Node.EndPosition(), capt.Node.Utf8Text(s.text), name /*match.PatternIndex,*/ /*capt.Index,*/)

		if spn.startIdx != startIdx || spn.endIdx != endIdx {
			s.spans <- spn
			spn = span{startIdx, endIdx, name}
		} else {
			spn.name = name
		}
	}

	s.spans <- spn

	close(s.spans)

	fmt.Fprintf(s.log, "highlight: [%v] completed\n", time.Since(s.started))
}

const maxChunkLen = 1024 * 4

func (s *Syntax) parse(start, endPos textbuf.Pos) {
	startPos, _ := s.buffer.Pos(max(0, start.Ln-1_000), 0)

	s.parser.SetIncludedRanges([]treeSitter.Range{{
		StartByte:  uint(startPos.Idx),
		EndByte:    uint(endPos.Idx),
		StartPoint: treeSitter.NewPoint(uint(startPos.Ln), uint(startPos.ColIdx)),
		EndPoint:   treeSitter.NewPoint(uint(endPos.Ln), uint(endPos.ColIdx)),
	}})

	oldTree := s.tree

	s.tree = s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		text := s.buffer.Chunk(i)

		if len(text) > maxChunkLen {
			text = text[0:maxChunkLen]
		}

		fmt.Fprintf(s.log, "parse: [%v] reading chunk %d, %+v, %d\n", time.Since(s.started), i, p, len(text))

		return []byte(text)
	}, oldTree, nil)

	fmt.Fprintf(s.log, "parse: [%v] completed\n", time.Since(s.started))
}

func (s *Syntax) prepareText(startPos, endPos textbuf.Pos) {
	if s.buffer.Count() > len(s.text) {
		s.text = make([]byte, s.buffer.Count())
	}

	copy(
		s.text[startPos.Idx:endPos.Idx],
		std.IterToStr(s.buffer.Slice(startPos.Idx, endPos.Idx)),
	)
}

func (s *Syntax) initLogger() {
	log, err := os.OpenFile("tmp/syntax.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	/*
		i := 0

		s.parser.SetLogger(func(t treeSitter.LogType, msg string) {
			var tp string

			switch t {
			case treeSitter.LogTypeParse:
				tp = "parse"
			case treeSitter.LogTypeLex:
				tp = "lex"
			}

			fmt.Fprintf(log, "%d: %s: %s\n", i, tp, msg)

			i += 1
		})
	*/

	s.log = log
}
