package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Delete struct {
	Editor Editor
}

func (h *Delete) Match(k key.Key) bool {
	return k.Name == "DELETE"
}

func (h *Delete) Handle(key.Key) bool {
	return h.Editor.Delete()
}
