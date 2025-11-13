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

func (z *Zen) Option() *palette.Option {
	return &z.option
}

func (z *Zen) Match(k key.Key) bool {
	return k.Name == "F11"
}

func (z *Zen) Run() {
	z.app.setZenEnabled(!z.app.zenEnabled)
}
