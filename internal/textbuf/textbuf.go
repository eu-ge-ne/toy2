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

func (buf *TextBuf) Index(ln, col int) (int, bool) {
	lnIndex, ok := buf.lnIndex(ln)
	if !ok {
		return 0, false
	}

	colIndex := 0

	for i, cell := range buf.IterLine(ln, true) {
		if i == col {
			return lnIndex + colIndex, true
		}

		colIndex += len(cell.G.Seg)
	}

	return 0, false
}

func (buf TextBuf) lnIndex(ln int) (int, bool) {
	if buf.Count() == 0 {
		return 0, false
	}

	if ln == 0 {
		return 0, true
	}

	eolIndex := ln - 1
	x := buf.tree.Root
	i := 0

	for x != node.NIL {
		if eolIndex < x.Left.TotalEolsLen {
			x = x.Left
			continue
		}

		eolIndex -= x.Left.TotalEolsLen
		i += x.Left.TotalLen

		if eolIndex < x.EolsLen {
			buf := buf.content.Table[x.PieceIndex]
			eolEnd := buf.Eols[x.EolsStart+eolIndex].End
			return i + eolEnd - x.Start, true
		}

		eolIndex -= x.EolsLen
		i += x.Len
		x = x.Right
	}

	return 0, false
}
