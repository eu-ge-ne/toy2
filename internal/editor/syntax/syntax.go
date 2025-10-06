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
	return Syntax{buffer: buffer}
}

func (s *Syntax) Reset() {
	s.parser = treeSitter.NewParser()

	s.parser.SetLanguage(treeSitter.NewLanguage(treeSitterTs.LanguageTypescript()))

	s.tree = s.parser.Parse([]byte(s.buffer.Text()), nil)
}

func (s *Syntax) Edit(startLn, startCol, endLn, endCol int) {
	/*
		s.tree.Edit(&treeSitter.InputEdit{
			StartByte:      startByte,
			OldEndByte:     oldEndByte,
			NewEndByte:     newEndByte,
			StartPosition:  treeSitter.NewPoint(startPosLn, startPosCol),
			OldEndPosition: treeSitter.NewPoint(oldEndPosLn, oldEndPosCol),
			NewEndPosition: treeSitter.NewPoint(newEndPosLn, newEndPosCol),
		})

		s.tree = s.parser.Parse([]byte(s.buffer.Text()), s.tree)
	*/
}
