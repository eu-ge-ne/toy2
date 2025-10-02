package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type StoneThemeCommand struct {
	app    *App
	option *palette.Option
}

func NewStoneThemeCommand(app *App) *StoneThemeCommand {
	return &StoneThemeCommand{
		app:    app,
		option: palette.NewOption("Theme Stone", "Theme: Stone", []key.Key{}),
	}
}

func (c *StoneThemeCommand) Option() *palette.Option {
	return c.option
}

func (c *StoneThemeCommand) Match(key.Key) bool {
	return false
}

func (c *StoneThemeCommand) Run() {
	c.app.setColors(theme.Stone{})

	c.app.Render()
}
