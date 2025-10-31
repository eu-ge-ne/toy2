package frame

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

type Frame struct {
	area              ui.Area
	indexEnabled      bool
	wrapEnabled       bool
	whitespaceEnabled bool

	indexWidth int
	textWidth  int
	wrapWidth  int

	scrollLn  int
	scrollCol int
	cursorY   int
	cursorX   int

	colorMainBg     []byte
	colorMainFg     []byte
	colorSelectedBg []byte
	colorVoidBg     []byte
	colorIndex      []byte
	colorCharFg     map[string][]byte

	buffer *textbuf.TextBuf
	cursor *cursor.Cursor
	syntax *syntax.Syntax
}

func New(buffer *textbuf.TextBuf, cursor *cursor.Cursor, syntax *syntax.Syntax) *Frame {
	return &Frame{
		buffer: buffer,
		cursor: cursor,
		syntax: syntax,
	}
}

func (fr *Frame) SetColors(t theme.Theme) {
	fr.colorMainBg = t.MainBg()
	fr.colorMainFg = t.Light1Fg()
	fr.colorSelectedBg = t.Light2Bg()
	fr.colorVoidBg = t.Dark0Bg()
	fr.colorIndex = append(t.Light0Bg(), t.Dark0Fg()...)

	fr.colorCharFg = map[string][]byte{
		"_text":        fr.colorMainFg,
		"_ws_enabled":  t.Dark0Fg(),
		"_ws_disabled": t.MainFg(),
		"keyword":      vt.CharFg(color.Fuchsia300),
		"comment":      vt.CharFg([3]byte{0x6A, 0x99, 0x55}),
		"function":     vt.CharFg([3]byte{0xDC, 0xDC, 0xAA}),
	}
}

func (fr *Frame) SetArea(a ui.Area) {
	fr.area = a
}

func (fr *Frame) SetIndexEnabled(e bool) {
	fr.indexEnabled = e
}

func (fr *Frame) SetWrapEnabled(e bool) {
	fr.wrapEnabled = e
}

func (fr *Frame) ToggleWrapEnabled() {
	fr.wrapEnabled = !fr.wrapEnabled
	fr.cursor.Home(false)
}

func (fr *Frame) SetWhitespaceEnabled(e bool) {
	fr.whitespaceEnabled = e
}

func (fr *Frame) ToggleWhitespaceEnabled() {
	fr.whitespaceEnabled = !fr.whitespaceEnabled
	fr.cursor.Home(false)
}

func (fr *Frame) Render(setCursor bool) {
	fr.scroll()

	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(vt.SaveCursor)
	vt.Buf.Write(fr.colorMainBg)
	vt.ClearArea(vt.Buf, fr.area)

	if fr.area.W >= fr.indexWidth {
		fr.renderLines()
	}

	if setCursor {
		vt.SetCursor(vt.Buf, fr.cursorY, fr.cursorX)
		vt.Buf.Write(vt.ShowCursor)
	} else {
		vt.Buf.Write(vt.RestoreCursor)
		vt.Buf.Write(vt.ShowCursor)
	}

	vt.Buf.Flush()

	vt.Sync.Esu()
}

func (fr *Frame) scroll() {
	if fr.indexEnabled && fr.buffer.LineCount() > 0 {
		fr.indexWidth = int(math.Log10(float64(fr.buffer.LineCount()))) + 3
	} else {
		fr.indexWidth = 0
	}

	fr.textWidth = fr.area.W - fr.indexWidth

	if fr.wrapEnabled {
		fr.wrapWidth = fr.textWidth
	} else {
		fr.wrapWidth = math.MaxInt
	}

	fr.cursorY = fr.area.Y
	fr.cursorX = fr.area.X + fr.indexWidth

	grapheme.Graphemes.SetWcharPos(fr.area.Y, fr.area.X+fr.indexWidth)

	fr.scrollV()

	fr.syntax.Highlight(fr.buffer, fr.scrollLn, fr.scrollLn+fr.area.H)

	fr.scrollH()
}

