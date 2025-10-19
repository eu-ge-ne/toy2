package textbuf

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
)

type InsertCase int

const (
	InsertRoot InsertCase = iota
	InsertLeft
	InsertRight
	InsertSplit
)

func (buf *TextBuf) Insert(index int, text string) {
	if index > buf.Count() {
		return
	}

	insertCase := InsertRoot
	p := node.NIL
	x := buf.tree.Root

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

	if (insertCase == InsertRight) && buf.content.Growable(p) {
		buf.content.Grow(p, text)
		node.Bubble(p)
		return
	}

	child := buf.content.Create(text)

	switch insertCase {
	case InsertRoot:
		buf.tree.Root = child
		buf.tree.Root.Red = false
	case InsertLeft:
		buf.tree.InsertLeft(p, child)
	case InsertRight:
		buf.tree.InsertRight(p, child)
	case InsertSplit:
		y := buf.content.Split(p, index, 0)
		buf.tree.InsertAfter(p, y)
		buf.tree.InsertBefore(y, child)
	}
}

func (buf *TextBuf) Insert2(ln, col int, text string) {
	index, ok := buf.LnIndex(ln)

	if !ok {
		if ln == 0 {
			index = 0
		} else {
			index = max(0, buf.LineCount()-1)
		}
	}

	for cell := range buf.IterLine(ln, false, 0, math.MaxInt) {
		if cell.Col == col {
			break
		}
		index += len(cell.Gr.Seg)
	}

	buf.Insert(index, text)
}

func (buf *TextBuf) Append(text string) {
	buf.Insert(buf.Count(), text)
}
