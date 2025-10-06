package textbuf

import (
	"iter"
	"slices"
	"strings"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Seg struct {
	G   *grapheme.Grapheme
	Ln  int
	Col int
}

func (tb *TextBuf) IterSegLine(ln int, extra bool) iter.Seq2[int, Seg] {
	return func(yield func(int, Seg) bool) {
		seg := Seg{}

		n := 0
		w := 0

		for i, g := range grapheme.Graphemes.Iter(tb.ReadPosRange(ln, 0, ln+1, 0)) {
			seg.G = g

			if seg.G.Width < 0 {
				seg.G.Width = vt.Wchar(tb.MeasureY, tb.MeasureX, seg.G.Bytes)
			}

			w += seg.G.Width
			if w > tb.WrapWidth {
				w = seg.G.Width
				seg.Ln += 1
				seg.Col = 0
			}

			if !yield(n, seg) {
				return
			}

			seg.Col += 1
			n = i
		}

		if extra {
			seg.G = grapheme.Graphemes.Get(" ")

			w += seg.G.Width
			if w > tb.WrapWidth {
				w = seg.G.Width
				seg.Ln += 1
				seg.Col = 0
			}

			if !yield(n, seg) {
				return
			}
		}
	}
}

func (tb *TextBuf) IterSegLineSlice(ln int, extra bool, start, end int) iter.Seq2[int, Seg] {
	return func(yield func(int, Seg) bool) {
		i := 0
		for j, c := range tb.IterSegLine(ln, extra) {
			if j >= start && j < end {
				if !yield(i, c) {
					return
				}
				i += 1
			}
		}
	}
}

func (tb *TextBuf) ReadSegPosRange(startLn, startCol, endLn, endCol int) string {
	startCol = tb.segCol(startLn, startCol)
	endCol = tb.segCol(endLn, endCol)

	it := tb.ReadPosRange(startLn, startCol, endLn, endCol)

	return strings.Join(slices.Collect(it), "")
}

func (tb *TextBuf) InsertSegPos(ln, col int, text string) {
	col = tb.segCol(ln, col)

	tb.InsertPos(ln, col, text)
}

func (tb *TextBuf) DeleteSegPosRange(startLn, startCol, endLn, endCol int) {
	startCol = tb.segCol(startLn, startCol)
	endCol = tb.segCol(endLn, endCol)

	tb.DeletePosRange(startLn, startCol, endLn, endCol)
}

func (tb *TextBuf) segCol(ln, col int) int {
	c := 0

	for i, cell := range tb.IterSegLine(ln, false) {
		if i == col {
			break
		}

		if i < col {
			c += len(cell.G.Seg)
		}
	}

	return c
}
