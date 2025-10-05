package command

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Whitespace struct {
	app    App
	option palette.Option
}

func NewWhitespace(app App) *Whitespace {
	return &Whitespace{
		app:    app,
		option: palette.NewOption("Whitespace", "View: Toggle Render Whitespace", []key.Key{{Name: "F5"}}),
	}
}

func (c *Whitespace) Option() *palette.Option {
	return &c.option
}

func (c *Whitespace) Match(k key.Key) bool {
	return k.Name == "F5"
}

func (c *Whitespace) Run() {
	c.app.Whitespace()
}
