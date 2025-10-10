package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Exit struct {
	app    *App
	option palette.Option
}

func NewExit(app *App) *Exit {
	return &Exit{
		app:    *app,
		option: palette.NewOption("Exit", "Global: Exit", []key.Key{{Name: "F10"}}),
	}
}

func (c *Exit) Option() *palette.Option {
	return &c.option
}

func (c *Exit) Match(k key.Key) bool {
	return k.Name == "F10"
}

func (c *Exit) Run() {
	c.app.editor.SetEnabled(false)

	if c.app.editor.HasChanges() {
		save := make(chan bool)
		go c.app.ask.Open("Save changes?", save)

		if <-save {
			c.app.save()
		}
	}

	c.app.exit()
}
