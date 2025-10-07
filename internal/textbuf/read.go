package textbuf

import (
	"iter"
	"math"
	"slices"
	"strings"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (tb *TextBuf) Iter() iter.Seq[string] {
	return tb.ReadSlice(0, math.MaxInt)
}

func (tb *TextBuf) Read() string {
	return strings.Join(slices.Collect(tb.Iter()), "")
}

func (tb *TextBuf) ReadSlice(start int, end int) iter.Seq[string] {
	x, offset := tb.tree.Root.Find(start)
	if x == nil {
		return none
	}

	return tb.content.Read(x, offset, end-start)
}

func (tb *TextBuf) ReadSlice2(startLn, startCol, endLn, endCol int) string {
	start, ok := tb.lnColToIndex(startLn, startCol)
	if !ok {
		return ""
	}

	end, ok := tb.lnColToIndex(endLn, endCol)
	if !ok {
		end = math.MaxInt
	}

	it := tb.ReadSlice(start, end)

	return strings.Join(slices.Collect(it), "")
}

func none(yield func(string) bool) {
}

type LineCell struct {
	G   *grapheme.Grapheme
	Ln  int
	Col int
}

func (tb *TextBuf) IterLine(ln int, extra bool) iter.Seq2[int, LineCell] {
	return func(yield func(int, LineCell) bool) {
		start, ok := tb.lnToIndex(ln)
		if !ok {
			return
		}

		end, ok := tb.lnToIndex(ln + 1)
		if !ok {
			end = math.MaxInt
		}

		cell := LineCell{}

		n := 0
		w := 0

		for i, g := range grapheme.Graphemes.Iter(tb.ReadSlice(start, end)) {
			cell.G = g

			if cell.G.Width < 0 {
				cell.G.Width = vt.Wchar(tb.MeasureY, tb.MeasureX, cell.G.Bytes)
			}

			w += cell.G.Width
			if w > tb.WrapWidth {
				w = cell.G.Width
				cell.Ln += 1
				cell.Col = 0
			}

			if !yield(i, cell) {
				return
			}

			cell.Col += 1
			n = i
		}

		if extra {
			cell.G = grapheme.Graphemes.Get(" ")

			w += cell.G.Width
			if w > tb.WrapWidth {
				w = cell.G.Width
				cell.Ln += 1
				cell.Col = 0
			}

			if !yield(n, cell) {
				return
			}
		}
	}
}

func (tb *TextBuf) IterLineSlice(ln int, extra bool, start, end int) iter.Seq2[int, LineCell] {
	return func(yield func(int, LineCell) bool) {
		i := 0
		for j, c := range tb.IterLine(ln, extra) {
			if j >= start && j < end {
				if !yield(i, c) {
					return
				}
				i += 1
			}
		}
	}
}
