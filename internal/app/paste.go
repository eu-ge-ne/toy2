package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Paste struct {
	app    *App
	option palette.Option
}

func NewPaste(app *App) *Paste {
	return &Paste{
		app: app,
		option: palette.NewOption(
			"Paste",
			"Edit: Paste",
			[]key.Key{{Name: "v", Mods: key.Ctrl}, {Name: "v", Mods: key.Super}},
		),
	}
}

func (c *Paste) Option() *palette.Option {
	return &c.option
}

func (c *Paste) Match(key.Key) bool {
	return false
}

func (c *Paste) Run() {
	if c.app.editor.Handlers["PASTE"].Run(key.Key{}) {
		c.app.editor.Render()
	}
}
