package piece

type Piece struct {
	data []byte
	Len  int
	Eols []Eol
}

type Eol struct {
	Start int
	End   int
}

func Create(data []byte) Piece {
	x := Piece{}

	x.appendEols(data)
	x.data = data
	x.Len = len(data)

	return x
}

func (p *Piece) Append(data []byte) {
	p.appendEols(data)

	p.data = append(p.data, data...)
	p.Len += len(data)
}

func (p *Piece) Read(start int, end int) []byte {
	return p.data[start:end]
}

func (p *Piece) FindEolIndex(index int, a int) int {
	b := len(p.Eols) - 1

outer:
	for a <= b {
		i := (a + b) / 2
		start := p.Eols[i].Start
		end := p.Eols[i].End

		switch {
		case index >= end:
			a = i + 1
		case index < start:
			b = i - 1
		case index == start:
			a = i
			break outer
		default:
			panic("Invalid arguments")
		}
	}

	return a
}

func (p *Piece) appendEols(data []byte) {
	var pr byte

	for i, r := range data {
		switch {
		case pr == '\r' && r == '\n':
			a := p.Len + i - 1
			b := a + 2
			p.Eols = append(p.Eols, Eol{Start: a, End: b})
		case r == '\n':
			a := p.Len + i
			b := a + 1
			p.Eols = append(p.Eols, Eol{Start: a, End: b})
		}
		pr = r
	}

	if pr == '\r' {
		a := p.Len + len(data) - 1
		b := a + 1
		p.Eols = append(p.Eols, Eol{Start: a, End: b})
	}
}
