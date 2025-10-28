package render

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/std"
)

func (r *Render) scroll(curLn, curCol int) {
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

	r.scrollV(curLn)
	r.scrollH(curLn, curCol)
}

func (r *Render) scrollV(curLn int) {
	deltaLn := curLn - r.ScrollLn

	// Above?
	if deltaLn <= 0 {
		r.ScrollLn = curLn
		return
	}

	// Below?

	if deltaLn > r.area.H {
		r.ScrollLn = curLn - r.area.H
	}

	xs := make([]int, curLn+1-r.ScrollLn)
	for i := 0; i < len(xs); i += 1 {
		xs[i] = 1
		for cell := range wrap(r.buffer.LineGraphemes(r.ScrollLn+i), r.wrapWidth, false) {
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

func (r *Render) scrollH(curLn, curCol int) {
	var cell *cell = nil
	for c := range wrap(r.buffer.LineGraphemes(curLn), r.wrapWidth, true) {
		if c.Col >= curCol {
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
	for c := range wrap(r.buffer.LineGraphemes(curLn), r.wrapWidth, true) {
		if c.Col >= curCol-deltaCol && c.Col < curCol {
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
