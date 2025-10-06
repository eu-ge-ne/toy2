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
		i := 0
		gg := uniseg.NewGraphemes(text)

		for gg.Next() {
			if !yield(i, p.Get(gg.Str())) {
				return
			}

			i += 1
		}
	}
}
