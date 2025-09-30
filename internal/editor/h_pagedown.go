package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type PageDownHandler struct {
	editor *Editor
}

func (h *PageDownHandler) Match(key key.Key) bool {
	return key.Name == "PAGE_DOWN"
}

func (h *PageDownHandler) Handle(key key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.Cursor.Down(h.editor.area.H, key.Shift)
}
