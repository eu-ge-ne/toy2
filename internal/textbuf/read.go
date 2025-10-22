package textbuf

import (
	"iter"
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
)

func (buf *TextBuf) Chunk(idx int) string {
	x, offset := buf.tree.Root.Find(idx)
	if x == nil {
		return ""
	}

	return buf.content.Chunk(x, offset)
}

func (buf *TextBuf) Slice(startIdx int, endIdx int) iter.Seq[string] {
	x, offset := buf.tree.Root.Find(startIdx)
	if x == nil {
		return func(yield func(string) bool) {}
	}

	return buf.content.Read(x, offset, endIdx-startIdx)
}

func (buf *TextBuf) Read(startLn, startCol, endLn, endCol int) iter.Seq[string] {
	startPos, ok := buf.Pos(startLn, startCol)
	if !ok {
		return func(yield func(string) bool) {}
	}

	endPos := buf.PosMax(endLn, endCol)

	return buf.Slice(startPos.Idx, endPos.Idx)
}

func (buf *TextBuf) ReadLine(ln int) iter.Seq[string] {
	startIdx, ok := buf.lnIdx(ln)
	if !ok {
		return func(yield func(string) bool) {}
	}

	endIdx, ok := buf.lnIdx(ln + 1)
	if !ok {
		endIdx = math.MaxInt
	}

	return buf.Slice(startIdx, endIdx)
}

func (buf *TextBuf) LineGraphemes(ln int) iter.Seq2[int, *grapheme.Grapheme] {
	return grapheme.Graphemes.FromString(buf.ReadLine(ln))
}
