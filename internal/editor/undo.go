package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Undo struct {
	editor *Editor
}

func (h *Undo) Match(k key.Key) bool {
	return k.Name == "z" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Undo) Run(key.Key) bool {
	return h.editor.Undo()
}
