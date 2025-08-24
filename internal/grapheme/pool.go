package grapheme

type GraphemePool struct {
	pool map[string]*Grapheme
}

func NewGraphemePool(pool map[string]*Grapheme) *GraphemePool {
	return &GraphemePool{pool: pool}
}

func (p *GraphemePool) Get(seg string) *Grapheme {
	g, ok := p.pool[seg]

	if !ok {
		g = NewGrapheme(seg, []byte(seg), -1)
		p.pool[seg] = g
	}

	return g
}
