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

func (tb TextBuf) posToIndex(ln, col int) (int, bool) {
	i, ok := tb.findLineStart(ln)

	if !ok {
		return 0, false
	}

	return i + col, true
}

func (tb TextBuf) findLineStart(ln int) (int, bool) {
	if ln == 0 {
		return 0, true
	}

	eolIndex := ln - 1
	x := tb.tree.Root
	i := 0

	for x != node.NIL {
		if eolIndex < x.Left.TotalEolsLen {
			x = x.Left
			continue
		}

		eolIndex -= x.Left.TotalEolsLen
		i += x.Left.TotalLen

		if eolIndex < x.EolsLen {
			buf := tb.content.Table[x.PieceIndex]
			eolEnd := buf.Eols[x.EolsStart+eolIndex].End
			return i + eolEnd - x.Start, true
		}

		eolIndex -= x.EolsLen
		i += x.Len
		x = x.Right
	}

	return 0, false
}
