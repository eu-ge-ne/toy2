package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Delete struct {
	editor *Editor
}

func (h *Delete) Match(k key.Key) bool {
	return k.Name == "DELETE"
}

func (h *Delete) Run(key.Key) bool {
	if h.editor.cursor.Selecting {
		h.editor.deleteSelection()
	} else {
		h.editor.deleteChar()
	}

	return true
}
