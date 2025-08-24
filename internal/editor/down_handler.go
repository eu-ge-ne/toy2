package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type DownHandler struct {
	editor *Editor
}

func (h *DownHandler) Match(key key.Key) bool {
	return key.Name == "DOWN"
}

func (h *DownHandler) Handle(key key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.cursor.Down(1, key.Shift)
}
