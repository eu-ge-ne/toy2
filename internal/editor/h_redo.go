package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type RedoHandler struct {
	editor *Editor
}

func (h *RedoHandler) Match(k key.Key) bool {
	return k.Name == "y" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *RedoHandler) Handle(key.Key) bool {
	return h.editor.History.Redo()
}
