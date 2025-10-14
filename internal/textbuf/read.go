package textbuf

import (
	"iter"
	"math"
)

var empty = func(yield func([]byte) bool) {}

func (buf *TextBuf) Chunk(i int) []byte {
	x, offset := buf.tree.Root.Find(i)
	if x == nil {
		return nil
	}

	return buf.content.Chunk(x, offset)
}

func (buf *TextBuf) Read(start int, end int) iter.Seq[[]byte] {
	x, offset := buf.tree.Root.Find(start)
	if x == nil {
		return empty
	}

	return buf.content.Read(x, offset, end-start)
}

func (buf *TextBuf) Read2(startLn, startCol, endLn, endCol int) iter.Seq[[]byte] {
	start, ok := buf.Index(startLn, startCol)
	if !ok {
		return empty
	}

	end, ok := buf.Index(endLn, endCol)
	if !ok {
		end = math.MaxInt
	}

	return buf.Read(start, end)
}
