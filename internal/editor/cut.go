package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Cut struct {
	editor *Editor
}

func (h *Cut) Match(k key.Key) bool {
	return k.Name == "x" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Cut) Run(key.Key) bool {
	ed := h.editor

	if !ed.enabled {
		return false
	}

	cur := ed.cursor

	if cur.Selecting {
		ed.clipboard = std.IterToStr(ed.buffer.Read(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol))
		ed.deleteSelection()
	} else {
		ed.clipboard = std.IterToStr(ed.buffer.Read(cur.Ln, cur.Col, cur.Ln, cur.Col+1))
		ed.deleteChar()
	}

	vt.CopyToClipboard(vt.Sync, ed.clipboard)

	return true
}
