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

func (buf *Piece) Append(text string) {
	buf.appendEols(text)

	buf.text += text
	buf.Len += len(text)
}

func (buf *Piece) Read(start int, end int) string {
	return buf.text[start:end]
}

func (buf *Piece) FindEolIndex(index int, a int) int {
	b := len(buf.Eols) - 1

outer:
	for a <= b {
		i := (a + b) / 2
		start := buf.Eols[i].Start
		end := buf.Eols[i].End

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

func (buf *Piece) appendEols(text string) {
	var p rune

	for i, r := range text {
		switch {
		case p == '\r' && r == '\n':
			a := buf.Len + i - 1
			b := a + 2
			buf.Eols = append(buf.Eols, Eol{Start: a, End: b})
		case r == '\n':
			a := buf.Len + i
			b := a + 1
			buf.Eols = append(buf.Eols, Eol{Start: a, End: b})
		}
		p = r
	}

	if p == '\r' {
		a := buf.Len + len(text) - 1
		b := a + 1
		buf.Eols = append(buf.Eols, Eol{Start: a, End: b})
	}
}
