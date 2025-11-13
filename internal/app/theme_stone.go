package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type ThemeStone struct {
	app    *App
	option palette.Option
}

func NewThemeStone(app *App) *ThemeStone {
	return &ThemeStone{
		app:    app,
		option: palette.NewOption("Theme Stone", "Theme: Stone", []key.Key{}),
	}
}

func (t *ThemeStone) Option() *palette.Option {
	return &t.option
}

func (t *ThemeStone) Match(key.Key) bool {
	return false
}

func (t *ThemeStone) Run() {
	t.app.setColors(theme.Stone{})
}
