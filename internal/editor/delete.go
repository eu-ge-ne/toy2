package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Delete struct {
	editor *Editor
}

func (h *Delete) Match(k key.Key) bool {
	return k.Name == "DELETE"
}

func (h *Delete) Run(key.Key) bool {
	return h.editor.Delete()
}
