package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type CopyHandler struct {
	editor *Editor
}

func (h *CopyHandler) Match(key *key.Key) bool {
	return key.Name == "c" && (key.Ctrl || key.Super)
}

func (h *CopyHandler) Handle(key *key.Key) bool {
	return h.editor.Copy()
}
