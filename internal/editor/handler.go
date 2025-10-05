package editor

import (
	"math"
	"slices"
	"time"

	"github.com/rivo/uniseg"

	"github.com/eu-ge-ne/toy2/internal/editor/handler"
	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (ed *Editor) HandleKey(key key.Key) bool {
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
	cur := ed.cursor

	if cur.Selecting {
		ed.clipboard = ed.buffer.ReadSegPosRange(cur.FromLn, cur.FromCol, cur.ToLn, cur.ToCol+1)
		cur.Set(cur.Ln, cur.Col, false)
	} else {
		ed.clipboard = ed.buffer.ReadSegPosRange(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
	}

	vt.CopyToClipboard(vt.Sync, ed.clipboard)

	return false
}

func (ed *Editor) Cut() bool {
	cur := ed.cursor

	if cur.Selecting {
		ed.clipboard = ed.buffer.ReadSegPosRange(cur.FromLn, cur.FromCol, cur.ToLn, cur.ToCol+1)
		ed.deleteSelection()
	} else {
		ed.clipboard = ed.buffer.ReadSegPosRange(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
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

	ed.Insert("\n")

	return true
}

func (ed *Editor) Home(sel bool) bool {
	return ed.cursor.Home(sel)
}

func (ed *Editor) Insert(text string) bool {
	if ed.cursor.Selecting {
		ed.buffer.DeleteSegPosRange(ed.cursor.FromLn, ed.cursor.FromCol, ed.cursor.ToLn, ed.cursor.ToCol+1)
		ed.cursor.Set(ed.cursor.FromLn, ed.cursor.FromCol, false)
	}

	ed.buffer.InsertSegPos(ed.cursor.Ln, ed.cursor.Col, text)

	eolCount := 0
	lastEolIndex := 0

	gs := uniseg.NewGraphemes(text)
	i := 0
	for gs.Next() {
		g := grapheme.Graphemes.Get(gs.Str())
		if g.IsEol {
			eolCount += 1
			lastEolIndex = i
		}
		i += 1
	}

	if eolCount == 0 {
		ed.cursor.Forward(uniseg.GraphemeClusterCount(text))
	} else {
		col := uniseg.GraphemeClusterCount(text) - lastEolIndex - 1

		ed.cursor.Set(ed.cursor.Ln+eolCount, col, false)
	}

	ed.history.Push()

	return true
}

func (ed *Editor) Left(sel bool) bool {
	return ed.cursor.Left(sel)
}

func (ed *Editor) PageDown(sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Down(ed.area.H, sel)
}

func (ed *Editor) PageUp(sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Up(ed.area.H, sel)
}

func (ed *Editor) Paste() bool {
	if len(ed.clipboard) == 0 {
		return false
	}

	ed.Insert(ed.clipboard)

	return true
}

func (ed *Editor) Redo() bool {
	return ed.history.Redo()
}

func (ed *Editor) Right(sel bool) bool {
	return ed.cursor.Right(sel)
}

func (ed *Editor) SelectAll() bool {
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
	return ed.history.Undo()
}

func (ed *Editor) Up(sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Up(1, sel)
}
