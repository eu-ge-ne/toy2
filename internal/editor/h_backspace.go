package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type BackspaceHandler struct {
	editor *Editor
}

func (h *BackspaceHandler) Match(k key.Key) bool {
	return k.Name == "BACKSPACE"
}

func (h *BackspaceHandler) Handle(key.Key) bool {
	if h.editor.cursor.Selecting {
		h.editor.deleteSelection()
	} else {
		h.editor.backspace()
	}

	return true
}
