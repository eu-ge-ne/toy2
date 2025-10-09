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
	start, oldEnd, ok := s.buffer.Index2(startLn, startCol, endLn, endCol)
	if !ok {
		panic("in Syntax.Delete")
	}

	s.edit(start, oldEnd, start+1, startLn, startCol, endLn, endCol, startLn, startCol+1)

	s.parse()
}

func (s *Syntax) Insert(startLn, startCol, endLn, endCol int) {
	start, newEnd, ok := s.buffer.Index2(startLn, startCol, endLn, endCol)
	if !ok {
		panic("in Syntax.Insert")
	}

	s.edit(start, start+1, newEnd, startLn, startCol, startLn, startCol+1, endLn, endCol)

	s.parse()
}

func (s *Syntax) edit(start, oldEnd, newEnd, startLn, startCol, oldEndLn, oldEndCol, newEndLn, newEndCol int) {
	s.tree.Edit(&treeSitter.InputEdit{
		StartByte:      uint(start),
		OldEndByte:     uint(oldEnd),
		NewEndByte:     uint(newEnd),
		StartPosition:  treeSitter.NewPoint(uint(startLn), uint(startCol)),
		OldEndPosition: treeSitter.NewPoint(uint(oldEndLn), uint(oldEndCol)),
		NewEndPosition: treeSitter.NewPoint(uint(newEndLn), uint(newEndCol)),
	})
}

func (s *Syntax) parse() {
	s.tree = s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		return []byte(s.buffer.Chunk(i))
	}, s.tree, nil)
}
