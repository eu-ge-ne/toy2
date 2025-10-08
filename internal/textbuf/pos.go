package textbuf

import "github.com/eu-ge-ne/toy2/internal/std"

type Pos struct {
	Ln     int
	Col    int
	Idx    int
	ColIdx int
}

func (buf *TextBuf) Pos(ln, col int) (Pos, bool) {
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
	ln = std.Clamp(ln, 0, buf.LineCount()-1)

	lnIdx, _ := buf.lnIdx(ln)

	colIdx, ok := buf.colIdx(ln, col)
	if !ok {
		colIdx = buf.endColIdx(ln)
		col = max(0, buf.ColumnCount(ln)-1)
	}

	return Pos{Ln: ln, Col: col, Idx: lnIdx + colIdx, ColIdx: colIdx}
}
