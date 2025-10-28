package render

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/ui"
)

type scroll struct {
	buffer *textbuf.TextBuf

	indexEnabled      bool
	whitespaceEnabled bool
	wrapEnabled       bool

	indexWidth int
	textWidth  int
	wrapWidth  int

	cursorY int
	cursorX int

	ln  int
	col int
}

func (s *scroll) scroll(area ui.Area, curLn, curCol int) {
	if s.indexEnabled && s.buffer.LineCount() > 0 {
		s.indexWidth = int(math.Log10(float64(s.buffer.LineCount()))) + 3
	} else {
		s.indexWidth = 0
	}

	s.textWidth = area.W - s.indexWidth

	if s.wrapEnabled {
		s.wrapWidth = s.textWidth
	} else {
		s.wrapWidth = math.MaxInt
	}

	s.cursorY = area.Y
	s.cursorX = area.X + s.indexWidth

	grapheme.Graphemes.SetWcharPos(area.Y, area.X+s.indexWidth)

	s.scrollV(area, curLn)
	s.scrollH(curLn, curCol)
}

func (s *scroll) scrollV(area ui.Area, curLn int) {
	deltaLn := curLn - s.ln

	// Above?
	if deltaLn <= 0 {
		s.ln = curLn
		return
	}

	// Below?

	if deltaLn > area.H {
		s.ln = curLn - area.H
	}

	xs := make([]int, curLn+1-s.ln)
	for i := 0; i < len(xs); i += 1 {
		xs[i] = 1
		for cell := range wrap(s.buffer.LineGraphemes(s.ln+i), s.wrapWidth, false) {
			if cell.Col > 0 && cell.WrapCol == 0 {
				xs[i] += 1
			}
		}
	}

	i := 0
	height := std.Sum(xs)

	for height > area.H {
		height -= xs[i]
		s.ln += 1
		i += 1
	}

	for i < len(xs)-1 {
		s.cursorY += xs[i]
		i += 1
	}
}

func (s *scroll) scrollH(curLn, curCol int) {
	var cell *cell = nil
	for c := range wrap(s.buffer.LineGraphemes(curLn), s.wrapWidth, true) {
		if c.Col >= curCol {
			cell = &c
			break
		}
	}
	if cell != nil {
		s.cursorY += cell.WrapLn
	}

	col := 0
	if cell != nil {
		col = cell.WrapCol
	}

	deltaCol := col - s.col

	// Before?

	if deltaCol <= 0 {
		s.col = col
		return
	}

	// After?

	xs := make([]int, deltaCol)
	xsI := 0
	for c := range wrap(s.buffer.LineGraphemes(curLn), s.wrapWidth, true) {
		if c.Col >= curCol-deltaCol && c.Col < curCol {
			xs[xsI] = c.Gr.Width
			xsI += 1
		}
	}

	width := std.Sum(xs)

	for _, w := range xs {
		if width < s.textWidth {
			break
		}

		s.col += 1
		width -= w
	}

	s.cursorX += width
}
