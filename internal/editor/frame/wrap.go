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

func wrap(line iter.Seq[*grapheme.Grapheme], width int, extra bool) iter.Seq[cell] {
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

type cell2 struct {
	Gr  *grapheme.Grapheme
	Col int
}

func width(line iter.Seq[*grapheme.Grapheme]) iter.Seq[cell2] {
	return func(yield func(cell2) bool) {
		cell := cell2{}

		for gr := range line {
			cell.Gr = gr

			if !yield(cell) {
				return
			}

			cell.Col += 1
		}

		cell.Gr = nil

		yield(cell)
	}
}
