package palette

import (
	"strings"

	"github.com/eu-ge-ne/toy2/internal/key"
)

type Option struct {
	Id          string
	description string
	shortcuts   string
}

func NewOption(id string, description string, keys []key.Key) *Option {
	shortcuts := make([]string, len(keys))

	for i, key := range keys {
		s := ""
		if key.Shift {
			s += "⇧"
		}
		if key.Ctrl {
			s += "⌃"
		}
		if key.Alt {
			s += "⌥"
		}
		if key.Super {
			s += "⌘"
		}
		s += strings.ToUpper(key.Name)
		shortcuts[i] = s
	}

	return &Option{
		Id:          id,
		description: description,
		shortcuts:   strings.Join(shortcuts, " "),
	}
}
