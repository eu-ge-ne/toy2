package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type ThemeNeutral struct {
	app    *App
	option palette.Option
}

func NewThemeNeutral(app *App) *ThemeNeutral {
	return &ThemeNeutral{
		app:    app,
		option: palette.NewOption("Theme Neutral", "Theme: Neutral", []key.Key{}),
	}
}

func (t *ThemeNeutral) Option() *palette.Option {
	return &t.option
}

func (t *ThemeNeutral) Match(key.Key) bool {
	return false
}

func (t *ThemeNeutral) Run() {
	t.app.setColors(theme.Neutral{})
}
