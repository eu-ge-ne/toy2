package textbuf

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/textbuf/content"
	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
	"github.com/eu-ge-ne/toy2/internal/textbuf/tree"
)

type TextBuf struct {
	content content.Content
	tree    tree.Tree
}

type Snapshot struct {
	node *node.Node
}

func New() *TextBuf {
	return &TextBuf{
		content: content.Content{},
		tree:    tree.Tree{Root: node.NIL},
	}
}

func (buf *TextBuf) Count() int {
	return buf.tree.Root.TotalLen
}

func (buf *TextBuf) Save() Snapshot {
	return Snapshot{
		node: buf.tree.Root.Clone(node.NIL),
	}
}

func (buf *TextBuf) Restore(s Snapshot) {
	buf.tree.Root = s.node.Clone(node.NIL)
}

func (buf *TextBuf) Reset(text string) {
	buf.delete(0, math.MaxInt)

	if len(text) > 0 {
		buf.insert(0, text)
	}
}

func (buf *TextBuf) Validate() {
	buf.tree.Root.Validate()
}
