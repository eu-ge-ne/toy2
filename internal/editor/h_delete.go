package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type DeleteHandler struct {
	editor *Editor
}

func (h *DeleteHandler) Match(key *key.Key) bool {
	return key.Name == "DELETE"
}

func (h *DeleteHandler) Handle(key *key.Key) bool {
	if h.editor.Cursor.Selecting {
		h.editor.deleteSelection()
	} else {
		h.editor.deleteChar()
	}

	return true
}
