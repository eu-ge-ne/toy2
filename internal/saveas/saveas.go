package saveas

import (
	"github.com/eu-ge-ne/toy2/internal/editor"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type SaveAs struct {
	area    ui.Area
	enabled bool
	editor  *editor.Editor

	colorBackground []byte
	colorText       []byte
}

func New() *SaveAs {
	return &SaveAs{editor: editor.New(false)}
}

func (sv *SaveAs) SetColors(t theme.Theme) {
	sv.colorBackground = t.Light1Bg()
	sv.colorText = append(t.Light1Bg(), t.Light1Fg()...)

	sv.editor.SetColors(t)
}

func (sv *SaveAs) SetArea(a ui.Area) {
	w := std.Clamp(60, 0, a.W)
	h := std.Clamp(10, 0, a.H)

	sv.area = ui.Area{
		Y: a.Y + (a.H-h)/2,
		X: a.X + (a.W-w)/2,
		W: w,
		H: h,
	}

	sv.editor.SetArea(ui.Area{
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

func (sv *SaveAs) Open(filePath string) <-chan string {
	done := make(chan string)

	go func() {
		sv.enabled = true
		sv.editor.SetEnabled(true)

		sv.editor.SetText(filePath)
		sv.editor.End(false)

		sv.Render()

		result := sv.processInput()

		sv.enabled = false
		sv.editor.SetEnabled(false)

		done <- result
	}()

	return done
}

func (sv *SaveAs) processInput() string {
	for {
		key := vt.ReadKey()

		switch key.Name {
		case "ESC":
			return ""
		case "ENTER":
			fp := sv.editor.GetText()
			if len(fp) > 0 {
				return fp
			}
		}

		if sv.editor.HandleKey(key) {
			sv.editor.Render()
		}
	}
}
