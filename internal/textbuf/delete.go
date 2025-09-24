package textbuf

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/textbuf/internal/node"
)

func (buf *TextBuf) Delete(start int, end int) {
	n, offset := buf.tree.Root.Find(start)
	if n == nil {
		return
	}

	count := end - start
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
			buf.tree.InsertAfter(n, &y)
		}
	} else {
		x := n
		i := 0

		if offset != 0 {
			y := buf.content.Split(n, offset, 0)
			buf.tree.InsertAfter(n, &y)
			x = &y
		}

		lastNode, lastOffset := buf.tree.Root.Find(end)
		if lastNode != nil && lastOffset != 0 {
			y := buf.content.Split(lastNode, lastOffset, 0)
			buf.tree.InsertAfter(lastNode, &y)
		}

		for x != node.NIL && (i < count) {
			i += x.Len

			next := node.Successor(x)

			buf.tree.Delete(x)

			x = next
		}
	}
}

func (buf *TextBuf) DeleteToEnd(start int) {
	buf.Delete(start, math.MaxInt)
}

func (buf *TextBuf) Delete2(startLn, startCol, endLn, endCol int) {
	start_i, ok := buf.posToIndex(startLn, startCol)
	if !ok {
		return
	}

	end_i, ok := buf.posToIndex(endLn, endCol)
	if !ok {
		end_i = math.MaxInt
	}

	buf.Delete(start_i, end_i)
}

func (buf *TextBuf) Delete2ToEnd(startLn, startCol int) {
	start_i, ok := buf.posToIndex(startLn, startCol)
	if !ok {
		return
	}

	buf.Delete(start_i, math.MaxInt)
}
