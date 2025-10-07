package textbuf

import (
	"iter"
	"math"
	"slices"
	"strings"
)

func (tb *TextBuf) Iter() iter.Seq[string] {
	return tb.ReadSlice(0, math.MaxInt)
}

func (tb *TextBuf) Read() string {
	return strings.Join(slices.Collect(tb.Iter()), "")
}

func (tb *TextBuf) ReadSlice(start int, end int) iter.Seq[string] {
	x, offset := tb.tree.Root.Find(start)
	if x == nil {
		return none
	}

	return tb.content.Read(x, offset, end-start)
}

func (tb *TextBuf) ReadSlice2(startLn, startCol, endLn, endCol int) string {
	start, ok := tb.lnColToIndex(startLn, startCol)
	if !ok {
		return ""
	}

	end, ok := tb.lnColToIndex(endLn, endCol)
	if !ok {
		end = math.MaxInt
	}

	it := tb.ReadSlice(start, end)

	return strings.Join(slices.Collect(it), "")
}

func none(yield func(string) bool) {
}
