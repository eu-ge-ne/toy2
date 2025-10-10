package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Redo struct {
	data *Data
}

func (h *Redo) Match(k key.Key) bool {
	return k.Name == "y" && (k.Mods&key.Ctrl != 0 || k.Mods&key.Super != 0)
}

func (h *Redo) Handle(key.Key) bool {
	return h.data.Redo()
}
