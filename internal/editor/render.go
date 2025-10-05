package editor

import (
	"fmt"
	"math"
	"time"

	"github.com/eu-ge-ne/toy2/internal/segbuf"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (ed *Editor) Render() {
	t0 := time.Now()

	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(vt.SaveCursor)
	vt.Buf.Write(ed.colors.background)
	vt.ClearArea(vt.Buf, ed.area)

	ed.determineLayout()

	if ed.area.W >= ed.indexWidth {
		ed.scrollV()
		ed.scrollH()
		ed.renderLines()
	}

	if ed.Enabled {
		vt.SetCursor(vt.Buf, ed.cursorY, ed.cursorX)
		vt.Buf.Write(vt.ShowCursor)
	} else {
		vt.Buf.Write(vt.RestoreCursor)
		vt.Buf.Write(vt.ShowCursor)
	}

	vt.Buf.Flush()

	vt.Sync.Esu()

	if ed.OnCursor != nil {
		ed.OnCursor(ed.cursor.Ln, ed.cursor.Col, ed.Buffer.LineCount())
	}

	if ed.OnRender != nil {
		ed.OnRender(time.Since(t0))
	}
}

func (ed *Editor) determineLayout() {
	if ed.IndexEnabled && ed.Buffer.LineCount() > 0 {
		ed.indexWidth = int(math.Log10(float64(ed.Buffer.LineCount()))) + 3
	} else {
		ed.indexWidth = 0
	}

	ed.textWidth = ed.area.W - ed.indexWidth

	if ed.wrapEnabled {
		ed.Buffer.WrapWidth = ed.textWidth
	} else {
		ed.Buffer.WrapWidth = math.MaxInt
	}

	ed.cursorY = ed.area.Y
	ed.cursorX = ed.area.X + ed.indexWidth

	ed.Buffer.MeasureY = ed.area.Y
	ed.Buffer.MeasureX = ed.area.X + ed.indexWidth
}

func (ed *Editor) scrollV() {
	deltaLn := ed.cursor.Ln - ed.scrollLn

	// Above?
	if deltaLn <= 0 {
		ed.scrollLn = ed.cursor.Ln
		return
	}

	// Below?

	if deltaLn > ed.area.H {
		ed.scrollLn = ed.cursor.Ln - ed.area.H
	}

	xs := make([]int, ed.cursor.Ln+1-ed.scrollLn)
	for i := 0; i < len(xs); i += 1 {
		xs[i] = 1
		for j, cell := range ed.Buffer.Line(ed.scrollLn+i, false) {
			if j > 0 && cell.Col == 0 {
				xs[i] += 1
			}
		}
	}

	i := 0
	height := std.Sum(xs)

	for height > ed.area.H {
		height -= xs[i]
		ed.scrollLn += 1
		i += 1
	}

	for i < len(xs)-1 {
		ed.cursorY += xs[i]
		i += 1
	}
}

func (ed *Editor) scrollH() {
	var cell *segbuf.Seg = nil
	for _, c := range ed.Buffer.LineSlice(ed.cursor.Ln, true, ed.cursor.Col, math.MaxInt) {
		cell = &c
		break
	}
	if cell != nil {
		ed.cursorY += cell.Ln
	}

	col := 0
	if cell != nil {
		col = cell.Col
	}

	deltaCol := col - ed.scrollCol

	// Before?

	if deltaCol <= 0 {
		ed.scrollCol = col
		return
	}

	// After?

	xs := make([]int, deltaCol)
	for i, c := range ed.Buffer.LineSlice(ed.cursor.Ln, true, ed.cursor.Col-deltaCol, ed.cursor.Col) {
		xs[i] = c.G.Width
	}

	width := std.Sum(xs)

	for _, w := range xs {
		if width < ed.textWidth {
			break
		}

		ed.scrollCol += 1
		width -= w
	}

	ed.cursorX += width
}

func (ed *Editor) renderLines() {
	row := ed.area.Y

	for ln := ed.scrollLn; ; ln += 1 {
		if ln < ed.Buffer.LineCount() {
			row = ed.renderLine(ln, row)
		} else {
			vt.SetCursor(vt.Buf, row, ed.area.X)
			vt.Buf.Write(ed.colors.void)
			vt.ClearLine(vt.Buf, ed.area.W)
		}

		row += 1
		if row >= ed.area.Y+ed.area.H {
			break
		}
	}
}

func (ed *Editor) renderLine(ln int, row int) int {
	availableW := 0
	currentColor := charColorUndefined

	for i, cell := range ed.Buffer.Line(ln, false) {
		if cell.Col == 0 {
			if i > 0 {
				row += 1
				if row >= ed.area.Y+ed.area.H {
					return row
				}
			}

			vt.SetCursor(vt.Buf, row, ed.area.X)

			if ed.indexWidth > 0 {
				if i == 0 {
					vt.Buf.Write(ed.colors.index)
					fmt.Fprintf(vt.Buf, "%*d ", ed.indexWidth-1, ln+1)
				} else {
					vt.Buf.Write(ed.colors.background)
					vt.WriteSpaces(vt.Buf, ed.indexWidth)
				}
			}

			availableW = ed.area.W - ed.indexWidth
		}

		if (cell.Col < ed.scrollCol) || (cell.G.Width > availableW) {
			continue
		}

		color := createCharColor(ed.cursor.IsSelected(ln, i), cell.G.IsVisible, ed.WhitespaceEnabled)
		if color != currentColor {
			currentColor = color
			vt.Buf.Write(ed.colors.char[color])
		}

		vt.Buf.Write(cell.G.Bytes)

		availableW -= cell.G.Width
	}

	return row
}
