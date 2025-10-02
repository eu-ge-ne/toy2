package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type WhitespaceCommand struct {
	app    *App
	option palette.Option
}

func NewWhitespaceCommand(app *App) *WhitespaceCommand {
	return &WhitespaceCommand{
		app:    app,
		option: palette.NewOption("Whitespace", "View: Toggle Render Whitespace", []key.Key{{Name: "F5"}}),
	}
}

func (c *WhitespaceCommand) Option() *palette.Option {
	return &c.option
}

func (c *WhitespaceCommand) Match(k key.Key) bool {
	return k.Name == "F5"
}

func (c *WhitespaceCommand) Run() {
	c.app.editor.WhitespaceEnabled = !c.app.editor.WhitespaceEnabled

	c.app.Render()
}
