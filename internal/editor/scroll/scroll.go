package scroll

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/ui"
)

type Scroll struct {
	Area         ui.Area
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
	cursor *cursor.Cursor
}

func New(buffer *textbuf.TextBuf, cursor *cursor.Cursor) *Scroll {
	return &Scroll{buffer: buffer, cursor: cursor}
}

func (s *Scroll) Scroll() {
	if s.IndexEnabled && s.buffer.LineCount() > 0 {
		s.IndexWidth = int(math.Log10(float64(s.buffer.LineCount()))) + 3
	} else {
		s.IndexWidth = 0
	}

	s.TextWidth = s.Area.W - s.IndexWidth

	if s.WrapEnabled {
		s.WrapWidth = s.TextWidth
	} else {
		s.WrapWidth = math.MaxInt
	}

	s.Y = s.Area.Y
	s.X = s.Area.X + s.IndexWidth

	grapheme.Graphemes.SetWcharPos(s.Area.Y, s.Area.X+s.IndexWidth)

	s.scrollV()
	s.scrollH()
}

func (s *Scroll) scrollV() {
	deltaLn := s.cursor.Ln - s.Ln

	// Above?
	if deltaLn <= 0 {
		s.Ln = s.cursor.Ln
		return
	}

	// Below?

	if deltaLn > s.Area.H {
		s.Ln = s.cursor.Ln - s.Area.H
	}

	xs := make([]int, s.cursor.Ln+1-s.Ln)
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

	for height > s.Area.H {
		height -= xs[i]
		s.Ln += 1
		i += 1
	}

	for i < len(xs)-1 {
		s.Y += xs[i]
		i += 1
	}
}

func (s *Scroll) scrollH() {
	var cell *grapheme.Cell = nil
	for c := range grapheme.Wrap(s.buffer.LineGraphemes(s.cursor.Ln), s.WrapWidth, true) {
		if c.Col >= s.cursor.Col {
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
	for c := range grapheme.Wrap(s.buffer.LineGraphemes(s.cursor.Ln), s.WrapWidth, true) {
		if c.Col >= s.cursor.Col-deltaCol && c.Col < s.cursor.Col {
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
