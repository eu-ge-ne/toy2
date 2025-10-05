package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/segbuf/textbuf/content"
	"github.com/eu-ge-ne/toy2/internal/segbuf/textbuf/node"
	"github.com/eu-ge-ne/toy2/internal/segbuf/textbuf/tree"
)

type TextBuf struct {
	content content.Content
	tree    tree.Tree
}

type Snapshot struct {
	node *node.Node
}

func New() TextBuf {
	return TextBuf{
		content: content.Content{},
		tree:    tree.Tree{Root: node.NIL},
	}
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

func (tb *TextBuf) Save() Snapshot {
	return Snapshot{
		node: tb.tree.Root.Clone(node.NIL),
	}
}

func (tb *TextBuf) Restore(s Snapshot) {
	tb.tree.Root = s.node.Clone(node.NIL)
}

func (tb *TextBuf) Reset(text string) {
	tb.Delete(0)

	if len(text) > 0 {
		tb.Insert(0, text)
	}
}

func (tb *TextBuf) Validate() {
	tb.tree.Root.Validate()
}
