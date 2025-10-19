package textbuf

import (
	"iter"
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
)

func (buf *TextBuf) IterLine(ln int, extra bool, iterStart, iterEnd int) iter.Seq[grapheme.IterCell] {
	start, ok := buf.LnIndex(ln)
	if !ok {
		return func(yield func(grapheme.IterCell) bool) {}
	}

	end, ok := buf.LnIndex(ln + 1)
	if !ok {
		end = math.MaxInt
	}

	opts := grapheme.IterOptions{
		WcharY:    buf.MeasureY,
		WcharX:    buf.MeasureX,
		WrapWidth: buf.WrapWidth,
		Extra:     extra,
		Start:     iterStart,
		End:       iterEnd,
	}

	return grapheme.Graphemes.IterString(buf.Read(start, end), opts)
}
