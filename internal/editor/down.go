package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Down struct {
	editor *Editor
}

func (h *Down) Match(k key.Key) bool {
	return k.Name == "DOWN"
}

func (h *Down) Run(k key.Key) bool {
	return h.editor.Down(1, k.Mods&key.Shift != 0)
}
