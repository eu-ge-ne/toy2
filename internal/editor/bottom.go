package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Bottom struct {
	editor *Editor
}

func (h *Bottom) Match(k key.Key) bool {
	return k.Name == "DOWN" && k.Mods&key.Super != 0
}

func (h *Bottom) Run(k key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.cursor.Bottom(k.Mods&key.Shift != 0)
}
