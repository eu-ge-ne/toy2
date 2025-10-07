package syntax

import (
	"math"

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

func (s *Syntax) Reset() {
	s.parser.SetLanguage(treeSitter.NewLanguage(treeSitterTs.LanguageTypescript()))

	s.parse()
}

func (s *Syntax) Delete(startLn, startCol, endLn, endCol int) {
	/*
		s.tree.Edit(&treeSitter.InputEdit{
			StartByte:      startByte,
			OldEndByte:     oldEndByte,
			NewEndByte:     newEndByte,
			StartPosition:  treeSitter.NewPoint(startPosLn, startPosCol),
			OldEndPosition: treeSitter.NewPoint(oldEndPosLn, oldEndPosCol),
			NewEndPosition: treeSitter.NewPoint(newEndPosLn, newEndPosCol),
		})
	*/

	s.parse()
}

func (s *Syntax) Insert(startLn, startCol int, text string) {
	s.parse()
}

func (s *Syntax) parse() {
	//s.tree = s.parser.Parse([]byte(s.buffer.Text()), nil)

	s.tree = s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		for text := range s.buffer.ReadSlice(i, math.MaxInt) {
			return []byte(text)
		}
		return make([]byte, 0)
	}, s.tree, nil)
}
