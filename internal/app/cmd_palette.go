package app

import (
	"slices"

	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type PaletteCommand struct {
	app *App
}

func NewPaletteCommand(app *App) *PaletteCommand {
	return &PaletteCommand{app}
}

func (c *PaletteCommand) Option() *palette.Option {
	return nil
}

func (c *PaletteCommand) Match(key *key.Key) bool {
	return key.Name == "F1"
}

func (c *PaletteCommand) Run() {
	c.app.editor.Enabled = false

	done := make(chan *palette.Option)

	go c.app.palette.Open(done)

	option := <-done

	c.app.editor.Enabled = true

	c.app.editor.Render()

	if option != nil {
		i := slices.IndexFunc(c.app.commands, func(c Command) bool {
			o := c.Option()
			return o != nil && o.Id == option.Id
		})

		if i >= 0 {
			c.app.commands[i].Run()
		}
	}
}
