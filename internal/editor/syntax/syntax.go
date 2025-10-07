package syntax

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

type Syntax struct {
	buffer *textbuf.TextBuf
	parser *treeSitter.Parser
	tree   *treeSitter.Tree
}

func New(buffer *textbuf.TextBuf) Syntax {
	return Syntax{
		buffer: buffer,
		parser: treeSitter.NewParser(),
	}
}

func (s *Syntax) SetLanguage() {
	err := s.parser.SetLanguage(treeSitter.NewLanguage(treeSitterTs.LanguageTypescript()))
	if err != nil {
		panic(err)
	}

	s.parse()
}

func (s *Syntax) Delete(startLn, startCol, endLn, endCol int) {
	startByte, ok := s.buffer.Index(startLn, startCol)
	if !ok {
		panic("in Syntax.Delete")
	}

	oldEndByte, ok := s.buffer.Index(endLn, endCol)
	if !ok {
		panic("in Syntax.Delete")
	}

	s.tree.Edit(&treeSitter.InputEdit{
		StartByte:      uint(startByte),
		OldEndByte:     uint(oldEndByte),
		NewEndByte:     uint(startByte + 1),
		StartPosition:  treeSitter.NewPoint(uint(startLn), uint(startCol)),
		OldEndPosition: treeSitter.NewPoint(uint(endLn), uint(endCol)),
		NewEndPosition: treeSitter.NewPoint(uint(startLn), uint(startCol+1)),
	})

	s.parse()
}

func (s *Syntax) Insert(startLn, startCol int, text string) {
	s.parse()
}

func (s *Syntax) parse() {
	s.tree = s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		return []byte(s.buffer.Chunk(i))
	}, s.tree, nil)
}
