package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Paste struct {
	editor *Editor
}

func (h *Paste) Match(k key.Key) bool {
	return k.Name == "v" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Paste) Run(key.Key) bool {
	return h.editor.Paste()
}
