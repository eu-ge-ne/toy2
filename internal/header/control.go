package header

import (
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (h *Header) Layout(a ui.Area) {
	h.area = ui.Area{
		Y: a.Y,
		X: a.X,
		W: a.W,
		H: std.Clamp(1, 0, a.H),
	}
}

func (h *Header) Render() {
	vt.Sync.Bsu()

	span := h.area.W

	vt.Buf.Write(
		vt.HideCursor,
		vt.SaveCursor,
		h.colorBackground,
	)
	vt.ClearArea(vt.Buf, h.area.Y, h.area.X, h.area.W, h.area.H)
	vt.SetCursor(vt.Buf, h.area.Y, h.area.X)
	vt.Buf.Write(
		h.colorFilePath,
		vt.WriteTextCenter(&span, h.filePath),
		vt.RestoreCursor,
		vt.ShowCursor,
	)

	vt.Buf.Flush()

	vt.Sync.Esu()
}
