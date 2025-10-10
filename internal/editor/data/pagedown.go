package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type PageDown struct {
	Editor Editor
}

func (h *PageDown) Match(k key.Key) bool {
	return k.Name == "PAGE_DOWN"
}

func (h *PageDown) Handle(k key.Key) bool {
	return h.Editor.PageDown(k.Mods&key.Shift != 0)
}
