package textbuf

import (
	"iter"
	"math"
)

var empty = func(yield func(string) bool) {}

func (buf *TextBuf) Chunk(i int) string {
	x, offset := buf.tree.Root.Find(i)
	if x == nil {
		return ""
	}

	return buf.content.Chunk(x, offset)
}

func (buf *TextBuf) Read(start int, end int) iter.Seq[string] {
	x, offset := buf.tree.Root.Find(start)
	if x == nil {
		return empty
	}

	return buf.content.Read(x, offset, end-start)
}

func (buf *TextBuf) Read2(startLn, startCol, endLn, endCol int) iter.Seq[string] {
	start, ok := buf.lnColToByte(startLn, startCol)
	if !ok {
		return empty
	}

	end, ok := buf.lnColToByte(endLn, endCol)
	if !ok {
		end = math.MaxInt
	}

	return buf.Read(start, end)
}
