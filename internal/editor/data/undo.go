package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Undo struct {
	Editor Editor
}

func (h *Undo) Match(k key.Key) bool {
	return k.Name == "z" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Undo) Handle(key.Key) bool {
	return h.Editor.Undo()
}
