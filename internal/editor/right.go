package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Right struct {
	editor *Editor
}

func (h *Right) Match(k key.Key) bool {
	return k.Name == "RIGHT"
}

func (h *Right) Run(k key.Key) bool {
	return h.editor.cursor.Right(k.Mods&key.Shift != 0)
}
