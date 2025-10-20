package textbuf

func (buf *TextBuf) ColToByte(ln, col int) (int, bool) {
	index := 0

	for cell := range buf.LineSegments(ln) {
		if cell.I == col {
			return index, true
		}

		index += len(cell.Gr.Str)
	}

	return 0, false
}
