package editor

import (
	"slices"
	"time"

	"github.com/rivo/uniseg"

	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/key"
)

func (ed *Editor) HandleKey(key key.Key) bool {
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
	if ed.cursor.Selecting {
		ed.Buffer.SegDelete2(ed.cursor.FromLn, ed.cursor.FromCol, ed.cursor.ToLn, ed.cursor.ToCol)
		ed.cursor.Set(ed.cursor.FromLn, ed.cursor.FromCol, false)
	}

	ed.Buffer.SegInsert2(ed.cursor.Ln, ed.cursor.Col, text)

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
}

func (ed *Editor) backspace() {
	if ed.cursor.Ln > 0 && ed.cursor.Col == 0 {
		l := 0
		for range ed.Buffer.SegLine(ed.cursor.Ln) {
			l += 1
			if l == 2 {
				break
			}
		}

		if l == 1 {
			ed.Buffer.SegDelete2(ed.cursor.Ln, ed.cursor.Col, ed.cursor.Ln, ed.cursor.Col+1)
			ed.cursor.Left(false)
		} else {
			ed.cursor.Left(false)
			ed.Buffer.SegDelete2(ed.cursor.Ln, ed.cursor.Col, ed.cursor.Ln, ed.cursor.Col+1)
		}
	} else {
		ed.Buffer.SegDelete2(ed.cursor.Ln, ed.cursor.Col-1, ed.cursor.Ln, ed.cursor.Col)
		ed.cursor.Left(false)
	}
}

func (ed *Editor) deleteChar() {
	ed.Buffer.SegDelete2(ed.cursor.Ln, ed.cursor.Col, ed.cursor.Ln, ed.cursor.Col+1)
}

func (ed *Editor) deleteSelection() {
	ed.Buffer.SegDelete2(ed.cursor.FromLn, ed.cursor.FromCol, ed.cursor.Ln, ed.cursor.Col+1)

	ed.cursor.Set(ed.cursor.FromLn, ed.cursor.FromCol, false)
}
