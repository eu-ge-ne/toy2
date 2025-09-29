package segbuf

import (
	"iter"
	"math"
	"slices"
	"strings"

	"github.com/rivo/uniseg"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type SegBuf struct {
	textbuf   textbuf.TextBuf
	WrapWidth int
	MeasureY  int
	MeasureX  int
}

type Snapshot = textbuf.Snapshot

func New() *SegBuf {
	return &SegBuf{
		textbuf:   textbuf.Create(),
		WrapWidth: math.MaxInt,
	}
}

func (sb *SegBuf) LineCount() int {
	return sb.textbuf.LineCount()
}

func (sb *SegBuf) Save() Snapshot {
	return sb.textbuf.Save()
}

func (sb *SegBuf) Restore(s Snapshot) {
	sb.textbuf.Restore(s)
}

func (sb *SegBuf) Reset(text string) {
	sb.textbuf.Reset(text)
}

func (sb *SegBuf) Append(text string) {
	sb.textbuf.Append(text)
}

func (sb *SegBuf) Iter() iter.Seq[string] {
	return sb.textbuf.Read(0)
}

func (sb *SegBuf) Text() string {
	return strings.Join(slices.Collect(sb.textbuf.Read(0)), "")
}

type Cell struct {
	G   *grapheme.Grapheme
	Ln  int
	Col int
}

func (sb *SegBuf) Line(ln int, extra bool) iter.Seq2[int, Cell] {
	return func(yield func(int, Cell) bool) {
		i := 0
		c := Cell{}
		w := 0

		for seg := range sb.segLine(ln) {
			c.G = grapheme.Graphemes.Get(seg)

			if c.G.Width < 0 {
				c.G.Width = vt.MeasureCursor(sb.MeasureY, sb.MeasureX, c.G.Bytes)
			}

			w += c.G.Width
			if w > sb.WrapWidth {
				w = c.G.Width
				c.Ln += 1
				c.Col = 0
			}

			if !yield(i, c) {
				return
			}

			i += 1
			c.Col += 1
		}

		if extra {
			c.G = grapheme.Graphemes.Get(" ")

			w += c.G.Width
			if w > sb.WrapWidth {
				w = c.G.Width
				c.Ln += 1
				c.Col = 0
			}

			if !yield(i, c) {
				return
			}
		}
	}
}

func (sb *SegBuf) LineSlice(ln int, extra bool, start, end int) iter.Seq2[int, Cell] {
	return func(yield func(int, Cell) bool) {
		i := 0
		for j, c := range sb.Line(ln, extra) {
			if j >= start && j < end {
				if !yield(i, c) {
					return
				}
				i += 1
			}
		}
	}
}

func (sb *SegBuf) Read(startLn, startCol, endLn, endCol int) string {
	startCol = sb.col(startLn, startCol)
	endCol = sb.col(endLn, endCol)

	return strings.Join(slices.Collect(sb.textbuf.Read2Range(startLn, startCol, endLn, endCol)), "")
}

func (sb *SegBuf) Insert(ln, col int, text string) {
	col = sb.col(ln, col)

	sb.textbuf.Insert2(ln, col, text)
}

func (sb *SegBuf) Delete(startLn, startCol, endLn, endCol int) {
	startCol = sb.col(startLn, startCol)
	endCol = sb.col(endLn, endCol)

	sb.textbuf.Delete2Range(startLn, startCol, endLn, endCol)
}

func (sb *SegBuf) col(ln, col int) int {
	col2 := 0
	i := 0

	for seg := range sb.segLine(ln) {
		if i == col {
			break
		}

		if i < col {
			col2 += len(seg)
		}

		i += 1
	}

	return col2
}

func (sb *SegBuf) segLine(ln int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for chunk := range sb.textbuf.Read2Range(ln, 0, ln+1, 0) {
			gr := uniseg.NewGraphemes(chunk)

			for gr.Next() {
				if !yield(gr.Str()) {
					return
				}
			}
		}
	}
}
