package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type CopyCommand struct {
	app    *App
	option *palette.Option
}

func NewCopyCommand(app *App) *CopyCommand {
	return &CopyCommand{
		app: app,
		option: palette.NewOption(
			"Copy",
			"Edit: Copy",
			[]key.Key{{Name: "c", Mods: key.Ctrl}, {Name: "c", Mods: key.Super}},
		),
	}
}

func (c *CopyCommand) Option() *palette.Option {
	return c.option
}

func (c *CopyCommand) Match(key.Key) bool {
	return false
}

func (c *CopyCommand) Run() {
	if c.app.editor.Enabled {
		if c.app.editor.Copy() {
			c.app.editor.Render()
		}
	}
}
