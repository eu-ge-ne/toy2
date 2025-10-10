package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type PageDown struct {
	editor *Editor
}

func (h *PageDown) Match(k key.Key) bool {
	return k.Name == "PAGE_DOWN"
}

func (h *PageDown) Run(k key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.cursor.Down(h.editor.pageSize, k.Mods&key.Shift != 0)
}
