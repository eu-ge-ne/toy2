package command

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Palette struct {
	app App
}

func NewPalette(app App) *Palette {
	return &Palette{app}
}

func (c *Palette) Option() *palette.Option {
	return nil
}

func (c *Palette) Match(k key.Key) bool {
	return k.Name == "F1"
}

func (c *Palette) Run() {
	c.app.Palette()
}
