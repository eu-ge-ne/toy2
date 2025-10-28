package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type PageUp struct {
	editor *Editor
}

func (h *PageUp) Match(k key.Key) bool {
	return k.Name == "PAGE_UP"
}

func (h *PageUp) Run(k key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.cursor.Up(h.editor.area.H, k.Mods&key.Shift != 0)
}
