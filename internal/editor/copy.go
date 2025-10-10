package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Copy struct {
	editor *Editor
}

func (h *Copy) Match(k key.Key) bool {
	return k.Name == "c" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Copy) Run(key.Key) bool {
	ed := h.editor

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
