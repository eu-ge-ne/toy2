package editor

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/key"
)

type SelectAllHandler struct {
	editor *Editor
}

func (h *SelectAllHandler) Match(key key.Key) bool {
	return key.Name == "a" && (key.Ctrl || key.Super)
}

func (h *SelectAllHandler) Handle(key key.Key) bool {
	h.editor.Cursor.Set(0, 0, false)
	h.editor.Cursor.Set(math.MaxInt, math.MaxInt, true)

	return true
}
