package editor

import (
	"math"
	"slices"
	"time"

	"github.com/eu-ge-ne/toy2/internal/editor/handler"
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (ed *Editor) HandleKey(key key.Key) bool {
	if !ed.enabled {
		return false
	}

	t0 := time.Now()

	i := slices.IndexFunc(ed.handlers, func(h handler.Handler) bool {
		return h.Match(key)
	})

	if i < 0 {
		return false
	}

	r := ed.handlers[i].Handle(key)

	if ed.OnKeyHandled != nil {
		ed.OnKeyHandled(time.Since(t0))
	}

	return r
}

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
