package handler

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Copy struct {
	Editor Editor
}

func (h *Copy) Match(k key.Key) bool {
	return k.Name == "c" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Copy) Handle(key.Key) bool {
	return h.Editor.Copy()
}
