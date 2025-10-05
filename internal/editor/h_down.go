package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type DownHandler struct {
	editor *Editor
}

func (h *DownHandler) Match(k key.Key) bool {
	return k.Name == "DOWN"
}

func (h *DownHandler) Handle(k key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.cursor.Down(1, k.Mods&key.Shift != 0)
}
