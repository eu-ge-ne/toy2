package header

import (
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
)

type Header struct {
	area     ui.Area
	filePath string
	Enabled  bool

	colorBackground  []byte
	colorFilePath    []byte
	colorUnsavedFlag []byte
}

func New() *Header {
	return &Header{}
}

func (h *Header) SetColors(t theme.Tokens) {
	h.colorBackground = t.Dark0Bg()
	h.colorFilePath = append(t.Dark0Bg(), t.Dark0Fg()...)
	h.colorUnsavedFlag = append(t.Dark0Bg(), t.Light2Fg()...)
}

func (h *Header) SetFilePath(filePath string) {
	h.filePath = filePath

	h.Render()
}
