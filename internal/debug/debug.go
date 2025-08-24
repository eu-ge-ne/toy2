package debug

import (
	"time"

	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
)

type Debug struct {
	area    ui.Area
	enabled bool

	inputTime  time.Duration
	renderTime time.Duration

	colorBackground []byte
	colorText       []byte
}

func New() *Debug {
	return &Debug{enabled: true}
}

func (d *Debug) SetColors(t theme.Tokens) {
	d.colorBackground = t.Light0Bg()
	d.colorText = append(t.Light0Bg(), t.Dark0Fg()...)
}

func (d *Debug) SetInputTime(elapsed time.Duration) {
	d.inputTime = elapsed

	if d.enabled {
		d.Render()
	}
}

func (d *Debug) SetRenderTime(elapsed time.Duration) {
	d.renderTime = elapsed

	if d.enabled {
		d.Render()
	}
}
