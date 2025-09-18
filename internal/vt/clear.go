package vt

func ClearArea(out out, y, x, w, h int) {
	SetCursor(out, y, x)

	for i := h; i > 0; i -= 1 {
		ECH(out, w)
		out.Write(CursorDown)
	}
}

func ClearLine(out out, w int) {
	ECH(out, w)
}
