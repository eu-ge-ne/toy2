package textbuf

import (
	"iter"
	"slices"
	"strings"

	"github.com/rivo/uniseg"

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
		i := 0
		c := Seg{}
		w := 0

		for chunk := range tb.ReadPosRange(ln, 0, ln+1, 0) {
			gr := uniseg.NewGraphemes(chunk)

			for gr.Next() {
				c.G = grapheme.Graphemes.Get(gr.Str())

				if c.G.Width < 0 {
					c.G.Width = vt.Wchar(tb.MeasureY, tb.MeasureX, c.G.Bytes)
				}

				w += c.G.Width
				if w > tb.WrapWidth {
					w = c.G.Width
					c.Ln += 1
					c.Col = 0
				}

				if !yield(i, c) {
					return
				}

				i += 1
				c.Col += 1
			}
		}

		if extra {
			c.G = grapheme.Graphemes.Get(" ")

			w += c.G.Width
			if w > tb.WrapWidth {
				w = c.G.Width
				c.Ln += 1
				c.Col = 0
			}

			if !yield(i, c) {
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
