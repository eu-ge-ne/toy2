package handler

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Top struct {
	Editor Editor
}

func (h *Top) Match(k key.Key) bool {
	return k.Name == "UP" && k.Mods&key.Super != 0
}

func (h *Top) Handle(k key.Key) bool {
	return h.Editor.Top(k.Mods&key.Shift != 0)
}
