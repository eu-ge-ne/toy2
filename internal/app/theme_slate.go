package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type ThemeSlate struct {
	app    *App
	option palette.Option
}

func NewThemeSlate(app *App) *ThemeSlate {
	return &ThemeSlate{
		app:    app,
		option: palette.NewOption("Theme Slate", "Theme: Slate", []key.Key{}),
	}
}

func (t *ThemeSlate) Option() *palette.Option {
	return &t.option
}

func (t *ThemeSlate) Match(key.Key) bool {
	return false
}

func (t *ThemeSlate) Run() {
	t.app.setColors(theme.Slate{})
}
