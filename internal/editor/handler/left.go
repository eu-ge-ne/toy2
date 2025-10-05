package handler

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Left struct {
	Editor Editor
}

func (h *Left) Match(k key.Key) bool {
	return k.Name == "LEFT"
}

func (h *Left) Handle(k key.Key) bool {
	return h.Editor.Left(k.Mods&key.Shift != 0)
}
