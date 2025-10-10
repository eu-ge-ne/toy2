package command

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type ThemeStone struct {
	app    App
	option palette.Option
}

func NewThemeStone(app App) *ThemeStone {
	return &ThemeStone{
		app:    app,
		option: palette.NewOption("Theme Stone", "Theme: Stone", []key.Key{}),
	}
}

func (c *ThemeStone) Option() *palette.Option {
	return &c.option
}

func (c *ThemeStone) Match(key.Key) bool {
	return false
}

func (c *ThemeStone) Run() {
	c.app.ThemeStone()
}
