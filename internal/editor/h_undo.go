package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type UndoHandler struct {
	editor *Editor
}

func (h *UndoHandler) Match(key *key.Key) bool {
	return key.Name == "z" && (key.Ctrl || key.Super)
}

func (h *UndoHandler) Handle(key *key.Key) bool {
	return h.editor.History.Undo()
}
