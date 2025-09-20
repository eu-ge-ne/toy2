package ask

import (
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Ask struct {
	area    ui.Area
	enabled bool
	text    string

	colorBackground []byte
	colorText       []byte
}

func New() *Ask {
	return &Ask{}
}

func (ask *Ask) SetColors(t theme.Tokens) {
	ask.colorBackground = t.Light1Bg()
	ask.colorText = append(t.Light1Bg(), t.Light1Fg()...)
}

func (ask *Ask) Open(text string, done chan<- bool) {
	ask.enabled = true

	ask.text = text
	ask.Render()

	result := ask.processInput()

	ask.enabled = false

	done <- result
}

func (ask *Ask) processInput() bool {
	for {
		for key := range vt.Read() {
			switch key.Name {
			case "ESC":
				return false
			case "ENTER":
				return true
			}
		}
	}
}
