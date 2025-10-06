package grapheme

import (
	"iter"

	"github.com/rivo/uniseg"
)

type GraphemePool struct {
	pool map[string]*Grapheme
}

func (p GraphemePool) Get(seg string) *Grapheme {
	g, ok := p.pool[seg]

	if !ok {
		g = NewGrapheme(seg, []byte(seg), -1)
		p.pool[seg] = g
	}

	return g
}

func (p GraphemePool) IterText(text string) iter.Seq2[int, *Grapheme] {
	return func(yield func(int, *Grapheme) bool) {
		var (
			i     = 0
			c     string
			state = -1
		)

		for len(text) > 0 {
			c, text, _, state = uniseg.StepString(text, state)
			g := p.Get(c)

			if !yield(i, g) {
				return
			}

			i += 1
		}
	}
}

func (p GraphemePool) Iter(it iter.Seq[string]) iter.Seq2[int, *Grapheme] {
	return func(yield func(int, *Grapheme) bool) {
		var (
			i     = 0
			c     string
			state = -1
		)

		for text := range it {
			for len(text) > 0 {
				c, text, _, state = uniseg.StepString(text, state)
				g := p.Get(c)

				if !yield(i, g) {
					return
				}

				i += 1
			}
		}
	}
}
