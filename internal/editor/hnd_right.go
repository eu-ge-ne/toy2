package editor

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/key"
)

type RightHandler struct {
	editor *Editor
}

func (h *RightHandler) Match(key key.Key) bool {
	return key.Name == "RIGHT"
}

func (h *RightHandler) Handle(key key.Key) bool {
	if h.editor.Cursor.Move(0, 1, key.Shift) {
		return true
	}

	if h.editor.Cursor.Ln < h.editor.Buffer.LineCount()-1 {
		return h.editor.Cursor.Move(1, math.MinInt, key.Shift)
	}

	return false
}
