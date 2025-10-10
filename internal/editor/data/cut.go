package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Cut struct {
	Editor Editor
}

func (h *Cut) Match(k key.Key) bool {
	return k.Name == "x" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Cut) Handle(key.Key) bool {
	return h.Editor.Cut()
}
