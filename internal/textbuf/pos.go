package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
)

type Pos struct {
	Idx int
	Ln  int
	Col int
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

	return Pos{lnIdx + colIdx, ln, colIdx}, true
}

func (buf *TextBuf) PosMax(ln, col int) Pos {
	maxLn := max(0, buf.LineCount()-1)
	if ln > maxLn {
		ln = maxLn
	}

	lnIdx, _ := buf.lnIdx(ln)

	colIdx, ok := buf.colIdx(ln, col)
	if !ok {
		colIdx = buf.colIdxMax(ln)
	}

	return Pos{lnIdx + colIdx, ln, colIdx}
}

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

func (buf *TextBuf) lnIdx(ln int) (int, bool) {
	if buf.Count() == 0 {
		return 0, false
	}

	if ln == 0 {
		return 0, true
	}

	eolIndex := ln - 1
	x := buf.tree.Root
	i := 0

	for x != node.NIL {
		if eolIndex < x.Left.TotalEolsLen {
			x = x.Left
			continue
		}

		eolIndex -= x.Left.TotalEolsLen
		i += x.Left.TotalLen

		if eolIndex < x.EolsLen {
			buf := buf.content.Table[x.PieceIndex]
			eolEnd := buf.Eols[x.EolsStart+eolIndex].End
			return i + eolEnd - x.Start, true
		}

		eolIndex -= x.EolsLen
		i += x.Len
		x = x.Right
	}

	return 0, false
}

func (buf *TextBuf) colIdx(ln, col int) (int, bool) {
	if col == 0 {
		return 0, true
	}

	idx := 0

	for i, gr := range buf.LineGraphemes(ln) {
		if i == col {
			return idx, true
		}

		idx += len(gr.Str)
	}

	return 0, false
}

func (buf *TextBuf) colIdxMax(ln int) int {
	idx := 0

	for _, gr := range buf.LineGraphemes(ln) {
		idx += len(gr.Str)
	}

	return idx
}
