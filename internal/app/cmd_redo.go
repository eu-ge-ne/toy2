package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type RedoCommand struct {
	app    *App
	option *palette.Option
}

func NewRedoCommand(app *App) *RedoCommand {
	return &RedoCommand{
		app: app,
		option: palette.NewOption(
			"Redo",
			"Edit: Redo",
			[]key.Key{{Name: "y", Mods: key.Ctrl}, {Name: "y", Mods: key.Super}},
		),
	}
}

func (c *RedoCommand) Option() *palette.Option {
	return c.option
}

func (c *RedoCommand) Match(key.Key) bool {
	return false
}

func (c *RedoCommand) Run() {
	if !c.app.editor.Enabled {
		return
	}

	if c.app.editor.History.Redo() {
		c.app.editor.Render()
	}
}
