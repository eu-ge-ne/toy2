package frame

import (
	"fmt"
	"math"

	"github.com/eu-ge-ne/toy2/internal/colors"
	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/syntax"
	"github.com/eu-ge-ne/toy2/internal/grapheme"
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

func (f *Frame) SetColors(t theme.Theme) {
	f.colorMainBg = t.MainBg()
	f.colorMainFg = t.Light1Fg()
	f.colorSelectedBg = t.Light2Bg()
	f.colorVoidBg = t.Dark0Bg()
	f.colorIndex = append(t.Light0Bg(), t.Dark0Fg()...)

	f.colorCharFg = map[string][]byte{
		"toy.text":            f.colorMainFg,
		"toy.wspace.on":       t.Dark0Fg(),
		"toy.wspace.off":      t.MainFg(),
		"keyword":             vt.CharFg(colors.Fuchsia300),
		"function":            vt.CharFg(colors.Yellow100),
		"punctuation.bracket": vt.CharFg(colors.Yellow300),
		"comment":             vt.CharFg(colors.Lime600),
		"variable":            vt.CharFg(colors.Sky200),
		"constructor":         vt.CharFg(colors.Sky200),
		"type":                vt.CharFg(colors.Sky200),
	}
}

func (f *Frame) SetArea(a ui.Area) {
	f.area = a
}

func (f *Frame) SetIndexEnabled(e bool) {
	f.indexEnabled = e
}

func (f *Frame) SetWrapEnabled(e bool) {
	f.wrapEnabled = e
}

func (f *Frame) ToggleWrapEnabled() {
	f.wrapEnabled = !f.wrapEnabled
	f.cursor.Home(false)
}

func (f *Frame) SetWhitespaceEnabled(e bool) {
	f.whitespaceEnabled = e
}

func (f *Frame) ToggleWhitespaceEnabled() {
	f.whitespaceEnabled = !f.whitespaceEnabled
	f.cursor.Home(false)
}

func (f *Frame) Render(setCursor bool) {
	f.scroll()

	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(vt.SaveCursor)
	vt.Buf.Write(f.colorMainBg)
	vt.ClearArea(vt.Buf, f.area)

	f.renderLines()

	if setCursor {
		vt.SetCursor(vt.Buf, f.cursorY, f.cursorX)
		vt.Buf.Write(vt.ShowCursor)
	} else {
		vt.Buf.Write(vt.RestoreCursor)
		vt.Buf.Write(vt.ShowCursor)
	}

	vt.Buf.Flush()

	vt.Sync.Esu()
}

func (f *Frame) scroll() {
	if f.indexEnabled && f.buffer.LineCount() > 0 {
		f.indexWidth = int(math.Log10(float64(f.buffer.LineCount()))) + 3
	} else {
		f.indexWidth = 0
	}

	f.textWidth = f.area.W - f.indexWidth

	if f.wrapEnabled {
		f.wrapWidth = f.textWidth
	} else {
		f.wrapWidth = math.MaxInt
	}

	grapheme.Graphemes.SetWcharPos(f.area.Y, f.area.X+f.indexWidth)

	f.scrollV()

	f.syntax.Highlight(f.buffer, f.scrollLn, f.scrollLn+f.area.H)

	f.scrollH()
}

func (f *Frame) scrollV() {
	f.cursorY = f.area.Y

	lnDelta := f.cursor.Ln - f.scrollLn

	// Above?
	if lnDelta <= 0 {
		f.scrollLn = f.cursor.Ln
		return
	}

	// Below?
	if lnDelta > f.area.H {
		lnDelta = f.area.H
		f.scrollLn = f.cursor.Ln - lnDelta
	}

	yDelta := 0

	for i := f.scrollLn + lnDelta; i >= f.scrollLn; i -= 1 {
		h := wrapCount(f.buffer.LineGraphemes(i), f.wrapWidth)

		if i == f.scrollLn+lnDelta {
			f.cursorY -= h
		}

		if yDelta+h > f.area.H {
			f.scrollLn += 1
		} else {
			yDelta += h
			f.cursorY += h
		}
	}
}

func (f *Frame) scrollH() {
	f.cursorX = f.area.X + f.indexWidth

	wrapLn, wrapCol := findWrapCol(f.buffer.LineGraphemes(f.cursor.Ln), f.wrapWidth, f.cursor.Col)
	f.cursorY += wrapLn

	delta := wrapCol - f.scrollCol

	// Before?
	if delta <= 0 {
		f.scrollCol = wrapCol
		return
	}

	// After?
	wSum, ww := sliceWidth(f.buffer.LineGraphemes(f.cursor.Ln), f.cursor.Col-delta, f.cursor.Col)

	i := 0

	for wSum >= f.textWidth {
		wSum -= ww[i]
		i += 1
		f.scrollCol += 1
	}

	f.cursorX += wSum
}

func (f *Frame) renderLines() {
	row := f.area.Y

	for ln := f.scrollLn; ; ln += 1 {
		if ln < f.buffer.LineCount() {
			row = f.renderLine(ln, row)
		} else {
			vt.SetCursor(vt.Buf, row, f.area.X)
			vt.Buf.Write(f.colorVoidBg)
			vt.ClearLine(vt.Buf, f.area.W)
		}

		row += 1
		if row >= f.area.Y+f.area.H {
			break
		}
	}
}

func (f *Frame) renderLine(ln int, row int) int {
	currentFg := ""
	currentBg := false
	availableW := 0

	for cell := range wrap(f.buffer.LineGraphemes(ln), f.wrapWidth) {
		if cell.WrapCol == 0 {
			if cell.Col > 0 {
				row += 1
				if row >= f.area.Y+f.area.H {
					return row
				}
			}

			vt.SetCursor(vt.Buf, row, f.area.X)

			if f.indexWidth > 0 {
				if cell.Col == 0 {
					vt.Buf.Write(f.colorIndex)
					fmt.Fprintf(vt.Buf, "%*d ", f.indexWidth-1, ln+1)
					vt.Buf.Write(f.colorMainBg)
				} else {
					vt.Buf.Write(f.colorMainBg)
					vt.WriteSpaces(vt.Buf, f.indexWidth)
				}
			}

			availableW = f.area.W - f.indexWidth
		}

		fg := f.syntax.Next(len(cell.Gr.Str))

		if (cell.WrapCol < f.scrollCol) || (cell.Gr.Width > availableW) {
			continue
		}

		bg := f.cursor.IsSelected(ln, cell.Col)
		if bg != currentBg {
			currentBg = bg
			if currentBg {
				vt.Buf.Write(f.colorSelectedBg)
			} else {
				vt.Buf.Write(f.colorMainBg)
			}
		}

		if len(fg) == 0 {
			if cell.Gr.IsVisible {
				fg = "toy.text"
			} else if f.whitespaceEnabled {
				fg = "toy.wspace.on"
			} else {
				fg = "toy.wspace.off"
			}
		}
		if fg != currentFg {
			currentFg = fg
			b, ok := f.colorCharFg[fg]
			if !ok {
				b = f.colorMainFg
			}
			vt.Buf.Write(b)
		}

		vt.Buf.Write(cell.Gr.Bytes)

		availableW -= cell.Gr.Width
	}

	return row
}
