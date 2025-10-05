package cursor

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/std"
)

type Cursor struct {
	ln0       int
	col0      int
	Ln        int
	Col       int
	Selecting bool
	FromLn    int
	FromCol   int
	ToLn      int
	ToCol     int

	buffer *textbuf.TextBuf
}

func New(buffer *textbuf.TextBuf) Cursor {
	return Cursor{buffer: buffer}
}

func (cur *Cursor) Set(ln, col int, sel bool) bool {
	oldLn := cur.Ln
	oldCol := cur.Col

	cur.setLn(ln)
	cur.setCol(col)
	cur.setSelection(oldLn, oldCol, sel)
	cur.setRange()

	return cur.Ln != oldLn || cur.Col != oldCol
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

	for _, c := range cur.buffer.IterSegLine(cur.Ln, false) {
		if c.G.IsEol {
			break
		}
		len += 1
	}

	cur.Col = std.Clamp(col, 0, len)
}

func (cur *Cursor) setSelection(ln, col int, sel bool) {
	if !sel {
		cur.Selecting = false
		return
	}

	if !cur.Selecting {
		cur.ln0 = ln
		cur.col0 = col
	}

	cur.Selecting = true
}

func (cur *Cursor) setRange() {
	if (cur.ln0 > cur.Ln) || (cur.ln0 == cur.Ln && cur.col0 > cur.Col) {
		cur.FromLn = cur.Ln
		cur.FromCol = cur.Col
		cur.ToLn = cur.ln0
		cur.ToCol = cur.col0
	} else {
		cur.FromLn = cur.ln0
		cur.FromCol = cur.col0
		cur.ToLn = cur.Ln
		cur.ToCol = cur.Col
	}
}
