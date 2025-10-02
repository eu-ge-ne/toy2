package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type UndoCommand struct {
	app    *App
	option *palette.Option
}

func NewUndoCommand(app *App) *UndoCommand {
	return &UndoCommand{
		app: app,
		option: palette.NewOption(
			"Undo",
			"Edit: Undo",
			[]key.Key{{Name: "z", Mods: key.Ctrl}, {Name: "z", Mods: key.Super}},
		),
	}
}

func (c *UndoCommand) Option() *palette.Option {
	return c.option
}

func (c *UndoCommand) Match(key.Key) bool {
	return false
}

func (c *UndoCommand) Run() {
	if !c.app.editor.Enabled {
		return
	}

	if c.app.editor.History.Undo() {
		c.app.editor.Render()
	}
}
