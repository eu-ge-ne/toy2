package editor

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/key"
)

type LeftHandler struct {
	editor *Editor
}

func (h *LeftHandler) Match(key key.Key) bool {
	return key.Name == "LEFT"
}

func (h *LeftHandler) Handle(key key.Key) bool {
	if h.editor.Cursor.Move(0, -1, key.Shift) {
		return true
	}

	if h.editor.Cursor.Ln > 0 {
		return h.editor.Cursor.Move(-1, math.MaxInt, key.Shift)
	}

	return false
}
