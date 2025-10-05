package handler

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Up struct {
	Editor Editor
}

func (h *Up) Match(k key.Key) bool {
	return k.Name == "UP"
}

func (h *Up) Handle(k key.Key) bool {
	return h.Editor.Up(k.Mods&key.Shift != 0)
}
