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
	return h.data.GoDown(k.Mods&key.Shift != 0)
}

func (d *Data) GoDown(sel bool) bool {
	if !d.multiLine {
		return false
	}

	return d.cursor.Down(1, sel)
}
