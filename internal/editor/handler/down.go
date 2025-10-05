package handler

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Down struct {
	Editor Editor
}

func (h *Down) Match(k key.Key) bool {
	return k.Name == "DOWN"
}

func (h *Down) Handle(k key.Key) bool {
	return h.Editor.Down(k.Mods&key.Shift != 0)
}
