package vt

func ClearArea(y, x, w, h int) []byte {
	b := SetCursor(y, x)

	for i := h; i > 0; i -= 1 {
		b = append(b, ECH(w)...)
		b = append(b, CursorDown...)
	}

	return b
}

func ClearLine(w int) []byte {
	return ECH(w)
}
