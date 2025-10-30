package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Redo struct {
	editor *Editor
}

func (h *Redo) Match(k key.Key) bool {
	return k.Name == "y" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Redo) Run(key.Key) bool {
	return h.editor.Redo()
}
