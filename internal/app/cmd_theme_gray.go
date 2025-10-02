package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type GrayThemeCommand struct {
	app    *App
	option *palette.Option
}

func NewGrayThemeCommand(app *App) *GrayThemeCommand {
	return &GrayThemeCommand{
		app:    app,
		option: palette.NewOption("Theme Gray", "Theme: Gray", []key.Key{}),
	}
}

func (c *GrayThemeCommand) Option() *palette.Option {
	return c.option
}

func (c *GrayThemeCommand) Match(key.Key) bool {
	return false
}

func (c *GrayThemeCommand) Run() {
	c.app.setColors(theme.Gray{})

	c.app.Render()
}
