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
	if tb.Count() == 0 {
		return 0
	}
	return tb.tree.Root.TotalEolsLen + 1
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
	tb.DeleteSlice(0, math.MaxInt)

	if len(text) > 0 {
		tb.Insert(0, text)
	}
}

func (tb *TextBuf) Validate() {
	tb.tree.Root.Validate()
}

func (tb TextBuf) lnIndex(ln int) (int, bool) {
	if tb.Count() == 0 {
		return 0, false
	}

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

func (tb *TextBuf) lnColToIndex(ln, col int) (int, bool) {
	lnIndex, ok := tb.lnIndex(ln)
	if !ok {
		return 0, false
	}

	colIndex := 0

	for i, cell := range tb.IterLine(ln, false) {
		if i == col {
			return lnIndex + colIndex, true
		}

		colIndex += len(cell.G.Seg)
	}

	return 0, false
}
