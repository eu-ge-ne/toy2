package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type RightHandler struct {
	editor *Editor
}

func (h *RightHandler) Match(k key.Key) bool {
	return k.Name == "RIGHT"
}

func (h *RightHandler) Handle(k key.Key) bool {
	return h.editor.cursor.Right(k.Mods&key.Shift != 0)
}
