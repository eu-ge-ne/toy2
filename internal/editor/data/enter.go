package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Enter struct {
	Editor Editor
}

func (h *Enter) Match(k key.Key) bool {
	return k.Name == "ENTER"
}

func (h *Enter) Handle(key.Key) bool {
	return h.Editor.Enter()
}
