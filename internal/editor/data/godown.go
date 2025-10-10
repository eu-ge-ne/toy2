package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type GoDown struct {
	data *Data
}

func (h *GoDown) Match(k key.Key) bool {
	return k.Name == "DOWN"
}

func (h *GoDown) Handle(k key.Key) bool {
	return h.data.Down(k.Mods&key.Shift != 0)
}
