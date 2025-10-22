package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
)

func (buf *TextBuf) Delete(startIdx int, endIdx int) {
	n, offset := buf.tree.Root.Find(startIdx)
	if n == nil {
		return
	}

	count := endIdx - startIdx
	offset2 := offset + count

	if offset2 == n.Len {
		if offset == 0 {
			buf.tree.Delete(n)
		} else {
			buf.content.TrimEnd(n, count)
			node.Bubble(n)
		}
	} else if offset2 < n.Len {
		if offset == 0 {
			buf.content.TrimStart(n, count)
			node.Bubble(n)
		} else {
			y := buf.content.Split(n, offset, count)
			buf.tree.InsertAfter(n, y)
		}
	} else {
		x := n
		i := 0

		if offset != 0 {
			y := buf.content.Split(n, offset, 0)
			buf.tree.InsertAfter(n, y)
			x = y
		}

		lastNode, lastOffset := buf.tree.Root.Find(endIdx)
		if lastNode != nil && lastOffset != 0 {
			y := buf.content.Split(lastNode, lastOffset, 0)
			buf.tree.InsertAfter(lastNode, y)
		}

		for x != node.NIL && (i < count) {
			i += x.Len

			next := node.Successor(x)

			buf.tree.Delete(x)

			x = next
		}
	}
}

func (buf *TextBuf) Delete2(startLn, startCol, endLn, endCol int) (Pos, Pos, bool) {
	startPos, ok := buf.Pos(startLn, startCol)
	if !ok {
		return Pos{}, Pos{}, false
	}

	endPos := buf.PosMax(endLn, endCol)

	buf.Delete(startPos.Idx, endPos.Idx)

	return startPos, endPos, true
}
