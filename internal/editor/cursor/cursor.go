package cursor

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type Cursor struct {
	Ln  int
	Col int

	Selecting bool
	StartLn   int
	StartCol  int
	EndLn     int
	EndCol    int

	buffer     *textbuf.TextBuf
	selFromLn  int
	selFromCol int
}

func New(buffer *textbuf.TextBuf) *Cursor {
	return &Cursor{buffer: buffer}
}

func (cur *Cursor) Set(ln, col int, sel bool) (ok bool) {
	oldLn := cur.Ln
	oldCol := cur.Col

	cur.Ln = std.Clamp(ln, 0, cur.buffer.LineCount()-1)

	if cur.Ln == cur.buffer.LineCount()-1 {
		cur.Col = std.Clamp(col, 0, cur.buffer.ColumnCount(cur.Ln))
	} else {
		cur.Col = std.Clamp(col, 0, cur.buffer.ColumnCount(cur.Ln)-1)
	}

	ok = cur.Ln != oldLn || cur.Col != oldCol

	if !sel {
		cur.Selecting = false
		return
	}

	if !cur.Selecting {
		cur.selFromLn = oldLn
		cur.selFromCol = oldCol
	}

	cur.Selecting = true

	if (cur.selFromLn > cur.Ln) || (cur.selFromLn == cur.Ln && cur.selFromCol > cur.Col) {
		cur.StartLn = cur.Ln
		cur.StartCol = cur.Col
		cur.EndLn = cur.selFromLn
		cur.EndCol = cur.selFromCol
	} else {
		cur.StartLn = cur.selFromLn
		cur.StartCol = cur.selFromCol
		cur.EndLn = cur.Ln
		cur.EndCol = cur.Col
	}

	return
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
