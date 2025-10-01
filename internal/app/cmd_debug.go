package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type DebugCommand struct {
	app    *App
	option *palette.Option
}

func NewDebugCommand(app *App) *DebugCommand {
	return &DebugCommand{
		app:    app,
		option: palette.NewOption("Debug", "Global: Toggle Debug Panel", []key.Key{}),
	}
}

func (c *DebugCommand) Option() *palette.Option {
	return c.option
}

func (c *DebugCommand) Match(key *key.Key) bool {
	return false
}

func (c *DebugCommand) Run() {
	c.app.debug.Enabled = !c.app.debug.Enabled

	c.app.editor.Render()
}
