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
		shortcuts[i] = k.Shortcut()
	}

	return &Option{
		Id:          id,
		Description: description,
		shortcuts:   strings.Join(shortcuts, " "),
	}
}
