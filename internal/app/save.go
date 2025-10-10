package command

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Save struct {
	app    App
	option palette.Option
}

func NewSave(app App) *Save {
	return &Save{
		app:    app,
		option: palette.NewOption("Save", "Global: Save", []key.Key{{Name: "F2"}}),
	}
}

func (c *Save) Option() *palette.Option {
	return &c.option
}

func (c *Save) Match(k key.Key) bool {
	return k.Name == "F2"
}

func (c *Save) Run() {
	c.app.Save()
}
