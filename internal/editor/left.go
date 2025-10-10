package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Left struct {
	editor *Editor
}

func (h *Left) Match(k key.Key) bool {
	return k.Name == "LEFT"
}

func (h *Left) Run(k key.Key) bool {
	return h.editor.cursor.Left(k.Mods&key.Shift != 0)
}
