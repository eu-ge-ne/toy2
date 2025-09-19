package textbuf

import (
	"iter"
	"math"
)

func (buf *TextBuf) Read(start int, end int) iter.Seq[string] {
	x, offset := buf.tree.Root.Find(start)
	if x == nil {
		return none
	}

	return buf.content.Read(x, offset, end-start)
}

func (buf *TextBuf) ReadToEnd(start int) iter.Seq[string] {
	return buf.Read(start, math.MaxInt)
}

func (buf *TextBuf) Read2(startLn, startCol, endLn, endCol int) iter.Seq[string] {
	start_i, ok := buf.posToIndex(startLn, startCol)
	if !ok {
		return none
	}

	end_i, ok := buf.posToIndex(endLn, endCol)
	if !ok {
		end_i = math.MaxInt
	}

	return buf.Read(start_i, end_i)
}

func (buf *TextBuf) Read2ToEnd(startLn, startCol int) iter.Seq[string] {
	start_i, ok := buf.posToIndex(startLn, startCol)
	if !ok {
		return none
	}

	return buf.Read(start_i, math.MaxInt)
}

func none(yield func(string) bool) {
}
