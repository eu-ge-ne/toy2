package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type ZincThemeCommand struct {
	app    *App
	option *palette.Option
}

func NewZincThemeCommand(app *App) *ZincThemeCommand {
	return &ZincThemeCommand{
		app:    app,
		option: palette.NewOption("Theme Zinc", "Theme: Zinc", []key.Key{}),
	}
}

func (c *ZincThemeCommand) Option() *palette.Option {
	return c.option
}

func (c *ZincThemeCommand) Match(key.Key) bool {
	return false
}

func (c *ZincThemeCommand) Run() {
	c.app.setColors(theme.Zinc{})

	c.app.Render()
}
