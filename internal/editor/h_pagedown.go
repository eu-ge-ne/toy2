package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type PageDownHandler struct {
	editor *Editor
}

func (h *PageDownHandler) Match(k key.Key) bool {
	return k.Name == "PAGE_DOWN"
}

func (h *PageDownHandler) Handle(k key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.Cursor.Down(h.editor.area.H, k.Mods&key.Shift != 0)
}
