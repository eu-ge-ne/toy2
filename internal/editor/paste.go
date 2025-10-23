package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Paste struct {
	editor *Editor
}

func (h *Paste) Match(k key.Key) bool {
	return k.Name == "v" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Paste) Run(key.Key) bool {
	if !h.editor.enabled {
		return false
	}

	if len(h.editor.clipboard) == 0 {
		return false
	}

	if h.editor.cursor.Selecting {
		h.editor.deleteSelection()
	}

	h.editor.insertText(h.editor.clipboard)

	return true
}
