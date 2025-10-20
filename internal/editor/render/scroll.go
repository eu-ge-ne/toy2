package render

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/std"
)

func (r *Render) scroll() {
	if r.indexEnabled && r.buffer.LineCount() > 0 {
		r.indexWidth = int(math.Log10(float64(r.buffer.LineCount()))) + 3
	} else {
		r.indexWidth = 0
	}

	r.textWidth = r.area.W - r.indexWidth

	if r.wrapEnabled {
		r.wrapWidth = r.textWidth
	} else {
		r.wrapWidth = math.MaxInt
	}

	r.cursorY = r.area.Y
	r.cursorX = r.area.X + r.indexWidth

	grapheme.Graphemes.SetWcharPos(r.area.Y, r.area.X+r.indexWidth)

	r.scrollV()
	r.scrollH()
}

func (r *Render) scrollV() {
	deltaLn := r.cursor.Ln - r.ScrollLn

	// Above?
	if deltaLn <= 0 {
		r.ScrollLn = r.cursor.Ln
		return
	}

	// Below?

	if deltaLn > r.area.H {
		r.ScrollLn = r.cursor.Ln - r.area.H
	}

	xs := make([]int, r.cursor.Ln+1-r.ScrollLn)
	for i := 0; i < len(xs); i += 1 {
		xs[i] = 1
		for cell := range r.wrapLine(r.ScrollLn+i, false) {
			if cell.Col > 0 && cell.WrapCol == 0 {
				xs[i] += 1
			}
		}
	}

	i := 0
	height := std.Sum(xs)

	for height > r.area.H {
		height -= xs[i]
		r.ScrollLn += 1
		i += 1
	}

	for i < len(xs)-1 {
		r.cursorY += xs[i]
		i += 1
	}
}

func (r *Render) scrollH() {
	var cell *cell = nil
	for c := range r.wrapLine(r.cursor.Ln, true) {
		if c.Col >= r.cursor.Col {
			cell = &c
			break
		}
	}
	if cell != nil {
		r.cursorY += cell.WrapLn
	}

	col := 0
	if cell != nil {
		col = cell.WrapCol
	}

	deltaCol := col - r.ScrollCol

	// Before?

	if deltaCol <= 0 {
		r.ScrollCol = col
		return
	}

	// After?

	xs := make([]int, deltaCol)
	xsI := 0
	for c := range r.wrapLine(r.cursor.Ln, true) {
		if c.Col >= r.cursor.Col-deltaCol && c.Col < r.cursor.Col {
			xs[xsI] = c.Gr.Width
			xsI += 1
		}
	}

	width := std.Sum(xs)

	for _, w := range xs {
		if width < r.textWidth {
			break
		}

		r.ScrollCol += 1
		width -= w
	}

	r.cursorX += width
}
