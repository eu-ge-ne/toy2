package textbuf

func (buf *TextBuf) ColMaxByte(ln int) int {
	n := 0

	for _, gr := range buf.LineGraphemes(ln) {
		n += len(gr.Str)
	}

	return n
}

func (buf *TextBuf) ColByte(ln, col int) (int, bool) {
	index := 0

	for i, gr := range buf.LineGraphemes(ln) {
		if i == col {
			return index, true
		}

		index += len(gr.Str)
	}

	return 0, false
}
