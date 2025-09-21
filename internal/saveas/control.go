package saveas

import (
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (sv *SaveAs) Layout(a ui.Area) {
	w := std.Clamp(60, 0, a.W)
	h := std.Clamp(10, 0, a.H)

	sv.area = ui.Area{
		Y: a.Y + (a.H-h)/2,
		X: a.X + (a.W-w)/2,
		W: w,
		H: h,
	}

	sv.editor.Layout(ui.Area{
		Y: sv.area.Y + 4,
		X: sv.area.X + 2,
		W: sv.area.W - 4,
		H: 1,
	})
}

func (sv *SaveAs) Render() {
	if !sv.enabled {
		return
	}

	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(sv.colorBackground)
	vt.ClearArea(vt.Buf, sv.area)
	vt.SetCursor(vt.Buf, sv.area.Y+1, sv.area.X)
	vt.Buf.Write(sv.colorText)
	span := sv.area.W
	vt.WriteTextCenter(vt.Buf, &span, "Save As")
	vt.SetCursor(vt.Buf, sv.area.Y+sv.area.H-2, sv.area.X)
	span = sv.area.W
	vt.WriteTextCenter(vt.Buf, &span, "ESC‧cancel    ENTER‧ok")

	sv.editor.Render()

	vt.Buf.Flush()

	vt.Sync.Esu()
}
