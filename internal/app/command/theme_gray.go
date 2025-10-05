package command

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type ThemeGray struct {
	app    App
	option palette.Option
}

func NewThemeGray(app App) *ThemeGray {
	return &ThemeGray{
		app:    app,
		option: palette.NewOption("Theme Gray", "Theme: Gray", []key.Key{}),
	}
}

func (c *ThemeGray) Option() *palette.Option {
	return &c.option
}

func (c *ThemeGray) Match(key.Key) bool {
	return false
}

func (c *ThemeGray) Run() {
	c.app.ThemeGray()
}
