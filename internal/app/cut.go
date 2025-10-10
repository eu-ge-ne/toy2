package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Cut struct {
	app    *App
	option palette.Option
}

func NewCut(app *App) *Cut {
	return &Cut{
		app: app,
		option: palette.NewOption(
			"Cut",
			"Edit: Cut",
			[]key.Key{{Name: "x", Mods: key.Ctrl}, {Name: "x", Mods: key.Super}},
		),
	}
}

func (c *Cut) Option() *palette.Option {
	return &c.option
}

func (c *Cut) Match(key.Key) bool {
	return false
}

func (c *Cut) Run() {
	if c.app.editor.Handlers["CUT"].Run(key.Key{}) {
		c.app.editor.Render()
	}
}
