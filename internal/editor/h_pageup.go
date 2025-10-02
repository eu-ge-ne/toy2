package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type PageUpHandler struct {
	editor *Editor
}

func (h *PageUpHandler) Match(k key.Key) bool {
	return k.Name == "PAGE_UP"
}

func (h *PageUpHandler) Handle(k key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.Cursor.Up(h.editor.area.H, k.Mods&key.Shift != 0)
}
