package cursor

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type Cursor struct {
	Ln        int
	Col       int
	Selecting bool
	StartLn   int
	StartCol  int
	EndLn     int
	EndCol    int

	buffer *textbuf.TextBuf
	ln0    int
	col0   int
}

func New(buffer *textbuf.TextBuf) *Cursor {
	return &Cursor{buffer: buffer}
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

func (cur *Cursor) Forward(dLn, dCol int) bool {
	if dLn == 0 {
		return cur.Set(cur.Ln, cur.Col+dCol, false)
	} else {
		return cur.Set(cur.Ln+dLn, dCol, false)
	}
}

func (cur *Cursor) IsSelected(ln, col int) bool {
	if !cur.Selecting {
		return false
	}

	if ln < cur.StartLn || ln > cur.EndLn {
		return false
	}

	if ln == cur.StartLn && col < cur.StartCol {
		return false
	}

	if ln == cur.EndLn && col >= cur.EndCol {
		return false
	}

	return true
}

func (cur *Cursor) Set(ln, col int, sel bool) bool {
	oldLn := cur.Ln
	oldCol := cur.Col

	cur.setLn(ln)
	cur.setCol(col)
	cur.setSelection(oldLn, oldCol, sel)

	return cur.Ln != oldLn || cur.Col != oldCol
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

	for cell := range cur.buffer.IterLine(cur.Ln, false) {
		if cell.Gr.IsEol {
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

	if (cur.ln0 > cur.Ln) || (cur.ln0 == cur.Ln && cur.col0 > cur.Col) {
		cur.StartLn = cur.Ln
		cur.StartCol = cur.Col
		cur.EndLn = cur.ln0
		cur.EndCol = cur.col0
	} else {
		cur.StartLn = cur.ln0
		cur.StartCol = cur.col0
		cur.EndLn = cur.Ln
		cur.EndCol = cur.Col
	}
}
