package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
)

func (buf *TextBuf) Insert(index int, data []byte) {
	if index > buf.Count() {
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
		buf.content.Grow(p, data)
		node.Bubble(p)
		return
	}

	child := buf.content.Create(data)

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

func (buf *TextBuf) InsertString(index int, text string) {
	buf.Insert(index, []byte(text))
}

func (buf *TextBuf) Insert2(ln, col int, data []byte) {
	index, ok := buf.LnIndex(ln)

	if !ok {
		if ln == 0 {
			index = 0
		} else {
			index = max(0, buf.LineCount()-1)
		}
	}

	for i, cell := range buf.IterLine(ln, false) {
		if i == col {
			break
		}
		index += len(cell.G.Seg)
	}

	buf.Insert(index, data)
}

func (buf *TextBuf) Append(data []byte) {
	buf.Insert(buf.Count(), data)
}
