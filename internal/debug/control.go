package debug

import (
	"fmt"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (d *Debug) ToggleEnabled() {
	d.enabled = !d.enabled
}

func (d *Debug) Layout(a ui.Area) {
	w := std.Clamp(30, 0, a.W)
	h := std.Clamp(7, 0, a.H)

	d.area = ui.Area{
		Y: a.Y + a.H - h,
		X: a.X + a.W - w,
		W: w,
		H: h,
	}
}

func (d *Debug) Render() {
	if !d.enabled {
		return
	}

	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(vt.SaveCursor)
	vt.Buf.Write(d.colorBackground)
	vt.ClearArea(vt.Buf, d.area.Y, d.area.X, d.area.W, d.area.H)
	vt.Buf.Write(d.colorText)
	vt.SetCursor(vt.Buf, d.area.Y+1, d.area.X+1)
	fmt.Fprintf(vt.Buf, "Input    : %v", d.inputTime)
	vt.SetCursor(vt.Buf, d.area.Y+2, d.area.X+1)
	fmt.Fprintf(vt.Buf, "Render   : %v", d.renderTime)
	vt.Buf.Write(vt.RestoreCursor)
	vt.Buf.Write(vt.ShowCursor)

	vt.Buf.Flush()

	vt.Sync.Esu()
}
