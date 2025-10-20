package render

import (
	"iter"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
)

type cell struct {
	Gr      *grapheme.Grapheme
	Col     int
	WrapLn  int
	WrapCol int
}

func (r *Render) wrapLine(ln int, extra bool) iter.Seq[cell] {
	return func(yield func(cell) bool) {
		cell := cell{}
		w := 0

		for _, gr := range r.buffer.LineGraphemes(ln) {
			cell.Gr = gr

			w += cell.Gr.Width
			if w > r.wrapWidth {
				w = cell.Gr.Width
				cell.WrapLn += 1
				cell.WrapCol = 0
			}

			if !yield(cell) {
				return
			}

			cell.Col += 1
			cell.WrapCol += 1
		}

		if extra {
			cell.Gr = grapheme.Graphemes.Get(" ")

			w += cell.Gr.Width
			if w > r.wrapWidth {
				w = cell.Gr.Width
				cell.WrapLn += 1
				cell.WrapCol = 0
			}

			if !yield(cell) {
				return
			}
		}
	}
}
