package cursor

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type Cursor struct {
	Ln  int
	Col int

	Selecting bool
	FromLn    int
	FromCol   int
	ToLn      int
	ToCol     int
	startLn   int
	startCol  int

	buffer *textbuf.TextBuf
}

func New(buffer *textbuf.TextBuf) *Cursor {
	return &Cursor{buffer: buffer}
}

func (cur *Cursor) Set(ln, col int, sel bool) bool {
	oldLn := cur.Ln
	oldCol := cur.Col

	cur.setLn(ln)
	cur.setCol(col)
	cur.setRange(oldLn, oldCol, cur.Ln, cur.Col, sel)

	return cur.Ln != oldLn || cur.Col != oldCol
}

func (cur *Cursor) Move(dy, dx int, sel bool) bool {
	return cur.Set(cur.Ln+dy, cur.Col+dx, sel)
}

func (cur *Cursor) Top(sel bool) bool {
	return cur.Set(0, 0, sel)
}

func (cur *Cursor) Bottom(sel bool) bool {
	return cur.Set(math.MaxInt, 0, sel)
}

func (cur *Cursor) Home(sel bool) bool {
	return cur.Set(cur.Ln, 0, sel)
}

func (cur *Cursor) End(sel bool) bool {
	return cur.Set(cur.Ln, math.MaxInt, sel)
}

func (cur *Cursor) Up(n int, sel bool) bool {
	return cur.Set(cur.Ln-n, cur.Col, sel)
}

func (cur *Cursor) Down(n int, sel bool) bool {
	return cur.Set(cur.Ln+n, cur.Col, sel)
}

func (cur *Cursor) Left(sel bool) bool {
	if cur.Set(cur.Ln, cur.Col-1, sel) {
		return true
	}

	if cur.Ln > 0 {
		return cur.Set(cur.Ln-1, math.MaxInt, sel)
	}

	return false
}

func (cur *Cursor) Right(sel bool) bool {
	if cur.Set(cur.Ln, cur.Col+1, sel) {
		return true
	}

	if cur.Ln < cur.buffer.LineCount()-1 {
		return cur.Set(cur.Ln+1, 0, sel)
	}

	return false
}

func (cur *Cursor) Forward(n int) bool {
	return cur.Set(cur.Ln, cur.Col+n, false)
}

func (cur *Cursor) IsSelected(ln, col int) bool {
	if !cur.Selecting {
		return false
	}

	if ln < cur.FromLn || ln > cur.ToLn {
		return false
	}

	if ln == cur.FromLn && col < cur.FromCol {
		return false
	}

	if ln == cur.ToLn && col > cur.ToCol {
		return false
	}

	return true
}

func (cur *Cursor) setLn(ln int) {
	max := cur.buffer.LineCount() - 1
	if max < 0 {
		max = 0
	}

	cur.Ln = std.Clamp(ln, 0, max)
}

func (cur *Cursor) setCol(col int) {
	len := 0

	for seg := range cur.buffer.SegLine(cur.Ln) {
		g := grapheme.Graphemes.Get(seg)
		if g.IsEol {
			break
		}
		len += 1
	}

	cur.Col = std.Clamp(col, 0, len)
}

func (cur *Cursor) setRange(oldLn, oldCol, newLn, newCol int, sel bool) {
	if !sel {
		cur.Selecting = false
		return
	}

	if !cur.Selecting {
		cur.startLn = oldLn
		cur.startCol = oldCol
	}

	cur.Selecting = true

	if (cur.startLn > newLn) || (cur.startLn == newLn && cur.startCol > newCol) {
		cur.FromLn = newLn
		cur.FromCol = newCol

		cur.ToLn = cur.startLn
		cur.ToCol = cur.startCol
	} else {
		cur.FromLn = cur.startLn
		cur.FromCol = cur.startCol

		cur.ToLn = newLn
		cur.ToCol = newCol
	}
}
