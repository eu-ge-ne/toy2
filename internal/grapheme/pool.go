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

func (p GraphemePool) Iter(it iter.Seq[string]) iter.Seq[*Grapheme] {
	return func(yield func(*Grapheme) bool) {
		var seg string

		for text := range it {
			state := -1

			for len(text) > 0 {
				seg, text, _, state = uniseg.StepString(text, state)

				if !yield(p.Get(seg)) {
					return
				}
			}
		}
	}
}

func (p GraphemePool) MeasureString(text string) (ln, col int) {
	var seg string
	state := -1

	for len(text) > 0 {
		seg, text, _, state = uniseg.StepString(text, state)

		if p.Get(seg).IsEol {
			ln += 1
			col = 0
		} else {
			col += 1
		}
	}

	return
}
