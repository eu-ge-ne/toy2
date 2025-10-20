package textbuf

func (buf *TextBuf) colToByte(ln, col int) (int, bool) {
	index := 0

	for i, gr := range buf.LineGraphemes(ln) {
		if i == col {
			return index, true
		}

		index += len(gr.Str)
	}

	return 0, false
}
