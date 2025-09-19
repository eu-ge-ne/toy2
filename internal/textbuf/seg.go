package textbuf

import (
	"iter"

	"github.com/rivo/uniseg"
)

func (buf *TextBuf) SegLine(ln int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for chunk := range buf.Read2(ln, 0, ln+1, 0) {
			gr := uniseg.NewGraphemes(chunk)

			for gr.Next() {
				if !yield(gr.Str()) {
					return
				}
			}
		}
	}
}

func (buf *TextBuf) SegRead2(startLn, startCol, endLn, endCol int) iter.Seq[string] {
	startLn, startCol = buf.segToPos(startLn, startCol)
	endLn, endCol = buf.segToPos(endLn, endCol)

	return buf.Read2(startLn, startCol, endLn, endCol)
}

func (buf *TextBuf) SegInsert2(ln, col int, text string) {
	ln, col = buf.segToPos(ln, col)

	buf.Insert2(ln, col, text)
}

func (buf *TextBuf) SegDelete2(startLn, startCol, endLn, endCol int) {
	startLn, startCol = buf.segToPos(startLn, startCol)
	endLn, endCol = buf.segToPos(endLn, endCol)

	buf.Delete2(startLn, startCol, endLn, endCol)
}

func (buf *TextBuf) segToPos(ln, col int) (int, int) {
	col2 := 0
	i := 0

	for seg := range buf.SegLine(ln) {
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
