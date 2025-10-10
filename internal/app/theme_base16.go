package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type ThemeBase16 struct {
	app    *App
	option palette.Option
}

func NewThemeBase16(app *App) *ThemeBase16 {
	return &ThemeBase16{
		app:    app,
		option: palette.NewOption("Theme Base16", "Theme: Base16", []key.Key{}),
	}
}

func (c *ThemeBase16) Option() *palette.Option {
	return &c.option
}

func (c *ThemeBase16) Match(key.Key) bool {
	return false
}

func (c *ThemeBase16) Run() {
	c.app.setColors(theme.Base16{})

	c.app.Render()
}
