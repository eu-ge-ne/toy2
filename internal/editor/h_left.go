package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type LeftHandler struct {
	editor *Editor
}

func (h *LeftHandler) Match(k key.Key) bool {
	return k.Name == "LEFT"
}

func (h *LeftHandler) Handle(k key.Key) bool {
	return h.editor.cursor.Left(k.Mods&key.Shift != 0)
}
