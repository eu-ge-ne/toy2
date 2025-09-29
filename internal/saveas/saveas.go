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
	return &SaveAs{
		editor: editor.New(false),
	}
}

func (sv *SaveAs) SetColors(t theme.Tokens) {
	sv.colorBackground = t.Light1Bg()
	sv.colorText = append(t.Light1Bg(), t.Light1Fg()...)

	sv.editor.SetColors(t)
}

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

func (sv *SaveAs) Open(filePath string, done chan<- string) {
	sv.enabled = true
	sv.editor.Enabled = true

	sv.editor.Buffer.Reset(filePath)
	sv.editor.Reset(false)

	sv.Render()

	result := sv.processInput()

	sv.enabled = false
	sv.editor.Enabled = false

	done <- result
}

func (sv *SaveAs) processInput() string {
	for {
		for key := range vt.Read() {
			switch key.Name {
			case "ESC":
				return ""
			case "ENTER":
				return sv.editor.Buffer.Text()
			}

			if sv.editor.HandleKey(key) {
				sv.editor.Render()
			}
		}
	}
}
