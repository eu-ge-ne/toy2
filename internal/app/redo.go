package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Redo struct {
	app    *App
	option palette.Option
}

func NewRedo(app *App) *Redo {
	return &Redo{
		app: app,
		option: palette.NewOption(
			"Redo",
			"Edit: Redo",
			[]key.Key{{Name: "y", Mods: key.Ctrl}, {Name: "y", Mods: key.Super}},
		),
	}
}

func (c *Redo) Option() *palette.Option {
	return &c.option
}

func (c *Redo) Match(key.Key) bool {
	return false
}

func (c *Redo) Run() bool {
	return c.app.editor.Redo()
}
