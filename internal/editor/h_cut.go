package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type CutHandler struct {
	editor *Editor
}

func (h *CutHandler) Match(key *key.Key) bool {
	return key.Name == "x" && (key.Ctrl || key.Super)
}

func (h *CutHandler) Handle(key *key.Key) bool {
	return h.editor.Cut()
}
