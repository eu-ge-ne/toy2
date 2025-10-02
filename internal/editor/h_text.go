package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type TextHandler struct {
	editor *Editor
}

func (h *TextHandler) Match(key key.Key) bool {
	return len(key.Text) != 0
}

func (h *TextHandler) Handle(key key.Key) bool {
	h.editor.insert(key.Text)

	return true
}
