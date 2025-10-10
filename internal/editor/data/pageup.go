package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type PageUp struct {
	data *Data
}

func (h *PageUp) Match(k key.Key) bool {
	return k.Name == "PAGE_UP"
}

func (h *PageUp) Handle(k key.Key) bool {
	return h.data.PageUp(k.Mods&key.Shift != 0)
}
