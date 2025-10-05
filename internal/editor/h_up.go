package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type UpHandler struct {
	editor *Editor
}

func (h *UpHandler) Match(k key.Key) bool {
	return k.Name == "UP"
}

func (h *UpHandler) Handle(k key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.cursor.Up(1, k.Mods&key.Shift != 0)
}
