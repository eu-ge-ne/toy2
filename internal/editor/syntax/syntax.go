package syntax

import (
	"log"
	"time"

	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/grammar"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type Syntax struct {
	parser *treeSitter.Parser
	tree   *treeSitter.Tree
	query  *treeSitter.Query

	grammar grammar.Grammar

	text    []byte
	started time.Time

	spans chan span
	span  span
	idx   int
}

type span struct {
	startIdx int
	endIdx   int
	name     string
}

func New() *Syntax {
	s := Syntax{}

	//s.parser.SetLogger(func(t treeSitter.LogType, msg string) { log.Printf("%v: %s", t, msg) })

	return &s
}

func (s *Syntax) SetGrammar(grm grammar.Grammar) {
	s.grammar = grm

	s.close()

	if grm != nil {
		s.init(grm)
	}
}

func (s *Syntax) init(grm grammar.Grammar) {
	s.parser = treeSitter.NewParser()

	err := s.parser.SetLanguage(grm.Lang())
	if err != nil {
		panic(err)
	}

	s.query = grm.Query()
}

func (s *Syntax) close() {
	if s.query != nil {
		s.query.Close()
		s.query = nil
	}

	if s.tree != nil {
		s.tree.Close()
		s.tree = nil
	}

	if s.parser != nil {
		s.parser.Close()
		s.parser = nil
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

	log.Printf("delete: %+v; %+v", change, e)
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

	log.Printf("insert: %+v; %+v", change, e)
}

func (s *Syntax) Highlight(buf *textbuf.TextBuf, startLn, endLn int) {
	if s.grammar == nil {
		return
	}

	s.started = time.Now()

	startPos, _ := buf.Pos(startLn, 0)
	endPos := buf.EndPos(endLn, 0)

	log.Printf("[%v] highlighting: %v:%v", time.Since(s.started), startPos, endPos)

	s.spans = make(chan span, 1024*2)
	s.span = span{startIdx: -1, endIdx: -1}
	s.idx = startPos.Idx

	go func() {
		s.parse(buf, startPos, endPos)
		s.prepareText(buf, startPos, endPos)

		qc := treeSitter.NewQueryCursor()
		qc.SetByteRange(uint(startPos.Idx), uint(endPos.Idx))
		defer qc.Close()

		var spn span

		capts := qc.Captures(s.query, s.tree.RootNode(), s.text)

		match, captIdx := capts.Next()
		if match != nil {
			capt := match.Captures[captIdx]
			spn = span{int(capt.Node.StartByte()), int(capt.Node.EndByte()), s.query.CaptureNames()[capt.Index]}
		}

		for ; match != nil; match, captIdx = capts.Next() {
			capt := match.Captures[captIdx]
			name := s.query.CaptureNames()[capt.Index]
			startIdx := int(capt.Node.StartByte())
			endIdx := int(capt.Node.EndByte())

			log.Printf("%v:%v %s (%s)\n", capt.Node.StartPosition(), capt.Node.EndPosition(), capt.Node.Utf8Text(s.text), name /*match.PatternIndex,*/ /*capt.Index,*/)

			if spn.startIdx != startIdx || spn.endIdx != endIdx {
				s.spans <- spn
				spn = span{startIdx, endIdx, name}
			} else {
				spn.name = name
			}
		}

		s.spans <- spn

		close(s.spans)

		log.Printf("[%v] highlighting completed", time.Since(s.started))
	}()
}

func (s *Syntax) Next(l int) string {
	if s.grammar == nil {
		return ""
	}

	defer func() { s.idx += l }()

	if s.idx >= s.span.endIdx {
		if spn, ok := <-s.spans; ok {
			s.span = spn
		}
	}

	if s.idx >= s.span.startIdx && s.idx < s.span.endIdx {
		return s.span.name
	}

	return ""
}

const maxChunkLen = 1024 * 4

func (s *Syntax) parse(buf *textbuf.TextBuf, start, endPos textbuf.Pos) {
	log.Printf("[%v] parsing", time.Since(s.started))

	startPos, _ := buf.Pos(max(0, start.Ln-2_000), 0)

	s.parser.SetIncludedRanges([]treeSitter.Range{{
		StartByte:  uint(startPos.Idx),
		EndByte:    uint(endPos.Idx),
		StartPoint: treeSitter.NewPoint(uint(startPos.Ln), uint(startPos.ColIdx)),
		EndPoint:   treeSitter.NewPoint(uint(endPos.Ln), uint(endPos.ColIdx)),
	}})

	oldTree := s.tree

	s.tree = s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		text := buf.Chunk(i)

		if len(text) > maxChunkLen {
			text = text[0:maxChunkLen]
		}

		log.Printf("[%v] reading chunk: %v;%+v;%v", time.Since(s.started), i, p, len(text))

		return []byte(text)
	}, oldTree, nil)

	log.Printf("[%v] parsing completed", time.Since(s.started))
}

func (s *Syntax) prepareText(buf *textbuf.TextBuf, startPos, endPos textbuf.Pos) {
	if buf.Count() > len(s.text) {
		s.text = make([]byte, buf.Count())
	}

	copy(
		s.text[startPos.Idx:endPos.Idx],
		std.IterToStr(buf.Slice(startPos.Idx, endPos.Idx)),
	)
}
