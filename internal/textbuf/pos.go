package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf/internal/node"
)

func (buf TextBuf) posToIndex(ln, col int) (int, bool) {
	i, ok := buf.findLineStart(ln)

	if !ok {
		return 0, false
	}

	return i + col, true
}

func (buf TextBuf) findLineStart(ln int) (int, bool) {
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
			buf := buf.content.Buffers[x.BufIndex]
			eolEnd := buf.Eols[x.EolsStart+eolIndex].End
			return i + eolEnd - x.Start, true
		}

		eolIndex -= x.EolsLen
		i += x.Len
		x = x.Right
	}

	return 0, false
}
