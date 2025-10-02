package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type UndoHandler struct {
	editor *Editor
}

func (h *UndoHandler) Match(k key.Key) bool {
	return k.Name == "z" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *UndoHandler) Handle(key.Key) bool {
	return h.editor.History.Undo()
}
