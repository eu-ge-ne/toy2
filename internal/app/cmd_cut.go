package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type CutCommand struct {
	app    *App
	option palette.Option
}

func NewCutCommand(app *App) *CutCommand {
	return &CutCommand{
		app: app,
		option: palette.NewOption(
			"Cut",
			"Edit: Cut",
			[]key.Key{{Name: "x", Mods: key.Ctrl}, {Name: "x", Mods: key.Super}},
		),
	}
}

func (c *CutCommand) Option() *palette.Option {
	return &c.option
}

func (c *CutCommand) Match(key.Key) bool {
	return false
}

func (c *CutCommand) Run() {
	if c.app.editor.Enabled {
		if c.app.editor.Cut() {
			c.app.editor.Render()
		}
	}
}
