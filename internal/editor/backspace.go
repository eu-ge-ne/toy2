package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Backspace struct {
	editor *Editor
}

func (h *Backspace) Match(k key.Key) bool {
	return k.Name == "BACKSPACE"
}

func (h *Backspace) Run(key.Key) bool {
	if h.editor.cursor.Selecting {
		h.editor.deleteSelection()
	} else {
		h.editor.deletePrevChar()
	}

	return true
}
