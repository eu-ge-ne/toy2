package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type BottomHandler struct {
	editor *Editor
}

func (h *BottomHandler) Match(k key.Key) bool {
	return k.Name == "DOWN" && k.Mods&key.Super != 0
}

func (h *BottomHandler) Handle(k key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.Cursor.Bottom(k.Mods&key.Shift != 0)
}
