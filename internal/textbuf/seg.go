package textbuf

import (
	"iter"

	"github.com/rivo/uniseg"
)

func (tb *TextBuf) SegLine(ln int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for chunk := range tb.Read2Range(ln, 0, ln+1, 0) {
			gr := uniseg.NewGraphemes(chunk)

			for gr.Next() {
				if !yield(gr.Str()) {
					return
				}
			}
		}
	}
}

func (tb *TextBuf) SegRead2(startLn, startCol, endLn, endCol int) iter.Seq[string] {
	startLn, startCol = tb.segToPos(startLn, startCol)
	endLn, endCol = tb.segToPos(endLn, endCol)

	return tb.Read2Range(startLn, startCol, endLn, endCol)
}

func (tb *TextBuf) SegInsert2(ln, col int, text string) {
	ln, col = tb.segToPos(ln, col)

	tb.Insert2(ln, col, text)
}

func (tb *TextBuf) SegDelete2(startLn, startCol, endLn, endCol int) {
	startLn, startCol = tb.segToPos(startLn, startCol)
	endLn, endCol = tb.segToPos(endLn, endCol)

	tb.Delete2Range(startLn, startCol, endLn, endCol)
}

func (tb *TextBuf) segToPos(ln, col int) (int, int) {
	col2 := 0
	i := 0

	for seg := range tb.SegLine(ln) {
		if i == col {
			break
		}

		if i < col {
			col2 += len(seg)
		}

		i += 1
	}

	return ln, col2
}
