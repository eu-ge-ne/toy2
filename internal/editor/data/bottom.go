package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Bottom struct {
	data *Data
}

func (h *Bottom) Match(k key.Key) bool {
	return k.Name == "DOWN" && k.Mods&key.Super != 0
}

func (h *Bottom) Handle(k key.Key) bool {
	return h.data.Bottom(k.Mods&key.Shift != 0)
}
