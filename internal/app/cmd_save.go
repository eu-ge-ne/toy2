package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type SaveCommand struct {
	app    *App
	option palette.Option
}

func NewSaveCommand(app *App) *SaveCommand {
	return &SaveCommand{
		app:    app,
		option: palette.NewOption("Save", "Global: Save", []key.Key{{Name: "F2"}}),
	}
}

func (c *SaveCommand) Option() *palette.Option {
	return &c.option
}

func (c *SaveCommand) Match(k key.Key) bool {
	return k.Name == "F2"
}

func (c *SaveCommand) Run() {
	c.app.editor.Enabled = false

	if c.app.save() {
		c.app.editor.ResetCursor()
	}

	c.app.editor.Enabled = true

	c.app.editor.Render()
}
