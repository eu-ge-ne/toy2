package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Up struct {
	data *Data
}

func (h *Up) Match(k key.Key) bool {
	return k.Name == "UP"
}

func (h *Up) Handle(k key.Key) bool {
	return h.data.Up(k.Mods&key.Shift != 0)
}
