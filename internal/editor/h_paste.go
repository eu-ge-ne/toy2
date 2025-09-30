package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type PasteHandler struct {
	editor *Editor
}

func (h *PasteHandler) Match(key key.Key) bool {
	return key.Name == "v" && (key.Ctrl || key.Super)
}

func (h *PasteHandler) Handle(key key.Key) bool {
	return h.editor.Paste()
}
