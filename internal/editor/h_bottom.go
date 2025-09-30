package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type BottomHandler struct {
	editor *Editor
}

func (h *BottomHandler) Match(key key.Key) bool {
	return key.Name == "DOWN" && key.Super
}

func (h *BottomHandler) Handle(key key.Key) bool {
	if !h.editor.multiLine {
		return false
	}

	return h.editor.Cursor.Bottom(key.Shift)
}
