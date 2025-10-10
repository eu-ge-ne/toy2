package editor

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/key"
)

type SelectAll struct {
	editor *Editor
}

func (h *SelectAll) Match(k key.Key) bool {
	return k.Name == "a" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *SelectAll) Run(key.Key) bool {
	if !h.editor.enabled {
		return false
	}

	h.editor.cursor.Set(0, 0, false)
	h.editor.cursor.Set(math.MaxInt, math.MaxInt, true)

	return true
}
