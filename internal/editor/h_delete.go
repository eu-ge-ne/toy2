package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type DeleteHandler struct {
	editor *Editor
}

func (h *DeleteHandler) Match(k key.Key) bool {
	return k.Name == "DELETE"
}

func (h *DeleteHandler) Handle(key.Key) bool {
	if h.editor.cursor.Selecting {
		h.editor.deleteSelection()
	} else {
		h.editor.deleteChar()
	}

	return true
}
