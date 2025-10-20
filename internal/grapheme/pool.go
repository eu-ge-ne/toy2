package grapheme

import (
	"iter"

	"github.com/rivo/uniseg"

	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Pool struct {
	pool   map[string]*Grapheme
	wCharY int
	wCharX int
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

func (p *Pool) FromString(it iter.Seq[string]) iter.Seq2[int, *Grapheme] {
	return func(yield func(int, *Grapheme) bool) {
		i := 0
		str := ""

		for text := range it {
			state := -1

			for len(text) > 0 {
				str, text, _, state = uniseg.StepString(text, state)

				gr := p.Get(str)
				if gr.Width < 0 {
					gr.Width = vt.Wchar(p.wCharY, p.wCharX, gr.Bytes)
				}

				if !yield(i, gr) {
					return
				}

				i += 1
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
