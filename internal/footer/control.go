package footer

import (
	"fmt"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (f *Footer) Layout(a ui.Area) {
	f.area = ui.Area{Y: a.Y + a.H - 1, X: a.X, W: a.W, H: std.Clamp(1, 0, a.H)}
}

func (f *Footer) Render() {
	vt.Bsu()

	vt.WriteBuf(
		vt.HideCursor,
		vt.SaveCursor,
		f.colorBackground,
		vt.ClearArea(f.area.Y, f.area.X, f.area.W, f.area.H),
		vt.SetCursor(f.area.Y, f.area.X),
		f.colorText,
		fmt.Appendf(nil, "%*s", f.area.W, f.cursorStatus),
	)

	vt.WriteBuf(
		vt.RestoreCursor,
		vt.ShowCursor,
	)

	vt.FlushBuf()

	vt.Esu()
}
