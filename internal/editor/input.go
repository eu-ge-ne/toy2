package editor

import (
	"slices"
	"time"

	"github.com/rivo/uniseg"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/key"
)

func (ed *Editor) HandleKey(key *key.Key) bool {
	t0 := time.Now()

	i := slices.IndexFunc(ed.handlers, func(h Handler) bool {
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

func (ed *Editor) insert(text string) {
	if ed.Cursor.Selecting {
		ed.Buffer.Delete(ed.Cursor.FromLn, ed.Cursor.FromCol, ed.Cursor.ToLn, ed.Cursor.ToCol+1)
		ed.Cursor.Set(ed.Cursor.FromLn, ed.Cursor.FromCol, false)
	}

	ed.Buffer.Insert(ed.Cursor.Ln, ed.Cursor.Col, text)

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
		ed.Cursor.Forward(uniseg.GraphemeClusterCount(text))
	} else {
		col := uniseg.GraphemeClusterCount(text) - lastEolIndex - 1

		ed.Cursor.Set(ed.Cursor.Ln+eolCount, col, false)
	}

	ed.History.Push()
}

func (ed *Editor) backspace() {
	if ed.Cursor.Ln > 0 && ed.Cursor.Col == 0 {
		l := 0
		for range ed.Buffer.Line(ed.Cursor.Ln, false) {
			l += 1
			if l == 2 {
				break
			}
		}

		if l == 1 {
			ed.Buffer.Delete(ed.Cursor.Ln, ed.Cursor.Col, ed.Cursor.Ln, ed.Cursor.Col+1)
			ed.Cursor.Left(false)
		} else {
			ed.Cursor.Left(false)
			ed.Buffer.Delete(ed.Cursor.Ln, ed.Cursor.Col, ed.Cursor.Ln, ed.Cursor.Col+1)
		}
	} else {
		ed.Buffer.Delete(ed.Cursor.Ln, ed.Cursor.Col-1, ed.Cursor.Ln, ed.Cursor.Col)
		ed.Cursor.Left(false)
	}

	ed.History.Push()
}

func (ed *Editor) deleteChar() {
	ed.Buffer.Delete(ed.Cursor.Ln, ed.Cursor.Col, ed.Cursor.Ln, ed.Cursor.Col+1)

	ed.History.Push()
}

func (ed *Editor) deleteSelection() {
	ed.Buffer.Delete(ed.Cursor.FromLn, ed.Cursor.FromCol, ed.Cursor.Ln, ed.Cursor.Col+1)

	ed.Cursor.Set(ed.Cursor.FromLn, ed.Cursor.FromCol, false)

	ed.History.Push()
}
