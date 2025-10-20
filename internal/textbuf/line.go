package textbuf

import (
	"iter"
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
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

func (buf *TextBuf) ReadLine(ln int) iter.Seq[string] {
	start, ok := buf.LnToByte(ln)
	if !ok {
		return func(yield func(string) bool) {}
	}

	end, ok := buf.LnToByte(ln + 1)
	if !ok {
		end = math.MaxInt
	}

	return buf.Read(start, end)
}

func (buf *TextBuf) LineSegments(ln int) iter.Seq[grapheme.Segment] {
	line := buf.ReadLine(ln)

	return grapheme.Graphemes.Segments(line, false)
}
