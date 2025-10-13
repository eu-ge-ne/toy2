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

	return string(buf.content.Chunk(x, offset))
}

func (buf *TextBuf) Read(start int, end int) iter.Seq[string] {
	return func(yield func(string) bool) {
		x, offset := buf.tree.Root.Find(start)
		if x == nil {
			return
		}

		for b := range buf.content.Read(x, offset, end-start) {
			if !yield(string(b)) {
				return
			}
		}
	}

}

func (buf *TextBuf) Read2(startLn, startCol, endLn, endCol int) string {
	start, ok := buf.Index(startLn, startCol)
	if !ok {
		return ""
	}

	end, ok := buf.Index(endLn, endCol)
	if !ok {
		end = math.MaxInt
	}

	it := buf.Read(start, end)

	return strings.Join(slices.Collect(it), "")
}
