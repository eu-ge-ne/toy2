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
	Area              ui.Area
	Enabled           bool
	IndexEnabled      bool
	WrapEnabled       bool
	WhitespaceEnabled bool

	indexWidth int
	textWidth  int
	wrapWidth  int

	ScrollLn  int
	ScrollCol int
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
}

func New(buffer *textbuf.TextBuf, cursor *cursor.Cursor) *Frame {
	return &Frame{
		buffer: buffer,
		cursor: cursor,
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

func (fr *Frame) Scroll() {
	if fr.IndexEnabled && fr.buffer.LineCount() > 0 {
		fr.indexWidth = int(math.Log10(float64(fr.buffer.LineCount()))) + 3
	} else {
		fr.indexWidth = 0
	}

	fr.textWidth = fr.Area.W - fr.indexWidth

	if fr.WrapEnabled {
		fr.wrapWidth = fr.textWidth
	} else {
		fr.wrapWidth = math.MaxInt
	}

	fr.cursorY = fr.Area.Y
	fr.cursorX = fr.Area.X + fr.indexWidth

	grapheme.Graphemes.SetWcharPos(fr.Area.Y, fr.Area.X+fr.indexWidth)

	fr.scrollV()
	fr.scrollH()
}

func (fr *Frame) Render(hl *syntax.Highlight) {
	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(vt.SaveCursor)
	vt.Buf.Write(fr.colorMainBg)
	vt.ClearArea(vt.Buf, fr.Area)

	if fr.Area.W >= fr.indexWidth {
		fr.renderLines(hl)
	}

	if fr.Enabled {
		vt.SetCursor(vt.Buf, fr.cursorY, fr.cursorX)
		vt.Buf.Write(vt.ShowCursor)
	} else {
		vt.Buf.Write(vt.RestoreCursor)
		vt.Buf.Write(vt.ShowCursor)
	}

	vt.Buf.Flush()

	vt.Sync.Esu()
}

func (fr *Frame) scrollV() {
	deltaLn := fr.cursor.Ln - fr.ScrollLn

	// Above?
	if deltaLn <= 0 {
		fr.ScrollLn = fr.cursor.Ln
		return
	}

	// Below?

	if deltaLn > fr.Area.H {
		fr.ScrollLn = fr.cursor.Ln - fr.Area.H
	}

	xs := make([]int, fr.cursor.Ln+1-fr.ScrollLn)
	for i := 0; i < len(xs); i += 1 {
		xs[i] = 1
		for cell := range grapheme.Wrap(fr.buffer.LineGraphemes(fr.ScrollLn+i), fr.wrapWidth, false) {
			if cell.Col > 0 && cell.WrapCol == 0 {
				xs[i] += 1
			}
		}
	}

	i := 0
	height := std.Sum(xs)

	for height > fr.Area.H {
		height -= xs[i]
		fr.ScrollLn += 1
		i += 1
	}

	for i < len(xs)-1 {
		fr.cursorY += xs[i]
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

	deltaCol := col - fr.ScrollCol

	// Before?

	if deltaCol <= 0 {
		fr.ScrollCol = col
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

		fr.ScrollCol += 1
		width -= w
	}

	fr.cursorX += width
}

func (fr *Frame) renderLines(hl *syntax.Highlight) {
	row := fr.Area.Y

	for ln := fr.ScrollLn; ; ln += 1 {
		if ln < fr.buffer.LineCount() {
			row = fr.renderLine(hl, ln, row)
		} else {
			vt.SetCursor(vt.Buf, row, fr.Area.X)
			vt.Buf.Write(fr.colorVoidBg)
			vt.ClearLine(vt.Buf, fr.Area.W)
		}

		row += 1
		if row >= fr.Area.Y+fr.Area.H {
			break
		}
	}
}

func (fr *Frame) renderLine(hl *syntax.Highlight, ln int, row int) int {
	currentFg := ""
	currentBg := false
	availableW := 0

	for cell := range grapheme.Wrap(fr.buffer.LineGraphemes(ln), fr.wrapWidth, false) {
		if cell.WrapCol == 0 {
			if cell.Col > 0 {
				row += 1
				if row >= fr.Area.Y+fr.Area.H {
					return row
				}
			}

			vt.SetCursor(vt.Buf, row, fr.Area.X)

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

			availableW = fr.Area.W - fr.indexWidth
		}

		fg := hl.Next(len(cell.Gr.Str))

		if (cell.WrapCol < fr.ScrollCol) || (cell.Gr.Width > availableW) {
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
			} else if fr.WhitespaceEnabled {
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
