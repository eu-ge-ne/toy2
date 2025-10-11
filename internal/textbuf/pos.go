package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
)

type Pos struct {
	ln  int
	col int
}

func (buf TextBuf) lnIndex(pos Pos) (int, bool) {
	if buf.Count() == 0 {
		return 0, false
	}

	if pos.ln == 0 {
		return 0, true
	}

	eolIndex := pos.ln - 1
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

func (buf *TextBuf) colIndex(pos Pos, lnIndex int) (int, bool) {
	colIndex := 0

	for i, cell := range buf.IterLine(pos.ln, true) {
		if i == pos.col {
			return lnIndex + colIndex, true
		}

		colIndex += len(cell.G.Seg)
	}

	return 0, false
}

func (buf *TextBuf) lnColIndex(pos Pos) (int, bool) {
	lnIndex, ok := buf.lnIndex(pos)
	if !ok {
		return 0, false
	}

	return buf.colIndex(pos, lnIndex)
}

func (buf *TextBuf) Index2(startLn, startCol, endLn, endCol int) (int, int, bool) {
	start, ok := buf.lnColIndex(Pos{startLn, startCol})
	if !ok {
		return 0, 0, false
	}

	end, ok := buf.lnColIndex(Pos{endLn, endCol})
	if !ok {
		return 0, 0, false
	}

	return start, end, true
}
