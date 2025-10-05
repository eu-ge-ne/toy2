package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type SelectAllCommand struct {
	app    *App
	option palette.Option
}

func NewSelectAllCommand(app *App) *SelectAllCommand {
	return &SelectAllCommand{
		app: app,
		option: palette.NewOption(
			"Select All",
			"Edit: Select All",
			[]key.Key{{Name: "a", Mods: key.Ctrl}, {Name: "a", Mods: key.Super}},
		),
	}
}

func (c *SelectAllCommand) Option() *palette.Option {
	return &c.option
}

func (c *SelectAllCommand) Match(key.Key) bool {
	return false
}

func (c *SelectAllCommand) Run() {
	if c.app.editor.SelectAll() {
		c.app.editor.Render()
	}
}
