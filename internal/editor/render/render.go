package render

import (
	"fmt"
	"math"

	"github.com/eu-ge-ne/toy2/internal/color"
	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/syntax"
	"github.com/eu-ge-ne/toy2/internal/grapheme"
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

	hlSpans chan syntax.Span
	hlSpan  syntax.Span
	hlPos   int

	colorMainBg     []byte
	colorMainFg     []byte
	colorSelectedBg []byte
	colorVoidBg     []byte
	colorIndex      []byte
	colorCharFg     map[string][]byte
}

func New(buffer *textbuf.TextBuf, cursor *cursor.Cursor) *Render {
	return &Render{
		buffer: buffer,
		cursor: cursor,
	}
}

func (r *Render) SetColors(t theme.Theme) {
	r.colorMainBg = t.MainBg()
	r.colorMainFg = t.Light1Fg()
	r.colorSelectedBg = t.Light2Bg()
	r.colorVoidBg = t.Dark0Bg()
	r.colorIndex = append(t.Light0Bg(), t.Dark0Fg()...)

	r.colorCharFg = map[string][]byte{
		"_text":        r.colorMainFg,
		"_ws_enabled":  t.Dark0Fg(),
		"_ws_disabled": t.MainFg(),
		"keyword":      vt.CharFg(color.Fuchsia300),
		"comment":      vt.CharFg([3]byte{0x6A, 0x99, 0x55}),
		"function":     vt.CharFg([3]byte{0xDC, 0xDC, 0xAA}),
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
	r.scroll()
	r.initHighlight()

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

func (r *Render) scroll() {
	if r.indexEnabled && r.buffer.LineCount() > 0 {
		r.indexWidth = int(math.Log10(float64(r.buffer.LineCount()))) + 3
	} else {
		r.indexWidth = 0
	}

	r.textWidth = r.area.W - r.indexWidth

	if r.wrapEnabled {
		grapheme.Graphemes.SetWrapWidth(r.textWidth)
	} else {
		grapheme.Graphemes.SetWrapWidth(math.MaxInt)
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
		line := r.buffer.ReadLine(r.ScrollLn + i)
		for cell := range grapheme.Graphemes.IterString(line, false, 0, math.MaxInt) {
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
	var cell *grapheme.IterCell = nil
	line := r.buffer.ReadLine(r.cursor.Ln)
	for c := range grapheme.Graphemes.IterString(line, true, r.cursor.Col, math.MaxInt) {
		cell = &c
		break
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
	line = r.buffer.ReadLine(r.cursor.Ln)
	for c := range grapheme.Graphemes.IterString(line, true, r.cursor.Col-deltaCol, r.cursor.Col) {
		xs[xsI] = c.Gr.Width
		xsI += 1
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

func (r *Render) initHighlight() {
	r.hlSpans = r.syntax.Highlight(r.ScrollLn, r.ScrollLn+r.area.H)
	r.hlSpan = syntax.Span{StartByte: -1, EndByte: -1}
	r.hlPos, _ = r.buffer.LnToByte(r.ScrollLn)
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
	currentFg := ""
	currentBg := false
	availableW := 0

	line := r.buffer.ReadLine(ln)
	for cell := range grapheme.Graphemes.IterString(line, false, 0, math.MaxInt) {
		if cell.WrapCol == 0 {
			if cell.Col > 0 {
				row += 1
				if row >= r.area.Y+r.area.H {
					return row
				}
			}

			vt.SetCursor(vt.Buf, row, r.area.X)

			if r.indexWidth > 0 {
				if cell.Col == 0 {
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

		spanName := r.nextSegSpanName(len(cell.Gr.Seg))

		if (cell.WrapCol < r.ScrollCol) || (cell.Gr.Width > availableW) {
			continue
		}

		bg := r.cursor.IsSelected(ln, cell.Col)
		if bg != currentBg {
			currentBg = bg
			if currentBg {
				vt.Buf.Write(r.colorSelectedBg)
			} else {
				vt.Buf.Write(r.colorMainBg)
			}
		}

		fg := spanName
		if len(fg) == 0 {
			if cell.Gr.IsVisible {
				fg = "_text"
			} else if r.whitespaceEnabled {
				fg = "_ws_enabled"
			} else {
				fg = "_ws_disabled"
			}
		}
		if fg != currentFg {
			currentFg = fg
			b, ok := r.colorCharFg[fg]
			if !ok {
				b = r.colorMainFg
			}
			vt.Buf.Write(b)
		}

		vt.Buf.Write(cell.Gr.Bytes)

		availableW -= cell.Gr.Width
	}

	return row
}

func (r *Render) nextSegSpanName(l int) string {
	var name string

	if r.hlPos >= r.hlSpan.EndByte {
		if s, ok := <-r.hlSpans; ok {
			r.hlSpan = s
		}
	}

	if r.hlPos >= r.hlSpan.StartByte && r.hlPos < r.hlSpan.EndByte {
		name = r.hlSpan.Name
	}

	r.hlPos += l

	return name
}
