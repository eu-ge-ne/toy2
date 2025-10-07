package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
)

func (tb *TextBuf) Insert(index int, text string) {
	if index > tb.Count() {
		return
	}

	type InsertCase int
	const (
		InsertRoot InsertCase = iota
		InsertLeft
		InsertRight
		InsertSplit
	)

	insertCase := InsertRoot
	p := node.NIL
	x := tb.tree.Root

	for x != node.NIL {
		if index <= x.Left.TotalLen {
			insertCase = InsertLeft
			p = x
			x = x.Left
			continue
		}

		index -= x.Left.TotalLen

		if index < x.Len {
			insertCase = InsertSplit
			p = x
			x = node.NIL
			continue
		}

		index -= x.Len

		insertCase = InsertRight
		p = x
		x = x.Right
	}

	if (insertCase == InsertRight) && tb.content.Growable(p) {
		tb.content.Grow(p, text)
		node.Bubble(p)
		return
	}

	child := tb.content.Create(text)

	switch insertCase {
	case InsertRoot:
		tb.tree.Root = child
		tb.tree.Root.Red = false
	case InsertLeft:
		tb.tree.InsertLeft(p, child)
	case InsertRight:
		tb.tree.InsertRight(p, child)
	case InsertSplit:
		y := tb.content.Split(p, index, 0)
		tb.tree.InsertAfter(p, y)
		tb.tree.InsertBefore(y, child)
	}
}

func (tb *TextBuf) Insert2(ln, col int, text string) {
	if ln == 0 && col == 0 {
		tb.Insert(0, text)
		return
	}

	index, ok := tb.lnColToIndex(ln, col)
	if !ok {
		return
	}

	tb.Insert(index, text)
}

func (tb *TextBuf) Append(text string) {
	tb.Insert(tb.Count(), text)
}
