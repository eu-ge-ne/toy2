package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Undo struct {
	app    *App
	option palette.Option
}

func NewUndo(app *App) *Undo {
	return &Undo{
		app: app,
		option: palette.NewOption(
			"Undo",
			"Edit: Undo",
			[]key.Key{{Name: "z", Mods: key.Ctrl}, {Name: "z", Mods: key.Super}},
		),
	}
}

func (c *Undo) Option() *palette.Option {
	return &c.option
}

func (c *Undo) Match(key.Key) bool {
	return false
}

func (c *Undo) Run() bool {
	return c.app.editor.Handlers["UNDO"].Run(key.Key{})
}
