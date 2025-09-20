package vt

import "io"

func ClearArea(out io.Writer, a struct{ Y, X, W, H int }) {
	SetCursor(out, a.Y, a.X)

	for i := a.H; i > 0; i -= 1 {
		ECH(out, a.W)
		out.Write(CursorDown)
	}
}

func ClearLine(out io.Writer, w int) {
	ECH(out, w)
}
