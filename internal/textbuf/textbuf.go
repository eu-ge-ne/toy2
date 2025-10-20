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

func (buf *TextBuf) LineCount() int {
	if buf.Count() == 0 {
		return 0
	}
	return buf.tree.Root.TotalEolsLen + 1
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
	buf.Delete(0, math.MaxInt)

	if len(text) > 0 {
		buf.Insert(0, text)
	}
}

func (buf *TextBuf) Validate() {
	buf.tree.Root.Validate()
}
