package command

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type ThemeSlate struct {
	app    App
	option palette.Option
}

func NewThemeSlate(app App) *ThemeSlate {
	return &ThemeSlate{
		app:    app,
		option: palette.NewOption("Theme Slate", "Theme: Slate", []key.Key{}),
	}
}

func (c *ThemeSlate) Option() *palette.Option {
	return &c.option
}

func (c *ThemeSlate) Match(key.Key) bool {
	return false
}

func (c *ThemeSlate) Run() {
	c.app.ThemeSlate()
}
