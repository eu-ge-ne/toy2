package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Home struct {
	editor *Editor
}

func (h *Home) Match(k key.Key) bool {
	switch {
	case k.Name == "HOME":
		return true
	case k.Name == "LEFT" && k.Mods&key.Super != 0:
		return true
	}
	return false
}

func (h *Home) Run(k key.Key) bool {
	return h.editor.Home(k.Mods&key.Shift != 0)
}
