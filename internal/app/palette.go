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

func (p *Palette) Option() *palette.Option {
	return nil
}

func (p *Palette) Match(k key.Key) bool {
	return k.Name == "F1"
}

func (p *Palette) Run() bool {
	p.app.editor.SetEnabled(false)

	defer func() {
		p.app.editor.SetEnabled(true)
	}()

	option := <-p.app.palette.Open()

	if option == nil {
		return false
	}

	for _, c := range p.app.commands {
		o := c.Option()
		if o != nil && o.Id == option.Id {
			c.Run()
			return true
		}
	}

	return true
}
