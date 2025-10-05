package textbuf

import "github.com/eu-ge-ne/toy2/internal/textbuf/node"

func (tb *TextBuf) InsertIndex(i int, text string) {
	if i > tb.Count() {
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
		if i <= x.Left.TotalLen {
			insertCase = InsertLeft
			p = x
			x = x.Left
			continue
		}

		i -= x.Left.TotalLen

		if i < x.Len {
			insertCase = InsertSplit
			p = x
			x = node.NIL
			continue
		}

		i -= x.Len

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
		y := tb.content.Split(p, i, 0)
		tb.tree.InsertAfter(p, y)
		tb.tree.InsertBefore(y, child)
	}
}

func (tb *TextBuf) InsertPos(ln, col int, text string) {
	i, ok := tb.posToIndex(ln, col)
	if !ok {
		return
	}

	tb.InsertIndex(i, text)
}

func (tb *TextBuf) Append(text string) {
	tb.InsertIndex(tb.Count(), text)
}
