package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Zen struct {
	app    *App
	option palette.Option
}

func NewZen(app *App) *Zen {
	return &Zen{
		app:    app,
		option: palette.NewOption("Zen", "Global: Toggle Zen Mode", []key.Key{{Name: "F11"}}),
	}
}

func (c *Zen) Option() *palette.Option {
	return &c.option
}

func (c *Zen) Match(k key.Key) bool {
	return k.Name == "F11"
}

func (c *Zen) Run() bool {
	c.app.setZenEnabled(!c.app.zenEnabled)

	return true
}
