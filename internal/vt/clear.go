package vt

import "io"

func ClearArea(out io.Writer, y, x, w, h int) {
	SetCursor(out, y, x)

	for i := h; i > 0; i -= 1 {
		ECH(out, w)
		out.Write(CursorDown)
	}
}

func ClearLine(out io.Writer, w int) {
	ECH(out, w)
}
