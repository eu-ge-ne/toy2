package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type PageUpHandler struct {
	editor *Editor
}

func (h *PageUpHandler) Match(key *key.Key) bool {
	return key.Name == "PAGE_UP"
}

func (h *PageUpHandler) Handle(key *key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.Cursor.Up(h.editor.area.H, key.Shift)
}
