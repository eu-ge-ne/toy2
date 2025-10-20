package textbuf

import (
	"iter"
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
	start, ok := buf.StartPos(startLn, startCol)
	if !ok {
		return empty
	}

	end := buf.EndPos(endLn, endCol)

	return buf.Read(start.Idx, end.Idx)
}
