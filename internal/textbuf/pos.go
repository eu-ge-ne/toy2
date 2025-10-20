package textbuf

type Pos struct {
	Idx int
	Ln  int
	Col int
}

func (buf *TextBuf) StartPos(ln, col int) (Pos, bool) {
	lnIdx, ok := buf.LnIdx(ln)
	if !ok {
		return Pos{}, false
	}

	colIdx, ok := buf.colIdx(ln, col)
	if !ok {
		return Pos{}, false
	}

	return Pos{lnIdx + colIdx, ln, colIdx}, true
}

func (buf *TextBuf) EndPos(ln, col int) Pos {
	maxLn := max(0, buf.LineCount()-1)
	if ln > maxLn {
		ln = maxLn
	}

	lnIdx, _ := buf.LnIdx(ln)

	colIdx, ok := buf.colIdx(ln, col)
	if !ok {
		colIdx = buf.colMaxIdx(ln)
	}

	return Pos{lnIdx + colIdx, ln, colIdx}
}
