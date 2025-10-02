package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type ExitCommand struct {
	app    *App
	option palette.Option
}

func NewExitCommand(app *App) *ExitCommand {
	return &ExitCommand{
		app:    app,
		option: palette.NewOption("Exit", "Global: Exit", []key.Key{{Name: "F10"}}),
	}
}

func (c *ExitCommand) Option() *palette.Option {
	return &c.option
}

func (c *ExitCommand) Match(k key.Key) bool {
	return k.Name == "F10"
}

func (c *ExitCommand) Run() {
	c.app.editor.Enabled = false

	if !c.app.editor.History.IsEmpty() {
		save := make(chan bool)
		go c.app.ask.Open("Save changes?", save)

		if <-save {
			c.app.save()
		}
	}

	c.app.exit()
}
