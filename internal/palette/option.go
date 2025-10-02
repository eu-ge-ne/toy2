package palette

import (
	"strings"

	"github.com/eu-ge-ne/toy2/internal/key"
)

type Option struct {
	Id          string
	Description string
	shortcuts   string
}

func NewOption(id string, description string, keys []key.Key) *Option {
	shortcuts := make([]string, len(keys))

	for i, k := range keys {
		s := ""
		if k.Mods&key.Shift != 0 {
			s += "⇧"
		}
		if k.Mods&key.Ctrl != 0 {
			s += "⌃"
		}
		if k.Mods&key.Alt != 0 {
			s += "⌥"
		}
		if k.Mods&key.Super != 0 {
			s += "⌘"
		}
		s += strings.ToUpper(k.Name)
		shortcuts[i] = s
	}

	return &Option{
		Id:          id,
		Description: description,
		shortcuts:   strings.Join(shortcuts, " "),
	}
}
