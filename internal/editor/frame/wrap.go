package frame

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

func wrap(line iter.Seq[*grapheme.Grapheme], width int) iter.Seq[cell] {
	return func(yield func(cell) bool) {
		cell := cell{}
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
	}
}

func findWrapCol(line iter.Seq[*grapheme.Grapheme], width int, col int) (int, int) {
	i, wrapLn, wrapCol := 0, 0, 0

	w := 0

	for gr := range line {
		w += gr.Width
		if w > width {
			w = gr.Width
			wrapLn += 1
			wrapCol = 0
		}

		if i == col {
			return wrapLn, wrapCol
		}

		i += 1
		wrapCol += 1
	}

	w += 1
	if w > width {
		w = 1
		wrapLn += 1
		wrapCol = 0
	}

	if i == col {
		return wrapLn, wrapCol
	}

	return 0, 0
}

func wrapCount(line iter.Seq[*grapheme.Grapheme], wrapWidth int) int {
	h := 1
	w := 0

	for gr := range line {
		w += gr.Width
		if w > wrapWidth {
			w = gr.Width
			h += 1
		}
	}

	return h
}

func width(line iter.Seq[*grapheme.Grapheme], start, end int) (int, []int) {
	sum := 0
	ww := make([]int, end-start)
	i := 0
	col := 0

	for gr := range line {
		if col >= start && col < end {
			sum += gr.Width
			ww[i] = gr.Width
			i += 1
		}

		col += 1
	}

	return sum, ww
}
