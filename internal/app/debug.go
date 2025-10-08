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

func (c *Debug) Option() *palette.Option {
	return &c.option
}

func (c *Debug) Match(key.Key) bool {
	return false
}

func (c *Debug) Run() bool {
	c.app.debug.ToggleEnabled()

	return true
}
