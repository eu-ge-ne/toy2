package editor

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/key"
)

type SelectAllHandler struct {
	editor *Editor
}

func (h *SelectAllHandler) Match(k key.Key) bool {
	return k.Name == "a" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *SelectAllHandler) Handle(key.Key) bool {
	h.editor.Cursor.Set(0, 0, false)
	h.editor.Cursor.Set(math.MaxInt, math.MaxInt, true)

	return true
}
