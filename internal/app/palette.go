package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Palette struct {
	app *App
}

func NewPalette(app *App) *Palette {
	return &Palette{app}
}

func (c *Palette) Option() *palette.Option {
	return nil
}

func (c *Palette) Match(k key.Key) bool {
	return k.Name == "F1"
}

func (c *Palette) Run() {
	c.app.editor.SetEnabled(false)

	done := make(chan *palette.Option)

	go c.app.palette.Open(done)

	option := <-done

	c.app.editor.SetEnabled(true)

	c.app.editor.Render()

	if option != nil {
		i := slices.IndexFunc(c.app.commands, func(c command.Command) bool {
			o := c.Option()
			return o != nil && o.Id == option.Id
		})

		if i >= 0 {
			app.commands[i].Run()
		}
	}
}
