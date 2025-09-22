package header

import (
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
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

func (h *Header) Layout(a ui.Area) {
	h.area = ui.Area{
		Y: a.Y,
		X: a.X,
		W: a.W,
		H: std.Clamp(1, 0, a.H),
	}
}

func (h *Header) Render() {
	if !h.Enabled {
		return
	}

	vt.Sync.Bsu()

	span := h.area.W

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(vt.SaveCursor)
	vt.Buf.Write(h.colorBackground)
	vt.ClearArea(vt.Buf, h.area)
	vt.SetCursor(vt.Buf, h.area.Y, h.area.X)
	vt.Buf.Write(h.colorFilePath)
	vt.WriteTextCenter(vt.Buf, &span, h.filePath)
	vt.Buf.Write(vt.RestoreCursor)
	vt.Buf.Write(vt.ShowCursor)

	vt.Buf.Flush()

	vt.Sync.Esu()
}

func (h *Header) SetFilePath(filePath string) {
	h.filePath = filePath

	h.Render()
}
