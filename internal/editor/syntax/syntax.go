package syntax

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	_ "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

//go:embed js/highlights.scm
var scmJsHighlights string

//go:embed ts/highlights.scm
var scmTsHighlights string

type Syntax struct {
	buffer *textbuf.TextBuf
	parser *treeSitter.Parser
	tree   *treeSitter.Tree

	query *treeSitter.Query
	spans chan span
	span  span
	idx   int

	log  *os.File
	text []byte
}

type span struct {
	startIdx int
	endIdx   int
	name     string
}

func New(buffer *textbuf.TextBuf) *Syntax {
	s := Syntax{
		buffer: buffer,
		parser: treeSitter.NewParser(),
	}

	//Log(s.parser)

	f, err := os.OpenFile("tmp/syntax.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	s.log = f

	return &s
}

func (s *Syntax) SetLanguage() {
	lang := treeSitter.NewLanguage(treeSitterTs.LanguageTypescript())

	err := s.parser.SetLanguage(lang)
	if err != nil {
		panic(err)
	}

	query, err0 := treeSitter.NewQuery(lang, scmJsHighlights+scmTsHighlights)
	if err0 != nil {
		panic(err0)
	}
	s.query = query
}

func (s *Syntax) Close() {
	if s.log != nil {
		s.log.Close()
		s.log = nil
	}

	s.tree.Close()
	s.tree = nil
}

func (s *Syntax) Delete(change textbuf.Change) {
	if s == nil || s.tree == nil {
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
	if s == nil || s.tree == nil {
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
	startPos, _ := s.buffer.Pos(startLn, 0)
	endPos := s.buffer.EndPos(endLn, 0)
	startPosParse, _ := s.buffer.Pos(max(0, startLn-2_000), 0)

	s.spans = make(chan span, 1024)
	s.span = span{-1, -1, ""}
	s.idx = startPos.Idx

	go s.highlight(startPos, endPos, startPosParse)
}

func (s *Syntax) NextSpan(l int) string {
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

func (s *Syntax) highlight(startPos textbuf.Pos, endPos textbuf.Pos, startPosParse textbuf.Pos) {
	started := time.Now()

	fmt.Fprintln(s.log, "highlight: started")

	s.parse(startPosParse, endPos)

	fmt.Fprintf(s.log, "highlight: parsed %v\n", time.Since(started))

	if s.buffer.Count() > len(s.text) {
		s.text = make([]byte, s.buffer.Count())
	}

	copy(
		s.text[startPos.Idx:endPos.Idx],
		std.IterToStr(s.buffer.Slice(startPos.Idx, endPos.Idx)),
	)

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()

	qc.SetByteRange(uint(startPos.Idx), uint(endPos.Idx))

	var spn span

	capts := qc.Captures(s.query, s.tree.RootNode(), s.text)

	match, captIdx := capts.Next()
	if match != nil {
		capt := match.Captures[captIdx]
		spn = span{
			int(capt.Node.StartByte()),
			int(capt.Node.EndByte()),
			s.query.CaptureNames()[capt.Index],
		}
	}

	for ; match != nil; match, captIdx = capts.Next() {
		capt := match.Captures[captIdx]
		name := s.query.CaptureNames()[capt.Index]

		/*
			fmt.Fprintf(s.log,
				"highlight: %v:%v %s (%s)\n",
				capt.Node.StartPosition(),
				capt.Node.EndPosition(),
				capt.Node.Utf8Text(s.text),
				name,
				//match.PatternIndex,
				//capt.Index,
			)
		*/

		startIdx := int(capt.Node.StartByte())
		endIdx := int(capt.Node.EndByte())

		if spn.startIdx != startIdx || spn.endIdx != endIdx {
			s.spans <- spn
			spn = span{startIdx, endIdx, name}
		} else {
			spn.name = name
		}
	}

	s.spans <- spn

	close(s.spans)

	fmt.Fprintf(s.log, "highlight: elapsed %v\n", time.Since(started))
}

const maxChunkLen = 1024 * 4

func (s *Syntax) parse(startPos, endPos textbuf.Pos) {
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

		fmt.Fprintf(s.log, "parse: reading chunk %d, %+v, %d\n", i, p, len(text))

		return []byte(text)
	}, oldTree, nil)
}
