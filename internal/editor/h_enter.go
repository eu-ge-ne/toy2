package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type EnterHandler struct {
	editor *Editor
}

func (h *EnterHandler) Match(k key.Key) bool {
	return k.Name == "ENTER"
}

func (h *EnterHandler) Handle(key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	h.editor.insert("\n")

	return true
}
