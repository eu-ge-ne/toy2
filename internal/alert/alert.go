package alert

import (
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Alert struct {
	area    ui.Area
	enabled bool
	text    string

	colorBackground []byte
	colorText       []byte
}

func New() *Alert {
	return &Alert{}
}

func (al *Alert) SetColors(t theme.Tokens) {
	al.colorBackground = t.DangerBg()
	al.colorText = append(t.DangerBg(), t.Light1Fg()...)
}

func (al *Alert) Open(text string, done chan<- struct{}) {
	al.enabled = true

	al.text = text
	al.Render()

	al.processInput()

	al.enabled = false

	done <- struct{}{}
}

func (al *Alert) processInput() {
	for {
		for key := range vt.Read() {
			switch key.Name {
			case "ESC":
				return
			case "ENTER":
				return
			}
		}
	}
}
