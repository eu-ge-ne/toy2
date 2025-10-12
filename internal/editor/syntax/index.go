package syntax

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func index2(buf *textbuf.TextBuf, startLn, startCol, endLn, endCol int) (int, int, bool) {
	start, ok := buf.Index(startLn, startCol)
	if !ok {
		return 0, 0, false
	}

	end, ok := buf.Index(endLn, endCol)
	if !ok {
		return 0, 0, false
	}

	return start, end, true
}
