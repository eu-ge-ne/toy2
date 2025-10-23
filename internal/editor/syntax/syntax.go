package syntax

import (
	_ "embed"
	"fmt"
	"os"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	_ "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

//go:embed js/highlights.scm
var scmJsHighlights string

//go:embed ts/highlights.scm
var scmTsHighlights string

type Syntax struct {
	buffer *textbuf.TextBuf
	parser *treeSitter.Parser
	query  *treeSitter.Query
	tree   *treeSitter.Tree

	log *os.File
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

func (s *Syntax) Highlight(startLn, endLn int) *Highlighter {
	return NewHighlighter(s, startLn, endLn)
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
