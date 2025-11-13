package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Debug struct {
	app    *App
	option palette.Option
}

func NewDebug(app *App) *Debug {
	return &Debug{
		app:    app,
		option: palette.NewOption("Debug", "Global: Toggle Debug Panel", []key.Key{}),
	}
}

func (d *Debug) Option() *palette.Option {
	return &d.option
}

func (d *Debug) Match(key.Key) bool {
	return false
}

func (d *Debug) Run() {
	d.app.debug.ToggleEnabled()
}
