package piece

type Piece struct {
	text string
	Len  int
	Eols []Eol
}

type Eol struct {
	Start int
	End   int
}

func Create(text string) Piece {
	x := Piece{}

	x.appendEols(text)
	x.text = text
	x.Len = len(text)

	return x
}

func (p *Piece) Append(text string) {
	p.appendEols(text)

	p.text += text
	p.Len += len(text)
}

func (p *Piece) Read(start int, end int) string {
	return p.text[start:end]
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

func (p *Piece) appendEols(text string) {
	var pr rune

	for i, r := range text {
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
		a := p.Len + len(text) - 1
		b := a + 1
		p.Eols = append(p.Eols, Eol{Start: a, End: b})
	}
}
