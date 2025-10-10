package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type ThemeZinc struct {
	app    *App
	option palette.Option
}

func NewThemeZinc(app *App) *ThemeZinc {
	return &ThemeZinc{
		app:    app,
		option: palette.NewOption("Theme Zinc", "Theme: Zinc", []key.Key{}),
	}
}

func (c *ThemeZinc) Option() *palette.Option {
	return &c.option
}

func (c *ThemeZinc) Match(key.Key) bool {
	return false
}

func (c *ThemeZinc) Run() bool {
	c.app.setColors(theme.Zinc{})

	return true
}
