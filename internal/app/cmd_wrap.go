package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type WrapCommand struct {
	app    *App
	option *palette.Option
}

func NewWrapCommand(app *App) *WrapCommand {
	return &WrapCommand{
		app:    app,
		option: palette.NewOption("Wrap", "View: Toggle Line Wrap", []key.Key{{Name: "F6"}}),
	}
}

func (c *WrapCommand) Option() *palette.Option {
	return c.option
}

func (c *WrapCommand) Match(key key.Key) bool {
	return key.Name == "F6"
}

func (c *WrapCommand) Run() {
	c.app.editor.WrapEnabled = !c.app.editor.WrapEnabled
	c.app.editor.Cursor.Home(false)

	c.app.Render()
}
