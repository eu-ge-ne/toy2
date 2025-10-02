package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type TextHandler struct {
	editor *Editor
}

func (h *TextHandler) Match(k key.Key) bool {
	return len(k.Text) != 0
}

func (h *TextHandler) Handle(k key.Key) bool {
	h.editor.insert(k.Text)

	return true
}
