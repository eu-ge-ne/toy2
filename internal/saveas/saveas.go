package saveas

import (
	"github.com/eu-ge-ne/toy2/internal/editor"
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
				return "todo"
			}

			if sv.editor.HandleKey(key) {
				sv.editor.Render()
			}
		}
	}
}
