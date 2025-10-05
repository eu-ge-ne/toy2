package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type TopHandler struct {
	editor *Editor
}

func (h *TopHandler) Match(k key.Key) bool {
	return k.Name == "UP" && k.Mods&key.Super != 0
}

func (h *TopHandler) Handle(k key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.cursor.Top(k.Mods&key.Shift != 0)
}
