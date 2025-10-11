package textbuf

import (
	"iter"
	"math"
	"slices"
	"strings"
)

func (buf *TextBuf) Iter() iter.Seq[string] {
	return buf.Read(0, math.MaxInt)
}

func (buf *TextBuf) All() string {
	return strings.Join(slices.Collect(buf.Iter()), "")
}

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
		return none
	}

	return buf.content.Read(x, offset, end-start)
}

func (buf *TextBuf) Read2(startLn, startCol, endLn, endCol int) string {
	start, ok := buf.lnColIndex(startLn, startCol)
	if !ok {
		return ""
	}

	end, ok := buf.lnColIndex(endLn, endCol)
	if !ok {
		end = math.MaxInt
	}

	it := buf.Read(start, end)

	return strings.Join(slices.Collect(it), "")
}

func none(yield func(string) bool) {
}
