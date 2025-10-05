package handler

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Backspace struct {
	Editor Editor
}

func (h *Backspace) Match(k key.Key) bool {
	return k.Name == "BACKSPACE"
}

func (h *Backspace) Handle(key.Key) bool {
	return h.Editor.Backspace()
}
