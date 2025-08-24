package editor

import (
	"iter"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type cell struct {
	g   *grapheme.Grapheme
	ln  int
	col int
}

func (ed *Editor) cells(lineIndex int, withTail bool) iter.Seq2[int, cell] {
	return func(yield func(int, cell) bool) {
		i := 0
		c := cell{}
		w := 0

		for seg := range ed.Buffer.SegLine(lineIndex) {
			c.g = grapheme.Graphemes.Get(seg)

			if c.g.Width < 0 {
				c.g.Width = vt.MeasureCursor(ed.measureY, ed.measureX, c.g.Bytes)
			}

			w += c.g.Width
			if w > ed.wrapWidth {
				w = c.g.Width
				c.ln += 1
				c.col = 0
			}

			if !yield(i, c) {
				return
			}

			i += 1
			c.col += 1
		}

		if withTail {
			c.g = grapheme.Graphemes.Get(" ")

			w += c.g.Width
			if w > ed.wrapWidth {
				w = c.g.Width
				c.ln += 1
				c.col = 0
			}

			if !yield(i, c) {
				return
			}
		}
	}
}

func (ed *Editor) sliceCells(lineIndex int, withTail bool, start, end int) iter.Seq2[int, cell] {
	return func(yield func(int, cell) bool) {
		i := 0
		for j, c := range ed.cells(lineIndex, withTail) {
			if j >= start && j < end {
				if !yield(i, c) {
					return
				}
				i += 1
			}
		}
	}
}
