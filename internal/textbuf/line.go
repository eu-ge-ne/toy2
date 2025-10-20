package textbuf

import (
	"iter"
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
)

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
