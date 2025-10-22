package textbuf

type Pos struct {
	Ln     int
	Col    int
	Idx    int
	ColIdx int
}

func (buf *TextBuf) StartPos(ln, col int) (Pos, bool) {
	lnIdx, ok := buf.lnIdx(ln)
	if !ok {
		return Pos{}, false
	}

	colIdx, ok := buf.colIdx(ln, col)
	if !ok {
		return Pos{}, false
	}

	return Pos{Ln: ln, Col: col, Idx: lnIdx + colIdx, ColIdx: colIdx}, true
}

func (buf *TextBuf) EndPos(ln, col int) Pos {
	maxLn := max(0, buf.LineCount()-1)
	if ln > maxLn {
		ln = maxLn
	}

	lnIdx, _ := buf.lnIdx(ln)

	colIdx, ok := buf.colIdx(ln, col)
	if !ok {
		colIdx = buf.maxColIdx(ln)
		col = max(0, buf.ColumnCount(ln)-1)
	}

	return Pos{Ln: ln, Col: col, Idx: lnIdx + colIdx, ColIdx: colIdx}
}
