package grapheme

import (
	"iter"

	"github.com/rivo/uniseg"

	"github.com/eu-ge-ne/toy2/internal/vt"
)

type GraphemePool struct {
	pool      map[string]*Grapheme
	wcharY    int
	wcharX    int
	wrapWidth int
}

func (p *GraphemePool) Get(seg string) *Grapheme {
	g, ok := p.pool[seg]

	if !ok {
		g = NewGrapheme(seg, []byte(seg), -1)
		p.pool[seg] = g
	}

	return g
}

func (p *GraphemePool) SetWcharPos(y, x int) {
	p.wcharY = y
	p.wcharX = x
}

func (p *GraphemePool) SetWrapWidth(wrapWidth int) {
	p.wrapWidth = wrapWidth
}

type IterCell struct {
	Gr      *Grapheme
	Col     int
	WrapLn  int
	WrapCol int
}

func (p *GraphemePool) IterString(it iter.Seq[string], extra bool) iter.Seq[IterCell] {
	return func(yield func(IterCell) bool) {
		cell := IterCell{}
		w := 0
		seg := ""

		for text := range it {
			state := -1

			for len(text) > 0 {
				seg, text, _, state = uniseg.StepString(text, state)

				cell.Gr = p.Get(seg)

				if cell.Gr.Width < 0 {
					cell.Gr.Width = vt.Wchar(p.wcharY, p.wcharX, cell.Gr.Bytes)
				}

				w += cell.Gr.Width
				if w > p.wrapWidth {
					w = cell.Gr.Width
					cell.WrapLn += 1
					cell.WrapCol = 0
				}

				if !yield(cell) {
					return
				}

				cell.Col += 1
				cell.WrapCol += 1
			}
		}

		if extra {
			cell.Gr = p.Get(" ")

			w += cell.Gr.Width
			if w > p.wrapWidth {
				w = cell.Gr.Width
				cell.WrapLn += 1
				cell.WrapCol = 0
			}

			if !yield(cell) {
				return
			}
		}
	}
}

func (p *GraphemePool) MeasureString(text string) (ln, col int) {
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
