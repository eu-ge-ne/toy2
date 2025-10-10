package command

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Wrap struct {
	app    App
	option palette.Option
}

func NewWrap(app App) *Wrap {
	return &Wrap{
		app:    app,
		option: palette.NewOption("Wrap", "View: Toggle Line Wrap", []key.Key{{Name: "F6"}}),
	}
}

func (c *Wrap) Option() *palette.Option {
	return &c.option
}

func (c *Wrap) Match(k key.Key) bool {
	return k.Name == "F6"
}

func (c *Wrap) Run() {
	c.app.Wrap()
}
