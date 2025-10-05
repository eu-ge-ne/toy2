package syntax

import (
	"github.com/eu-ge-ne/toy2/internal/segbuf"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

type Syntax struct {
	buffer *segbuf.SegBuf
}

func New(buffer *segbuf.SegBuf) Syntax {
	return Syntax{buffer: buffer}
}

func (s *Syntax) Reset() {
	parser := treeSitter.NewParser()
	defer parser.Close()

	parser.SetLanguage(treeSitter.NewLanguage(treeSitterTs.LanguageTypescript()))

	tree := parser.Parse([]byte(s.buffer.Text()), nil)
	defer tree.Close()
}
