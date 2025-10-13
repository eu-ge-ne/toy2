package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
)

func (buf TextBuf) LnIndex(ln int) (int, bool) {
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

func (buf *TextBuf) ColIndex(ln, col int) (int, bool) {
	index := 0

	for i, cell := range buf.IterLine(ln, true) {
		if i == col {
			return index, true
		}

		index += len(cell.G.Seg)
	}

	return 0, false
}

func (buf *TextBuf) Index(ln, col int) (int, bool) {
	lnIndex, ok := buf.LnIndex(ln)
	if !ok {
		return 0, false
	}

	colIndex, ok := buf.ColIndex(ln, col)
	if !ok {
		return 0, false
	}

	return lnIndex + colIndex, true
}
