package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type RedoHandler struct {
	editor *Editor
}

func (h *RedoHandler) Match(key key.Key) bool {
	return key.Name == "y" && (key.Ctrl || key.Super)
}

func (h *RedoHandler) Handle(key key.Key) bool {
	return h.editor.History.Redo()
}
