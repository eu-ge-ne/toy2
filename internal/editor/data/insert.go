package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Insert struct {
	Editor Editor
}

func (h *Insert) Match(k key.Key) bool {
	return len(k.Text) != 0
}

func (h *Insert) Handle(k key.Key) bool {
	return h.Editor.Insert(k.Text)
}
