package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type HomeHandler struct {
	editor *Editor
}

func (h *HomeHandler) Match(key key.Key) bool {
	switch {
	case key.Name == "HOME":
		return true
	case key.Name == "LEFT" && key.Super:
		return true
	}
	return false
}

func (h *HomeHandler) Handle(key key.Key) bool {
	return h.editor.Cursor.Home(key.Shift)
}
