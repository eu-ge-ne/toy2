package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Cut struct {
	editor *Editor
}

func (h *Cut) Match(k key.Key) bool {
	return k.Name == "x" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Cut) Run(key.Key) bool {
	return h.editor.Cut()
}
