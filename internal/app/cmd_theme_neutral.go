package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type NeutralThemeCommand struct {
	app    *App
	option *palette.Option
}

func NewNeutralThemeCommand(app *App) *NeutralThemeCommand {
	return &NeutralThemeCommand{
		app:    app,
		option: palette.NewOption("Theme Neutral", "Theme: Neutral", []key.Key{}),
	}
}

func (c *NeutralThemeCommand) Option() *palette.Option {
	return c.option
}

func (c *NeutralThemeCommand) Match(key key.Key) bool {
	return false
}

func (c *NeutralThemeCommand) Run() {
	c.app.setColors(theme.Neutral{})

	c.app.Render()
}
