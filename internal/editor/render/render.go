package render

import (
	"fmt"
	"math"

	"github.com/eu-ge-ne/toy2/internal/color"
	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/syntax"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Render struct {
	buffer *textbuf.TextBuf
	cursor *cursor.Cursor
	syntax *syntax.Syntax

	area              ui.Area
	enabled           bool
	indexEnabled      bool
	whitespaceEnabled bool
	wrapEnabled       bool

	indexWidth int
	textWidth  int
	cursorY    int
	cursorX    int
	ScrollLn   int
	ScrollCol  int

	colorMainBg     []byte
	colorSelectedBg []byte
	colorVoidBg     []byte
	colorIndex      []byte
	colorCharFg     map[syntax.CharFgColor][]byte
}

func New(buffer *textbuf.TextBuf, cursor *cursor.Cursor) *Render {
	return &Render{
		buffer: buffer,
		cursor: cursor,
	}
}

func (r *Render) SetColors(t theme.Theme) {
	r.colorMainBg = t.MainBg()
	r.colorSelectedBg = t.Light2Bg()
	r.colorVoidBg = t.Dark0Bg()
	r.colorIndex = append(t.Light0Bg(), t.Dark0Fg()...)

	r.colorCharFg = map[syntax.CharFgColor][]byte{
		syntax.CharFgColorVisible:    t.Light1Fg(),
		syntax.CharFgColorWhitespace: t.Dark0Fg(),
		syntax.CharFgColorEmpty:      t.MainFg(),
		syntax.CharFgColorVariable:  vt.CharFg(color.Sky200),
		syntax.CharFgColorKeyword:  vt.CharFg(color.Purple400),
	}
}

func (r *Render) SetArea(a ui.Area) {
	r.area = a
}

func (r *Render) SetEnabled(enabled bool) {
	r.enabled = enabled
}

func (r *Render) SetIndexEnabled(enabled bool) {
	r.indexEnabled = enabled
}

func (r *Render) SetWhitespaceEnabled(enabled bool) {
	r.whitespaceEnabled = enabled
}

func (r *Render) ToggleWhitespaceEnabled() {
	r.whitespaceEnabled = !r.whitespaceEnabled
}

func (r *Render) SetWrapEnabled(enabled bool) {
	r.wrapEnabled = enabled
}

func (r *Render) ToggleWrapEnabled() {
	r.wrapEnabled = !r.wrapEnabled
}

func (r *Render) SetSyntax(s *syntax.Syntax) {
	r.syntax = s
}

func (r *Render) Render() {
	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(vt.SaveCursor)
	vt.Buf.Write(r.colorMainBg)
	vt.ClearArea(vt.Buf, r.area)

	if r.area.W >= r.indexWidth {
		r.renderLines()
	}

	if r.enabled {
		vt.SetCursor(vt.Buf, r.cursorY, r.cursorX)
		vt.Buf.Write(vt.ShowCursor)
	} else {
		vt.Buf.Write(vt.RestoreCursor)
		vt.Buf.Write(vt.ShowCursor)
	}

	vt.Buf.Flush()

	vt.Sync.Esu()
}

func (r *Render) Scroll() {
	if r.indexEnabled && r.buffer.LineCount() > 0 {
		r.indexWidth = int(math.Log10(float64(r.buffer.LineCount()))) + 3
	} else {
		r.indexWidth = 0
	}

	r.textWidth = r.area.W - r.indexWidth

	if r.wrapEnabled {
		r.buffer.WrapWidth = r.textWidth
	} else {
		r.buffer.WrapWidth = math.MaxInt
	}

	r.cursorY = r.area.Y
	r.cursorX = r.area.X + r.indexWidth

	r.buffer.MeasureY = r.area.Y
	r.buffer.MeasureX = r.area.X + r.indexWidth

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
		for j, cell := range r.buffer.IterLine(r.ScrollLn+i, false) {
			if j > 0 && cell.Col == 0 {
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

	deltaCol := col - r.ScrollCol

	// Before?

	if deltaCol <= 0 {
		r.ScrollCol = col
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

		r.ScrollCol += 1
		width -= w
	}

	r.cursorX += width
}

func (r *Render) renderLines() {
	row := r.area.Y

	for ln := r.ScrollLn; ; ln += 1 {
		if ln < r.buffer.LineCount() {
			row = r.renderLine(ln, row)
		} else {
			vt.SetCursor(vt.Buf, row, r.area.X)
			vt.Buf.Write(r.colorVoidBg)
			vt.ClearLine(vt.Buf, r.area.W)
		}

		row += 1
		if row >= r.area.Y+r.area.H {
			break
		}
	}
}

func (r *Render) renderLine(ln int, row int) int {
	currentFg := syntax.CharFgColorUndefined
	currentBg := false
	availableW := 0

	for i, cell := range r.buffer.IterLine(ln, false) {
		if cell.Col == 0 {
			if i > 0 {
				row += 1
				if row >= r.area.Y+r.area.H {
					return row
				}
			}

			vt.SetCursor(vt.Buf, row, r.area.X)

			if r.indexWidth > 0 {
				if i == 0 {
					vt.Buf.Write(r.colorIndex)
					fmt.Fprintf(vt.Buf, "%*d ", r.indexWidth-1, ln+1)
					vt.Buf.Write(r.colorMainBg)
				} else {
					vt.Buf.Write(r.colorMainBg)
					vt.WriteSpaces(vt.Buf, r.indexWidth)
				}
			}

			availableW = r.area.W - r.indexWidth
		}

		if (cell.Col < r.ScrollCol) || (cell.G.Width > availableW) {
			continue
		}

		colorBg := r.cursor.IsSelected(ln, i)
		if colorBg != currentBg {
			currentBg = colorBg
			if currentBg {
				vt.Buf.Write(r.colorSelectedBg)
			} else {
				vt.Buf.Write(r.colorMainBg)
			}
		}

		start, _ := r.buffer.Index(ln, i)
		end := start + len(cell.G.Seg)
		colorFg := r.syntax.HighlightSpan(start, end)
		if colorFg == syntax.CharFgColorUndefined {
			if cell.G.IsVisible {
				colorFg = syntax.CharFgColorVisible
			} else if r.whitespaceEnabled {
				colorFg = syntax.CharFgColorWhitespace
			} else {
				colorFg = syntax.CharFgColorEmpty
			}
		}
		if colorFg != currentFg {
			currentFg = colorFg
			vt.Buf.Write(r.colorCharFg[colorFg])
		}

		vt.Buf.Write(cell.G.Bytes)

		availableW -= cell.G.Width
	}

	return row
}
