package textbuf

func (buf *TextBuf) ColMax(ln int) int {
	col := 0

	for range buf.LineGraphemes(ln) {
		col += 1
	}

	return col
}

func (buf *TextBuf) ColMaxNonEol(ln int) int {
	col := 0

	for _, gr := range buf.LineGraphemes(ln) {
		if gr.IsEol {
			break
		}
		col += 1
	}

	return col
}

func (buf *TextBuf) colIdx(ln, col int) (int, bool) {
	if col == 0 {
		return 0, true
	}

	colIdx := 0

	for i, gr := range buf.LineGraphemes(ln) {
		if i == col {
			return colIdx, true
		}

		colIdx += len(gr.Str)
	}

	return 0, false
}

func (buf *TextBuf) colIdxMax(ln int) int {
	colIdx := 0

	for _, gr := range buf.LineGraphemes(ln) {
		colIdx += len(gr.Str)
	}

	return colIdx
}
