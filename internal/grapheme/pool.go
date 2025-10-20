package grapheme

import (
	"iter"

	"github.com/rivo/uniseg"

	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Pool struct {
	pool      map[string]*Grapheme
	wCharY    int
	wCharX    int
	wrapWidth int
}

func (p *Pool) Get(seg string) *Grapheme {
	g, ok := p.pool[seg]

	if !ok {
		g = NewGrapheme(seg, []byte(seg), -1)
		p.pool[seg] = g
	}

	return g
}

func (p *Pool) SetWcharPos(y, x int) {
	p.wCharY = y
	p.wCharX = x
}

func (p *Pool) SetWrapWidth(w int) {
	p.wrapWidth = w
}

type Segment struct {
	Gr      *Grapheme
	Col     int
	WrapLn  int
	WrapCol int
}

func (p *Pool) Segments(it iter.Seq[string], extra bool) iter.Seq[Segment] {
	return func(yield func(Segment) bool) {
		seg := Segment{}
		w := 0
		str := ""

		for text := range it {
			state := -1

			for len(text) > 0 {
				str, text, _, state = uniseg.StepString(text, state)

				seg.Gr = p.Get(str)
				if seg.Gr.Width < 0 {
					seg.Gr.Width = vt.Wchar(p.wCharY, p.wCharX, seg.Gr.Bytes)
				}

				w += seg.Gr.Width
				if w > p.wrapWidth {
					w = seg.Gr.Width
					seg.WrapLn += 1
					seg.WrapCol = 0
				}

				if !yield(seg) {
					return
				}

				seg.Col += 1
				seg.WrapCol += 1
			}
		}

		if extra {
			seg.Gr = p.Get(" ")

			w += seg.Gr.Width
			if w > p.wrapWidth {
				w = seg.Gr.Width
				seg.WrapLn += 1
				seg.WrapCol = 0
			}

			if !yield(seg) {
				return
			}
		}
	}
}

func (p *Pool) MeasureString(text string) (ln, col int) {
	str := ""
	state := -1

	for len(text) > 0 {
		str, text, _, state = uniseg.StepString(text, state)

		if p.Get(str).IsEol {
			ln += 1
			col = 0
		} else {
			col += 1
		}
	}

	return
}
