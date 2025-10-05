package command

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Copy struct {
	app    App
	option palette.Option
}

func NewCopy(app App) *Copy {
	return &Copy{
		app: app,
		option: palette.NewOption(
			"Copy",
			"Edit: Copy",
			[]key.Key{{Name: "c", Mods: key.Ctrl}, {Name: "c", Mods: key.Super}},
		),
	}
}

func (c *Copy) Option() *palette.Option {
	return &c.option
}

func (c *Copy) Match(key.Key) bool {
	return false
}

func (c *Copy) Run() {
	c.app.Copy()
}
