package scroll

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/ui"
)

type Scroll struct {
	IndexEnabled bool
	WrapEnabled  bool
	IndexWidth   int
	TextWidth    int
	WrapWidth    int

	Ln  int
	Col int
	Y   int
	X   int

	buffer *textbuf.TextBuf
}

func New(buffer *textbuf.TextBuf) *Scroll {
	return &Scroll{buffer: buffer}
}

func (s *Scroll) Scroll(area ui.Area, toLn, toCol int) {
	if s.IndexEnabled && s.buffer.LineCount() > 0 {
		s.IndexWidth = int(math.Log10(float64(s.buffer.LineCount()))) + 3
	} else {
		s.IndexWidth = 0
	}

	s.TextWidth = area.W - s.IndexWidth

	if s.WrapEnabled {
		s.WrapWidth = s.TextWidth
	} else {
		s.WrapWidth = math.MaxInt
	}

	s.Y = area.Y
	s.X = area.X + s.IndexWidth

	grapheme.Graphemes.SetWcharPos(area.Y, area.X+s.IndexWidth)

	s.scrollV(area, toLn)
	s.scrollH(toLn, toCol)
}

func (s *Scroll) scrollV(area ui.Area, toLn int) {
	deltaLn := toLn - s.Ln

	// Above?
	if deltaLn <= 0 {
		s.Ln = toLn
		return
	}

	// Below?

	if deltaLn > area.H {
		s.Ln = toLn - area.H
	}

	xs := make([]int, toLn+1-s.Ln)
	for i := 0; i < len(xs); i += 1 {
		xs[i] = 1
		for cell := range grapheme.Wrap(s.buffer.LineGraphemes(s.Ln+i), s.WrapWidth, false) {
			if cell.Col > 0 && cell.WrapCol == 0 {
				xs[i] += 1
			}
		}
	}

	i := 0
	height := std.Sum(xs)

	for height > area.H {
		height -= xs[i]
		s.Ln += 1
		i += 1
	}

	for i < len(xs)-1 {
		s.Y += xs[i]
		i += 1
	}
}

func (s *Scroll) scrollH(toLn, toCol int) {
	var cell *grapheme.Cell = nil
	for c := range grapheme.Wrap(s.buffer.LineGraphemes(toLn), s.WrapWidth, true) {
		if c.Col >= toCol {
			cell = &c
			break
		}
	}
	if cell != nil {
		s.Y += cell.WrapLn
	}

	col := 0
	if cell != nil {
		col = cell.WrapCol
	}

	deltaCol := col - s.Col

	// Before?

	if deltaCol <= 0 {
		s.Col = col
		return
	}

	// After?

	xs := make([]int, deltaCol)
	xsI := 0
	for c := range grapheme.Wrap(s.buffer.LineGraphemes(toLn), s.WrapWidth, true) {
		if c.Col >= toCol-deltaCol && c.Col < toCol {
			xs[xsI] = c.Gr.Width
			xsI += 1
		}
	}

	width := std.Sum(xs)

	for _, w := range xs {
		if width < s.TextWidth {
			break
		}

		s.Col += 1
		width -= w
	}

	s.X += width
}
