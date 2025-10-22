package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
)

type insCase int

const (
	insCaseRoot insCase = iota
	insCaseLeft
	insCaseRight
	insCaseSplit
)

func (buf *TextBuf) insert(idx int, text string) {
	if idx > buf.Count() {
		return
	}

	insertCase := insCaseRoot
	p := node.NIL
	x := buf.tree.Root

	for x != node.NIL {
		if idx <= x.Left.TotalLen {
			insertCase = insCaseLeft
			p = x
			x = x.Left
			continue
		}

		idx -= x.Left.TotalLen

		if idx < x.Len {
			insertCase = insCaseSplit
			p = x
			x = node.NIL
			continue
		}

		idx -= x.Len

		insertCase = insCaseRight
		p = x
		x = x.Right
	}

	if (insertCase == insCaseRight) && buf.content.Growable(p) {
		buf.content.Grow(p, text)
		node.Bubble(p)
		return
	}

	child := buf.content.Create(text)

	switch insertCase {
	case insCaseRoot:
		buf.tree.Root = child
		buf.tree.Root.Red = false
	case insCaseLeft:
		buf.tree.InsertLeft(p, child)
	case insCaseRight:
		buf.tree.InsertRight(p, child)
	case insCaseSplit:
		y := buf.content.Split(p, idx, 0)
		buf.tree.InsertAfter(p, y)
		buf.tree.InsertBefore(y, child)
	}
}

func (buf *TextBuf) Insert(ln, col int, text string) (Pos, Pos) {
	startPos := buf.PosMax(ln, col)

	dLn, dCol := grapheme.Graphemes.MeasureString(text)
	var endLn, endCol int
	if dLn == 0 {
		endLn = ln
		endCol = col + dCol
	} else {
		endLn = ln + dLn
		endCol = dCol
	}
	endPos := buf.PosMax(endLn, endCol)

	buf.insert(startPos.Idx, text)

	return startPos, endPos
}

func (buf *TextBuf) Append(text string) {
	buf.insert(buf.Count(), text)
}
