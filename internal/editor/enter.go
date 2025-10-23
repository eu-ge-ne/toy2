package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Enter struct {
	editor *Editor
}

func (h *Enter) Match(k key.Key) bool {
	return k.Name == "ENTER"
}

func (h *Enter) Run(key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	if h.editor.cursor.Selecting {
		h.editor.deleteSelection()
	}

	h.editor.insertText("\n")

	return true
}
