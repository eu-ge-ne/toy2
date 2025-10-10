package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Down struct {
	data *Data
}

func (h *Down) Match(k key.Key) bool {
	return k.Name == "DOWN"
}

func (h *Down) Handle(k key.Key) bool {
	return h.data.Down(k.Mods&key.Shift != 0)
}
