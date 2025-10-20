package textbuf

import (
	"github.com/eu-ge-ne/toy2/internal/grapheme"
)

func (buf *TextBuf) Index(ln, col int) (int, bool) {
	lnIndex, ok := buf.LnToByte(ln)
	if !ok {
		return 0, false
	}

	line := buf.ReadLine(ln)
	colIndex, ok := grapheme.Graphemes.ColToByte(line, col)
	if !ok {
		return 0, false
	}

	return lnIndex + colIndex, true
}
