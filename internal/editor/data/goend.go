package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type GoEnd struct {
	data *Data
}

func (h *GoEnd) Match(k key.Key) bool {
	switch {
	case k.Name == "END":
		return true
	case k.Name == "RIGHT" && k.Mods&key.Super != 0:
		return true
	}
	return false
}

func (h *GoEnd) Handle(k key.Key) bool {
	return h.data.GoEnd(k.Mods&key.Shift != 0)
}

func (d *Data) GoEnd(sel bool) bool {
	return d.cursor.End(sel)
}
