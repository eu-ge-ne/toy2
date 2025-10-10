package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Right struct {
	Editor Editor
}

func (h *Right) Match(k key.Key) bool {
	return k.Name == "RIGHT"
}

func (h *Right) Handle(k key.Key) bool {
	return h.Editor.Right(k.Mods&key.Shift != 0)
}
