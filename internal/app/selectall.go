package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type SelectAll struct {
	app    *App
	option palette.Option
}

func NewSelectAll(app *App) *SelectAll {
	return &SelectAll{
		app: app,
		option: palette.NewOption(
			"Select All",
			"Edit: Select All",
			[]key.Key{{Name: "a", Mods: key.Ctrl}, {Name: "a", Mods: key.Super}},
		),
	}
}

func (c *SelectAll) Option() *palette.Option {
	return &c.option
}

func (c *SelectAll) Match(key.Key) bool {
	return false
}

func (c *SelectAll) Run() bool {
	return c.app.editor.SelectAll()
}
