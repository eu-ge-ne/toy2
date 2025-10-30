package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Up struct {
	editor *Editor
}

func (h *Up) Match(k key.Key) bool {
	return k.Name == "UP"
}

func (h *Up) Run(k key.Key) bool {
	return h.editor.Up(1, k.Mods&key.Shift != 0)
}
