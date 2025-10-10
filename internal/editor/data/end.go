package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type End struct {
	data *Data
}

func (h *End) Match(k key.Key) bool {
	switch {
	case k.Name == "END":
		return true
	case k.Name == "RIGHT" && k.Mods&key.Super != 0:
		return true
	}
	return false
}

func (h *End) Handle(k key.Key) bool {
	return h.data.GoEnd(k.Mods&key.Shift != 0)
}
