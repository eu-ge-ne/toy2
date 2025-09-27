package app

import (
	"math"

	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type SelectAllCommand struct {
	app    *App
	option *palette.Option
}

func NewSelectAllCommand(app *App) *SelectAllCommand {
	return &SelectAllCommand{
		app:    app,
		option: palette.NewOption("Select All", "Edit: Select All", []key.Key{{Name: "a", Ctrl: true}, {Name: "a", Super: true}}),
	}
}

func (c *SelectAllCommand) Option() *palette.Option {
	return c.option
}

func (c *SelectAllCommand) Match(key key.Key) bool {
	return false
}

func (c *SelectAllCommand) Run() {
	if !c.app.editor.Enabled {
		return
	}

	c.app.editor.Cursor.Set(0, 0, false)
	c.app.editor.Cursor.Set(math.MaxInt, math.MaxInt, true)

	c.app.editor.Render()
}
