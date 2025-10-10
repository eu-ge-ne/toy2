package editor

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (ed *Editor) Backspace() bool {
	if ed.cursor.Selecting {
		ed.deleteSelection()
	} else {
		ed.deletePrevChar()
	}

	return true
}

func (ed *Editor) Bottom(sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Bottom(sel)
}

func (ed *Editor) Copy() bool {
	if !ed.enabled {
		return false
	}

	cur := ed.cursor

	if cur.Selecting {
		ed.clipboard = ed.buffer.Read2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		cur.Set(cur.Ln, cur.Col, false)
	} else {
		ed.clipboard = ed.buffer.Read2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
	}

	vt.CopyToClipboard(vt.Sync, ed.clipboard)

	return false
}

func (ed *Editor) Cut() bool {
	if !ed.enabled {
		return false
	}

	cur := ed.cursor

	if cur.Selecting {
		ed.clipboard = ed.buffer.Read2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		ed.deleteSelection()
	} else {
		ed.clipboard = ed.buffer.Read2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
		ed.deleteChar()
	}

	vt.CopyToClipboard(vt.Sync, ed.clipboard)

	return true
}

func (ed *Editor) Delete() bool {
	if ed.cursor.Selecting {
		ed.deleteSelection()
	} else {
		ed.deleteChar()
	}

	return true
}

func (ed *Editor) Down(sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Down(1, sel)
}

func (ed *Editor) End(sel bool) bool {
	return ed.cursor.End(sel)
}

func (ed *Editor) Enter() bool {
	if !ed.multiLine {
		return false
	}

	return ed.Insert("\n")
}

func (ed *Editor) Home(sel bool) bool {
	return ed.cursor.Home(sel)
}

func (ed *Editor) Insert(text string) bool {
	ed.insertText(text)

	return true
}

func (ed *Editor) Left(sel bool) bool {
	return ed.cursor.Left(sel)
}

func (ed *Editor) PageDown(sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Down(ed.pageSize, sel)
}

func (ed *Editor) PageUp(sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Up(ed.pageSize, sel)
}

func (ed *Editor) Paste() bool {
	if !ed.enabled {
		return false
	}

	if len(ed.clipboard) == 0 {
		return false
	}

	return ed.Insert(ed.clipboard)
}

func (ed *Editor) Redo() bool {
	if !ed.enabled {
		return false
	}

	return ed.history.Redo()
}

func (ed *Editor) Right(sel bool) bool {
	return ed.cursor.Right(sel)
}

func (ed *Editor) SelectAll() bool {
	if !ed.enabled {
		return false
	}

	ed.cursor.Set(0, 0, false)
	ed.cursor.Set(math.MaxInt, math.MaxInt, true)

	return true
}

func (ed *Editor) Top(sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Top(sel)
}

func (ed *Editor) Undo() bool {
	if !ed.enabled {
		return false
	}

	return ed.history.Undo()
}

func (ed *Editor) Up(sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Up(1, sel)
}

func (ed *Editor) deleteChar() {
	cur := ed.cursor

	ed.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
	ed.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)

	ed.history.Push()
}

func (ed *Editor) deletePrevChar() {
	cur := ed.cursor

	if cur.Ln > 0 && cur.Col == 0 {
		l := 0
		for range ed.buffer.IterLine(cur.Ln, false) {
			l += 1
			if l == 2 {
				break
			}
		}

		if l == 1 {
			ed.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			ed.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			cur.Left(false)
		} else {
			cur.Left(false)
			ed.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			ed.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
		}
	} else {
		ed.buffer.Delete2(cur.Ln, cur.Col-1, cur.Ln, cur.Col)
		ed.syntax.Delete(cur.Ln, cur.Col-1, cur.Ln, cur.Col)
		cur.Left(false)
	}

	ed.history.Push()
}

func (ed *Editor) deleteSelection() {
	cur := ed.cursor

	ed.buffer.Delete2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	ed.syntax.Delete(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	ed.cursor.Set(cur.StartLn, cur.StartCol, false)

	ed.history.Push()
}

func (ed *Editor) insertText(text string) {
	cur := ed.cursor

	if cur.Selecting {
		ed.buffer.Delete2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		ed.syntax.Delete(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		cur.Set(cur.StartLn, cur.StartCol, false)
	}

	ed.buffer.Insert2(cur.Ln, cur.Col, text)

	startLn := cur.Ln
	startCol := cur.Col

	dLn, dCol := grapheme.Graphemes.MeasureText(text)
	cur.Forward(dLn, dCol)

	endLn := cur.Ln
	endCol := cur.Col

	ed.syntax.Insert(startLn, startCol, endLn, endCol)

	ed.history.Push()
}
