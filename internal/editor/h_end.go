package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type EndHandler struct {
	editor *Editor
}

func (h *EndHandler) Match(key *key.Key) bool {
	switch {
	case key.Name == "END":
		return true
	case key.Name == "RIGHT" && key.Super:
		return true
	}
	return false
}

func (h *EndHandler) Handle(key *key.Key) bool {
	return h.editor.Cursor.End(key.Shift)
}
