package grapheme

import (
	"iter"
)

type Cell struct {
	Gr      *Grapheme
	Col     int
	WrapLn  int
	WrapCol int
}

func Wrap(line iter.Seq[*Grapheme], width int, extra bool) iter.Seq[Cell] {
	return func(yield func(Cell) bool) {
		cell := Cell{}
		w := 0

		for gr := range line {
			cell.Gr = gr

			w += cell.Gr.Width
			if w > width {
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
			cell.Gr = nil

			w += 1
			if w > width {
				w = 1
				cell.WrapLn += 1
				cell.WrapCol = 0
			}

			yield(cell)
		}
	}
}

func WrapHeight(line iter.Seq[*Grapheme], width int) int {
	h := 1
	w := 0

	for gr := range line {
		w += gr.Width

		if w > width {
			w = gr.Width
			h += 1
		}
	}

	return h
}
