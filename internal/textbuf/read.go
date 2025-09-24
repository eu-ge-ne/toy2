package textbuf

import (
	"iter"
	"math"
)

func (tb *TextBuf) Read(start int, end int) iter.Seq[string] {
	x, offset := tb.tree.Root.Find(start)
	if x == nil {
		return none
	}

	return tb.content.Read(x, offset, end-start)
}

func (tb *TextBuf) ReadToEnd(start int) iter.Seq[string] {
	return tb.Read(start, math.MaxInt)
}

func (tb *TextBuf) Read2(startLn, startCol, endLn, endCol int) iter.Seq[string] {
	start_i, ok := tb.posToIndex(startLn, startCol)
	if !ok {
		return none
	}

	end_i, ok := tb.posToIndex(endLn, endCol)
	if !ok {
		end_i = math.MaxInt
	}

	return tb.Read(start_i, end_i)
}

func (tb *TextBuf) Read2ToEnd(startLn, startCol int) iter.Seq[string] {
	start_i, ok := tb.posToIndex(startLn, startCol)
	if !ok {
		return none
	}

	return tb.Read(start_i, math.MaxInt)
}

func none(yield func(string) bool) {
}
