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
	flag     bool
	enabled  bool

	colorBackground []byte
	colorFilePath   []byte
	colorFlag       []byte
}

func New() *Header {
	return &Header{}
}

func (h *Header) Enable(enable bool) {
	h.enabled = enable
}

func (h *Header) SetColors(t theme.Theme) {
	h.colorBackground = t.Dark0Bg()
	h.colorFilePath = append(t.Dark0Bg(), t.Dark0Fg()...)
	h.colorFlag = append(t.Dark0Bg(), t.Light2Fg()...)
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
	if !h.enabled {
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

	if h.flag {
		vt.Buf.Write(h.colorFlag)
		vt.WriteText(vt.Buf, &span, " +")
	}

	vt.Buf.Write(vt.RestoreCursor)
	vt.Buf.Write(vt.ShowCursor)

	vt.Buf.Flush()

	vt.Sync.Esu()
}

func (h *Header) SetFilePath(filePath string) {
	h.filePath = filePath

	h.Render()
}

func (h *Header) SetFlag(flag bool) {
	h.flag = flag

	h.Render()
}
