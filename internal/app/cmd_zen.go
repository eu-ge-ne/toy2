package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type ZenCommand struct {
	app    *App
	option *palette.Option
}

func NewZenCommand(app *App) *ZenCommand {
	return &ZenCommand{
		app:    app,
		option: palette.NewOption("Zen", "Global: Toggle Zen Mode", []key.Key{{Name: "F11"}}),
	}
}

func (c *ZenCommand) Option() *palette.Option {
	return c.option
}

func (c *ZenCommand) Match(key *key.Key) bool {
	return key.Name == "F11"
}

func (c *ZenCommand) Run() {
	c.app.enableZen(!c.app.zenEnabled)

	c.app.refresh()
}
