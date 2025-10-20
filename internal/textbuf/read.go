package textbuf

import (
	"iter"
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
)

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
		return func(yield func(string) bool) {}
	}

	return buf.content.Read(x, offset, end-start)
}

func (buf *TextBuf) Read2(startLn, startCol, endLn, endCol int) iter.Seq[string] {
	start, ok := buf.Pos(startLn, startCol)
	if !ok {
		return func(yield func(string) bool) {}
	}

	end := buf.PosNear(endLn, endCol)

	return buf.Read(start.Idx, end.Idx)
}

func (buf *TextBuf) ReadLine(ln int) iter.Seq[string] {
	start, ok := buf.lnIdx(ln)
	if !ok {
		return func(yield func(string) bool) {}
	}

	end, ok := buf.lnIdx(ln + 1)
	if !ok {
		end = math.MaxInt
	}

	return buf.Read(start, end)
}

func (buf *TextBuf) LineGraphemes(ln int) iter.Seq2[int, *grapheme.Grapheme] {
	line := buf.ReadLine(ln)

	return grapheme.Graphemes.FromString(line)
}
