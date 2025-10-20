package textbuf

func (buf *TextBuf) PosToStartByte(ln, col int) (int, int, bool) {
	lnByte, ok := buf.LnByte(ln)
	if !ok {
		return 0, 0, false
	}

	colByte, ok := buf.ColByte(ln, col)
	if !ok {
		return 0, 0, false
	}

	return lnByte + colByte, colByte, true
}

func (buf *TextBuf) PosToEndByte(ln, col int) (int, int) {
	maxLn := max(0, buf.LineCount()-1)
	if ln > maxLn {
		ln = maxLn
	}

	lnByte, ok := buf.LnByte(ln)
	if !ok {
		panic("in TextBuf.posToByte2")
	}

	colByte, ok := buf.ColByte(ln, col)
	if !ok {
		colByte = buf.ColMaxByte(ln)
	}

	return lnByte + colByte, colByte
}

func (buf *TextBuf) posToInsertByte(ln, col int) int {
	maxLn := max(0, buf.LineCount()-1)
	if ln > maxLn {
		ln = maxLn
	}

	lnByte, _ := buf.LnByte(ln)

	colByte, ok := buf.ColByte(ln, col)
	if !ok {
		colByte = buf.ColMaxByte(ln)
	}

	return lnByte + colByte
}
