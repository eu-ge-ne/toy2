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

	query         *treeSitter.Query
	startPos      textbuf.Pos
	endPos        textbuf.Pos
	startPosParse textbuf.Pos
	spans         chan span
	span          span
	idx           int

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

	f, err := os.OpenFile("tmp/syntax.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	s.log = f

	return &s
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

	var e treeSitter.InputEdit

	e.StartByte = uint(change.Start.Idx)
	e.OldEndByte = uint(change.End.Idx)

	e.StartPosition.Row = uint(change.Start.Ln)
	e.StartPosition.Column = uint(change.Start.ColIdx)
	e.OldEndPosition.Row = uint(change.End.Ln)
	e.OldEndPosition.Column = uint(change.End.ColIdx)

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

	var e treeSitter.InputEdit

	e.StartByte = uint(change.Start.Idx)
	e.NewEndByte = uint(change.End.Idx)

	e.StartPosition.Row = uint(change.Start.Ln)
	e.StartPosition.Column = uint(change.Start.ColIdx)
	e.NewEndPosition.Row = uint(change.End.Ln)
	e.NewEndPosition.Column = uint(change.End.ColIdx)

	e.OldEndByte = e.StartByte
	e.OldEndPosition = e.StartPosition

	s.tree.Edit(&e)

	fmt.Fprintf(s.log, "insert: change %+v\n", change)
	fmt.Fprintf(s.log, "insert: e %+v\n", e)
}

func (s *Syntax) Highlight(startLn, endLn int) {
	s.startPos, _ = s.buffer.Pos(startLn, 0)
	s.endPos = s.buffer.EndPos(endLn, 0)
	s.startPosParse, _ = s.buffer.Pos(max(0, startLn-2_000), 0)
	s.spans = make(chan span, 1024)
	s.span = span{-1, -1, ""}
	s.idx = s.startPos.Idx

	go s.highlight()
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

func (s *Syntax) highlight() {
	started := time.Now()

	fmt.Fprintln(s.log, "highlight: started")

	s.parse(s.startPosParse, s.endPos)

	fmt.Fprintf(s.log, "highlight: parsed %v\n", time.Since(started))

	if s.buffer.Count() > len(s.text) {
		s.text = make([]byte, s.buffer.Count())
	}

	copy(
		s.text[s.startPos.Idx:s.endPos.Idx],
		std.IterToStr(s.buffer.Slice(s.startPos.Idx, s.endPos.Idx)),
	)

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()

	qc.SetByteRange(uint(s.startPos.Idx), uint(s.endPos.Idx))

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

func (s *Syntax) Next(l int) string {
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
