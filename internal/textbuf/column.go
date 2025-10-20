package textbuf

func (buf *TextBuf) colMaxIdx(ln int) int {
	idx := 0

	for _, gr := range buf.LineGraphemes(ln) {
		idx += len(gr.Str)
	}

	return idx
}

func (buf *TextBuf) colIdx(ln, col int) (int, bool) {
	idx := 0

	for i, gr := range buf.LineGraphemes(ln) {
		if i == col {
			return idx, true
		}

		idx += len(gr.Str)
	}

	return 0, false
}
