package command

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type ThemeNeutral struct {
	app    App
	option palette.Option
}

func NewThemeNeutral(app App) *ThemeNeutral {
	return &ThemeNeutral{
		app:    app,
		option: palette.NewOption("Theme Neutral", "Theme: Neutral", []key.Key{}),
	}
}

func (c *ThemeNeutral) Option() *palette.Option {
	return &c.option
}

func (c *ThemeNeutral) Match(key.Key) bool {
	return false
}

func (c *ThemeNeutral) Run() {
	c.app.ThemeNeutral()
}
