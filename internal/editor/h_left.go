package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type LeftHandler struct {
	editor *Editor
}

func (h *LeftHandler) Match(key *key.Key) bool {
	return key.Name == "LEFT"
}

func (h *LeftHandler) Handle(key *key.Key) bool {
	return h.editor.Cursor.Left(key.Shift)
}
