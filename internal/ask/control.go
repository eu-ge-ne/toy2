package ask

import (
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (ask *Ask) Layout(a ui.Area) {
	w := std.Clamp(60, 0, a.W)
	h := std.Clamp(7, 0, a.H)

	ask.area = ui.Area{
		Y: a.Y + (a.H-h)/2,
		X: a.X + (a.W-w)/2,
		W: w,
		H: h,
	}
}

func (ask *Ask) Render() {
	if !ask.enabled {
		return
	}

	vt.Sync.Bsu()

	vt.Buf.Write(vt.HideCursor)
	vt.Buf.Write(ask.colorBackground)
	vt.ClearArea(vt.Buf, ask.area)

	runes := []rune(ask.text)

	for y := ask.area.Y + 1; y < ask.area.Y+ask.area.H-3; y += 1 {
		if len(runes) == 0 {
			break
		}

		span := ask.area.W - 2
		i := min(span, len(runes))
		line := runes[:i]
		runes = runes[i:]

		vt.SetCursor(vt.Buf, y, ask.area.X+1)
		vt.Buf.Write(ask.colorText)
		vt.WriteTextCenter(vt.Buf, &span, string(line))
	}

	vt.SetCursor(vt.Buf, ask.area.Y+ask.area.H-2, ask.area.X)
	span := ask.area.W
	vt.WriteTextCenter(vt.Buf, &span, "ESC‧no    ENTER‧yes")

	vt.Buf.Flush()

	vt.Sync.Esu()
}
