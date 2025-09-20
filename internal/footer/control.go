package footer

import (
	"fmt"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (f *Footer) Layout(a ui.Area) {
	f.area = ui.Area{
		Y: a.Y + a.H - 1,
		X: a.X,
		W: a.W,
		H: std.Clamp(1, 0, a.H),
	}
}

func (f *Footer) Render() {
	if !f.Enabled {
		return
	}

	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(vt.SaveCursor)
	vt.Buf.Write(f.colorBackground)
	vt.ClearArea(vt.Buf, f.area)
	vt.SetCursor(vt.Buf, f.area.Y, f.area.X)
	vt.Buf.Write(f.colorText)
	fmt.Fprintf(vt.Buf, "%*s", f.area.W, f.cursorStatus)
	vt.Buf.Write(vt.RestoreCursor)
	vt.Buf.Write(vt.ShowCursor)

	vt.Buf.Flush()

	vt.Sync.Esu()
}
