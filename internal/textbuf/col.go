package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/grapheme"
)

func (buf *TextBuf) ColumnCount(ln int) int {
	return grapheme.Graphemes.CountString(buf.ReadLine(ln))
}

func (buf *TextBuf) colIdx(ln, col int) (int, bool) {
	if col == 0 {
		return 0, true
	}

	i := 0
	colIdx := 0

	for gr := range buf.LineGraphemes(ln) {
		if i == col {
			break
		}

		i += 1
		colIdx += len(gr.Str)
	}

	if i == col {
		return colIdx, true
	}

	return 0, false
}

func (buf *TextBuf) maxColIdx(ln int) int {
	colIdx := 0

	for gr := range buf.LineGraphemes(ln) {
		colIdx += len(gr.Str)
	}

	return colIdx
}
