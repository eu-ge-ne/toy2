package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type SlateThemeCommand struct {
	app    *App
	option *palette.Option
}

func NewSlateThemeCommand(app *App) *SlateThemeCommand {
	return &SlateThemeCommand{
		app:    app,
		option: palette.NewOption("Theme Slate", "Theme: Slate", []key.Key{}),
	}
}

func (c *SlateThemeCommand) Option() *palette.Option {
	return c.option
}

func (c *SlateThemeCommand) Match(key key.Key) bool {
	return false
}

func (c *SlateThemeCommand) Run() {
	c.app.setColors(theme.Slate{})

	c.app.Render()
}
