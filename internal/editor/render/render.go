package render

import (
	"fmt"
	"math"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Render struct {
	buffer *textbuf.TextBuf
	cursor *cursor.Cursor

	Colors            Colors
	Area              ui.Area
	Enabled           bool
	IndexEnabled      bool
	WhitespaceEnabled bool
	WrapEnabled       bool

	indexWidth int
	textWidth  int
	cursorY    int
	cursorX    int
	scrollLn   int
	scrollCol  int
}

func New(buffer *textbuf.TextBuf, cursor *cursor.Cursor) Render {
	return Render{
		buffer: buffer,
		cursor: cursor,
	}
}

func (r *Render) Render() {
	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(vt.SaveCursor)
	vt.Buf.Write(r.Colors.background)
	vt.ClearArea(vt.Buf, r.Area)

	r.determineLayout()

	if r.Area.W >= r.indexWidth {
		r.scrollV()
		r.scrollH()
		r.renderLines()
	}

	if r.Enabled {
		vt.SetCursor(vt.Buf, r.cursorY, r.cursorX)
		vt.Buf.Write(vt.ShowCursor)
	} else {
		vt.Buf.Write(vt.RestoreCursor)
		vt.Buf.Write(vt.ShowCursor)
	}

	vt.Buf.Flush()

	vt.Sync.Esu()
}

func (r *Render) determineLayout() {
	if r.IndexEnabled && r.buffer.LineCount() > 0 {
		r.indexWidth = int(math.Log10(float64(r.buffer.LineCount()))) + 3
	} else {
		r.indexWidth = 0
	}

	r.textWidth = r.Area.W - r.indexWidth

	if r.WrapEnabled {
		r.buffer.WrapWidth = r.textWidth
	} else {
		r.buffer.WrapWidth = math.MaxInt
	}

	r.cursorY = r.Area.Y
	r.cursorX = r.Area.X + r.indexWidth

	r.buffer.MeasureY = r.Area.Y
	r.buffer.MeasureX = r.Area.X + r.indexWidth
}

func (r *Render) scrollV() {
	deltaLn := r.cursor.Ln - r.scrollLn

	// Above?
	if deltaLn <= 0 {
		r.scrollLn = r.cursor.Ln
		return
	}

	// Below?

	if deltaLn > r.Area.H {
		r.scrollLn = r.cursor.Ln - r.Area.H
	}

	xs := make([]int, r.cursor.Ln+1-r.scrollLn)
	for i := 0; i < len(xs); i += 1 {
		xs[i] = 1
		for j, cell := range r.buffer.IterLine(r.scrollLn+i, false) {
			if j > 0 && cell.Col == 0 {
				xs[i] += 1
			}
		}
	}

	i := 0
	height := std.Sum(xs)

	for height > r.Area.H {
		height -= xs[i]
		r.scrollLn += 1
		i += 1
	}

	for i < len(xs)-1 {
		r.cursorY += xs[i]
		i += 1
	}
}

func (r *Render) scrollH() {
	var cell *textbuf.LineCell = nil
	for _, c := range r.buffer.IterLine2(r.cursor.Ln, true, r.cursor.Col, math.MaxInt) {
		cell = &c
		break
	}
	if cell != nil {
		r.cursorY += cell.Ln
	}

	col := 0
	if cell != nil {
		col = cell.Col
	}

	deltaCol := col - r.scrollCol

	// Before?

	if deltaCol <= 0 {
		r.scrollCol = col
		return
	}

	// After?

	xs := make([]int, deltaCol)
	for i, c := range r.buffer.IterLine2(r.cursor.Ln, true, r.cursor.Col-deltaCol, r.cursor.Col) {
		xs[i] = c.G.Width
	}

	width := std.Sum(xs)

	for _, w := range xs {
		if width < r.textWidth {
			break
		}

		r.scrollCol += 1
		width -= w
	}

	r.cursorX += width
}

func (r *Render) renderLines() {
	row := r.Area.Y

	for ln := r.scrollLn; ; ln += 1 {
		if ln < r.buffer.LineCount() {
			row = r.renderLine(ln, row)
		} else {
			vt.SetCursor(vt.Buf, row, r.Area.X)
			vt.Buf.Write(r.Colors.void)
			vt.ClearLine(vt.Buf, r.Area.W)
		}

		row += 1
		if row >= r.Area.Y+r.Area.H {
			break
		}
	}
}

func (r *Render) renderLine(ln int, row int) int {
	availableW := 0
	currentColor := charColorUndefined

	for i, cell := range r.buffer.IterLine(ln, false) {
		if cell.Col == 0 {
			if i > 0 {
				row += 1
				if row >= r.Area.Y+r.Area.H {
					return row
				}
			}

			vt.SetCursor(vt.Buf, row, r.Area.X)

			if r.indexWidth > 0 {
				if i == 0 {
					vt.Buf.Write(r.Colors.index)
					fmt.Fprintf(vt.Buf, "%*d ", r.indexWidth-1, ln+1)
				} else {
					vt.Buf.Write(r.Colors.background)
					vt.WriteSpaces(vt.Buf, r.indexWidth)
				}
			}

			availableW = r.Area.W - r.indexWidth
		}

		if (cell.Col < r.scrollCol) || (cell.G.Width > availableW) {
			continue
		}

		color := newCharColor(r.cursor.IsSelected(ln, i), cell.G.IsVisible, r.WhitespaceEnabled)
		if color != currentColor {
			currentColor = color
			vt.Buf.Write(r.Colors.char[color])
		}

		vt.Buf.Write(cell.G.Bytes)

		availableW -= cell.G.Width
	}

	return row
}
