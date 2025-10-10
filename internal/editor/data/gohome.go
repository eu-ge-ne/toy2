package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type GoHome struct {
	data *Data
}

func (h *GoHome) Match(k key.Key) bool {
	switch {
	case k.Name == "HOME":
		return true
	case k.Name == "LEFT" && k.Mods&key.Super != 0:
		return true
	}
	return false
}

func (h *GoHome) Handle(k key.Key) bool {
	return h.data.GoHome(k.Mods&key.Shift != 0)
}

func (d *Data) GoHome(sel bool) bool {
	return d.cursor.Home(sel)
}
