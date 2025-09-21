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

func (d *DebugCommand) Option() *palette.Option {
	return d.option
}

func (d *DebugCommand) Match(key key.Key) bool {
	return false
}

func (d *DebugCommand) Run() {
	d.app.debug.ToggleEnabled()

	d.app.editor.Render()
}
