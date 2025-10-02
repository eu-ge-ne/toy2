package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type CutHandler struct {
	editor *Editor
}

func (h *CutHandler) Match(k key.Key) bool {
	return k.Name == "x" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *CutHandler) Handle(key.Key) bool {
	return h.editor.Cut()
}
