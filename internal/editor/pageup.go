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
	return h.editor.Up(h.editor.area.H, k.Mods&key.Shift != 0)
}
