package textbuf

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
)

func (buf *TextBuf) ColIndex(ln, col int) (int, bool) {
	index := 0

	line := buf.ReadLine(ln)
	for cell := range grapheme.Graphemes.IterString(line, true, 0, math.MaxInt) {
		if cell.Col == col {
			return index, true
		}

		index += len(cell.Gr.Seg)
	}

	return 0, false
}

func (buf *TextBuf) Index(ln, col int) (int, bool) {
	lnIndex, ok := buf.LnToByte(ln)
	if !ok {
		return 0, false
	}

	colIndex, ok := buf.ColIndex(ln, col)
	if !ok {
		return 0, false
	}

	return lnIndex + colIndex, true
}
