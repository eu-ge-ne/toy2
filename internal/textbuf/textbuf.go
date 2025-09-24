package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf/internal/content"
	"github.com/eu-ge-ne/toy2/internal/textbuf/internal/node"
	"github.com/eu-ge-ne/toy2/internal/textbuf/internal/tree"
)

type TextBuf struct {
	content content.Content
	tree    tree.Tree
}

func New(text string) *TextBuf {
	buf := TextBuf{
		content: content.Content{},
		tree:    tree.Tree{Root: node.NIL},
	}

	if len(text) > 0 {
		buf.tree.Root = buf.content.Create(text)
		buf.tree.Root.Red = false
	}

	return &buf
}

func (tb *TextBuf) Count() int {
	return tb.tree.Root.TotalLen
}

func (tb *TextBuf) LineCount() int {
	if tb.tree.Root.TotalLen == 0 {
		return 0
	} else {
		return tb.tree.Root.TotalEolsLen + 1
	}
}

func (tb *TextBuf) Save() *node.Node {
	return tb.tree.Root.Clone(node.NIL)
}

func (tb *TextBuf) Restore(n *node.Node) {
	tb.tree.Root = n.Clone(node.NIL)
}

func (tb *TextBuf) Reset(text string) {
	tb.DeleteToEnd(0)

	if len(text) > 0 {
		tb.Insert(0, text)
	}
}

func (tb *TextBuf) Validate() {
	tb.tree.Root.Validate()
}
