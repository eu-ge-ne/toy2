package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type CopyHandler struct {
	editor *Editor
}

func (h *CopyHandler) Match(k key.Key) bool {
	return k.Name == "c" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *CopyHandler) Handle(key.Key) bool {
	return h.editor.Copy()
}
