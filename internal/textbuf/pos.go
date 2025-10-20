package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
)

func (buf TextBuf) LnToByte(ln int) (int, bool) {
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

func (buf *TextBuf) LnColToByte(ln, col int) (int, bool) {
	lnIndex, ok := buf.LnToByte(ln)
	if !ok {
		return 0, false
	}

	colIndex, ok := buf.ColToByte(ln, col)
	if !ok {
		return 0, false
	}

	return lnIndex + colIndex, true
}
