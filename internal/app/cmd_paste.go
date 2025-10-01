package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type PasteCommand struct {
	app    *App
	option *palette.Option
}

func NewPasteCommand(app *App) *PasteCommand {
	return &PasteCommand{
		app:    app,
		option: palette.NewOption("Paste", "Edit: Paste", []key.Key{{Name: "v", Ctrl: true}, {Name: "v", Super: true}}),
	}
}

func (c *PasteCommand) Option() *palette.Option {
	return c.option
}

func (c *PasteCommand) Match(key *key.Key) bool {
	return false
}

func (c *PasteCommand) Run() {
	if c.app.editor.Enabled {
		if c.app.editor.Paste() {
			c.app.editor.Render()
		}
	}
}
