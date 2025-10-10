package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Insert struct {
	editor *Editor
}

func (h *Insert) Match(k key.Key) bool {
	return len(k.Text) != 0
}

func (h *Insert) Run(k key.Key) bool {
	h.editor.insertText(k.Text)

	return true
}
