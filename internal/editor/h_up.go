package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type UpHandler struct {
	editor *Editor
}

func (h *UpHandler) Match(key key.Key) bool {
	return key.Name == "UP"
}

func (h *UpHandler) Handle(key key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.Cursor.Up(1, key.Shift)
}
