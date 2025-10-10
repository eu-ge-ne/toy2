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

func (c *Palette) Run() bool {
	c.app.editor.SetEnabled(false)

	done := make(chan *palette.Option)

	go c.app.palette.Open(done)

	option := <-done

	c.app.editor.SetEnabled(true)

	for _, c := range c.app.commands {
		o := c.Option()
		if o != nil && o.Id == option.Id {
			c.Run()
			return true
		}
	}

	return true
}
