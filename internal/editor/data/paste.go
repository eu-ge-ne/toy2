package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Paste struct {
	Editor Editor
}

func (h *Paste) Match(k key.Key) bool {
	return k.Name == "v" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Paste) Handle(key.Key) bool {
	return h.Editor.Paste()
}