func (fr *Frame) scrollV() {
	delta := fr.cursor.Ln - fr.scrollLn

	// Above?
	if delta <= 0 {
		fr.scrollLn = fr.cursor.Ln
		return
	}

	// Below?

	if delta > fr.area.H {
		fr.scrollLn = fr.cursor.Ln - fr.area.H
	}

	heights := make([]int, fr.cursor.Ln+1-fr.scrollLn)

	for i := 0; i < len(heights); i += 1 {
		heights[i] = 1

		for cell := range grapheme.Wrap(fr.buffer.LineGraphemes(fr.scrollLn+i), fr.wrapWidth, false) {
			if cell.Col > 0 && cell.WrapCol == 0 {
				heights[i] += 1
			}
		}
	}

	i := 0
	height := std.Sum(heights)

	for height > fr.area.H {
		height -= heights[i]
		fr.scrollLn += 1
		i += 1
	}

	for i < len(heights)-1 {
		fr.cursorY += heights[i]
		i += 1
	}
}

func (fr *Frame) scrollH() {
	var cell *grapheme.Cell = nil
	for c := range grapheme.Wrap(fr.buffer.LineGraphemes(fr.cursor.Ln), fr.wrapWidth, true) {
		if c.Col >= fr.cursor.Col {
			cell = &c
			break
		}
	}
	if cell != nil {
		fr.cursorY += cell.WrapLn
	}

	col := 0
	if cell != nil {
		col = cell.WrapCol
	}

	deltaCol := col - fr.scrollCol

	// Before?

	if deltaCol <= 0 {
		fr.scrollCol = col
		return
	}

	// After?

	xs := make([]int, deltaCol)
	xsI := 0
	for c := range grapheme.Wrap(fr.buffer.LineGraphemes(fr.cursor.Ln), fr.wrapWidth, true) {
		if c.Col >= fr.cursor.Col-deltaCol && c.Col < fr.cursor.Col {
			xs[xsI] = c.Gr.Width
			xsI += 1
		}
	}

	width := std.Sum(xs)

	for _, w := range xs {
		if width < fr.textWidth {
			break
		}

		fr.scrollCol += 1
		width -= w
	}

	fr.cursorX += width
}

func (fr *Frame) renderLines() {
	row := fr.area.Y

	for ln := fr.scrollLn; ; ln += 1 {
		if ln < fr.buffer.LineCount() {
			row = fr.renderLine(ln, row)
		} else {
			vt.SetCursor(vt.Buf, row, fr.area.X)
			vt.Buf.Write(fr.colorVoidBg)
			vt.ClearLine(vt.Buf, fr.area.W)
		}

		row += 1
		if row >= fr.area.Y+fr.area.H {
			break
		}
	}
}

func (fr *Frame) renderLine(ln int, row int) int {
	currentFg := ""
	currentBg := false
	availableW := 0

	for cell := range grapheme.Wrap(fr.buffer.LineGraphemes(ln), fr.wrapWidth, false) {
		if cell.WrapCol == 0 {
			if cell.Col > 0 {
				row += 1
				if row >= fr.area.Y+fr.area.H {
					return row
				}
			}

			vt.SetCursor(vt.Buf, row, fr.area.X)

			if fr.indexWidth > 0 {
				if cell.Col == 0 {
					vt.Buf.Write(fr.colorIndex)
					fmt.Fprintf(vt.Buf, "%*d ", fr.indexWidth-1, ln+1)
					vt.Buf.Write(fr.colorMainBg)
				} else {
					vt.Buf.Write(fr.colorMainBg)
					vt.WriteSpaces(vt.Buf, fr.indexWidth)
				}
			}

			availableW = fr.area.W - fr.indexWidth
		}

		fg := fr.syntax.Next(len(cell.Gr.Str))

		if (cell.WrapCol < fr.scrollCol) || (cell.Gr.Width > availableW) {
			continue
		}

		bg := fr.cursor.IsSelected(ln, cell.Col)
		if bg != currentBg {
			currentBg = bg
			if currentBg {
				vt.Buf.Write(fr.colorSelectedBg)
			} else {
				vt.Buf.Write(fr.colorMainBg)
			}
		}

		if len(fg) == 0 {
			if cell.Gr.IsVisible {
				fg = "_text"
			} else if fr.whitespaceEnabled {
				fg = "_ws_enabled"
			} else {
				fg = "_ws_disabled"
			}
		}
		if fg != currentFg {
			currentFg = fg
			b, ok := fr.colorCharFg[fg]
			if !ok {
				b = fr.colorMainFg
			}
			vt.Buf.Write(b)
		}

		vt.Buf.Write(cell.Gr.Bytes)

		availableW -= cell.Gr.Width
	}

	return row
}
