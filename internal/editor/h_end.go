package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type EndHandler struct {
	editor *Editor
}

func (h *EndHandler) Match(k key.Key) bool {
	switch {
	case k.Name == "END":
		return true
	case k.Name == "RIGHT" && k.Mods&key.Super != 0:
		return true
	}
	return false
}

func (h *EndHandler) Handle(k key.Key) bool {
	return h.editor.Cursor.End(k.Mods&key.Shift != 0)
}
