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

func (p GraphemePool) IterText(text string) iter.Seq[*Grapheme] {
	return func(yield func(*Grapheme) bool) {
		var (
			state = -1
			seg   string
		)

		for len(text) > 0 {
			seg, text, _, state = uniseg.StepString(text, state)

			if !yield(p.Get(seg)) {
				return
			}

		}
	}
}

func (p GraphemePool) Iter(it iter.Seq[string]) iter.Seq[*Grapheme] {
	return func(yield func(*Grapheme) bool) {
		var (
			state = -1
			seg   string
		)

		for text := range it {
			for len(text) > 0 {
				seg, text, _, state = uniseg.StepString(text, state)

				if !yield(p.Get(seg)) {
					return
				}
			}
		}
	}
}
