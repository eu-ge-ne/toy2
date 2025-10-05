package textbuf

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/textbuf/content"
	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
	"github.com/eu-ge-ne/toy2/internal/textbuf/tree"
)

type TextBuf struct {
	WrapWidth int
	MeasureY  int
	MeasureX  int

	content content.Content
	tree    tree.Tree
}

type Snapshot struct {
	node *node.Node
}

func New() TextBuf {
	return TextBuf{
		WrapWidth: math.MaxInt,

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
	tb.DeleteIndex(0)

	if len(text) > 0 {
		tb.InsertIndex(0, text)
	}
}

func (tb *TextBuf) Validate() {
	tb.tree.Root.Validate()
}
