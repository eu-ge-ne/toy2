package textbuf

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/textbuf/internal/node"
)

func (tb *TextBuf) DeleteRange(start int, end int) {
	n, offset := tb.tree.Root.Find(start)
	if n == nil {
		return
	}

	count := end - start
	offset2 := offset + count

	if offset2 == n.Len {
		if offset == 0 {
			tb.tree.Delete(n)
		} else {
			tb.content.TrimEnd(n, count)
			node.Bubble(n)
		}
	} else if offset2 < n.Len {
		if offset == 0 {
			tb.content.TrimStart(n, count)
			node.Bubble(n)
		} else {
			y := tb.content.Split(n, offset, count)
			tb.tree.InsertAfter(n, y)
		}
	} else {
		x := n
		i := 0

		if offset != 0 {
			y := tb.content.Split(n, offset, 0)
			tb.tree.InsertAfter(n, y)
			x = y
		}

		lastNode, lastOffset := tb.tree.Root.Find(end)
		if lastNode != nil && lastOffset != 0 {
			y := tb.content.Split(lastNode, lastOffset, 0)
			tb.tree.InsertAfter(lastNode, y)
		}

		for x != node.NIL && (i < count) {
			i += x.Len

			next := node.Successor(x)

			tb.tree.Delete(x)

			x = next
		}
	}
}

func (tb *TextBuf) Delete(start int) {
	tb.DeleteRange(start, math.MaxInt)
}

func (tb *TextBuf) Delete2Range(startLn, startCol, endLn, endCol int) {
	start_i, ok := tb.posToIndex(startLn, startCol)
	if !ok {
		return
	}

	end_i, ok := tb.posToIndex(endLn, endCol)
	if !ok {
		end_i = math.MaxInt
	}

	tb.DeleteRange(start_i, end_i)
}

func (tb *TextBuf) Delete2(startLn, startCol int) {
	start_i, ok := tb.posToIndex(startLn, startCol)
	if !ok {
		return
	}

	tb.DeleteRange(start_i, math.MaxInt)
}
