package alert

import (
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (al *Alert) Layout(a ui.Area) {
	w := std.Clamp(60, 0, a.W)
	h := std.Clamp(10, 0, a.H)

	al.area = ui.Area{
		Y: a.Y + (a.H-h)/2,
		X: a.X + (a.W-w)/2,
		W: w,
		H: h,
	}
}

func (al *Alert) Render() {
	if !al.enabled {
		return
	}

	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(al.colorBackground)
	vt.ClearArea(vt.Buf, al.area)

	runes := []rune(al.text)

	for y := al.area.Y + 1; y < al.area.Y+al.area.H-3; y += 1 {
		if len(runes) == 0 {
			break
		}

		span := al.area.W - 2
		i := min(span, len(runes))
		line := runes[:i]
		runes = runes[i:]

		vt.SetCursor(vt.Buf, y, al.area.X+1)
		vt.Buf.Write(al.colorText)
		vt.WriteTextCenter(vt.Buf, &span, string(line))
	}

	vt.SetCursor(vt.Buf, al.area.Y+al.area.H-2, al.area.X)
	span := al.area.W
	vt.WriteTextCenter(vt.Buf, &span, "ENTER‧ok")

	vt.Buf.Flush()

	vt.Sync.Esu()
}
