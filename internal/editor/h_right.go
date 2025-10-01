package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type RightHandler struct {
	editor *Editor
}

func (h *RightHandler) Match(key *key.Key) bool {
	return key.Name == "RIGHT"
}

func (h *RightHandler) Handle(key *key.Key) bool {
	return h.editor.Cursor.Right(key.Shift)
}
