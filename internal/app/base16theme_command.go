package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type Base16ThemeCommand struct {
	app    *App
	option *palette.Option
}

func NewBase16ThemeCommand(app *App) *Base16ThemeCommand {
	return &Base16ThemeCommand{
		app:    app,
		option: palette.NewOption("Theme Base16", "Theme: Base16", []key.Key{}),
	}
}

func (c *Base16ThemeCommand) Option() *palette.Option {
	return c.option
}

func (c *Base16ThemeCommand) Match(key key.Key) bool {
	return false
}

func (c *Base16ThemeCommand) Run() {
	c.app.setColors(theme.Base16{})

	c.app.Render()
}
