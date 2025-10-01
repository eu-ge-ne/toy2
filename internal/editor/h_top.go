package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type TopHandler struct {
	editor *Editor
}

func (h *TopHandler) Match(key *key.Key) bool {
	return key.Name == "UP" && key.Super
}

func (h *TopHandler) Handle(key *key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.Cursor.Top(key.Shift)
}
