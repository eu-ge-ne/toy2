package textbuf

import (
	"bytes"
	"iter"
	"math"
	"slices"
)

var empty = func(yield func([]byte) bool) {}

func (buf *TextBuf) Iter() iter.Seq[[]byte] {
	return buf.Read(0, math.MaxInt)
}

func (buf *TextBuf) All() string {
	return string(bytes.Join(slices.Collect(buf.Iter()), []byte{}))
}

func (buf *TextBuf) Chunk(i int) string {
	x, offset := buf.tree.Root.Find(i)
	if x == nil {
		return ""
	}

	return string(buf.content.Chunk(x, offset))
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
