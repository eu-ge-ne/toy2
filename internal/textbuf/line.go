package textbuf

import (
	"iter"
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
)

func (buf *TextBuf) IterLine(ln int, extra bool) iter.Seq[grapheme.IterCell] {
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
	}

	return grapheme.Graphemes.IterString(buf.Read(start, end), opts)
}

func (buf *TextBuf) IterLine2(ln int, extra bool, start, end int) iter.Seq[grapheme.IterCell] {
	return func(yield func(grapheme.IterCell) bool) {
		for cell := range buf.IterLine(ln, extra) {
			if cell.Col >= start && cell.Col < end {
				if !yield(cell) {
					return
				}
			}
		}
	}
}
