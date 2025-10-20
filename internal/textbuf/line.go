package textbuf

import (
	"iter"
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type LineCell struct {
	G   *grapheme.Grapheme
	Ln  int
	Col int
}

func (buf *TextBuf) IterLine(ln int, extra bool) iter.Seq2[int, LineCell] {
	return func(yield func(int, LineCell) bool) {
		start, ok := buf.LnToByte(ln)
		if !ok {
			return
		}

		end, ok := buf.LnToByte(ln + 1)
		if !ok {
			end = math.MaxInt
		}

		cell := LineCell{}

		i := 0
		w := 0

		for g := range grapheme.Graphemes.Iter(buf.Read(start, end)) {
			cell.G = g

			if cell.G.Width < 0 {
				cell.G.Width = vt.Wchar(buf.MeasureY, buf.MeasureX, cell.G.Bytes)
			}

			w += cell.G.Width
			if w > buf.WrapWidth {
				w = cell.G.Width
				cell.Ln += 1
				cell.Col = 0
			}

			if !yield(i, cell) {
				return
			}

			cell.Col += 1
			i += 1
		}

		if extra {
			cell.G = grapheme.Graphemes.Get(" ")

			w += cell.G.Width
			if w > buf.WrapWidth {
				w = cell.G.Width
				cell.Ln += 1
				cell.Col = 0
			}

			if !yield(i, cell) {
				return
			}
		}
	}
}
