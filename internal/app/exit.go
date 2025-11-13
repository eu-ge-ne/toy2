package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Exit struct {
	app    *App
	option palette.Option
}

func NewExit(app *App) *Exit {
	return &Exit{
		app:    app,
		option: palette.NewOption("Exit", "Global: Exit", []key.Key{{Name: "F10"}}),
	}
}

func (e *Exit) Option() *palette.Option {
	return &e.option
}

func (e *Exit) Match(k key.Key) bool {
	return k.Name == "F10"
}

func (e *Exit) Run() {
	e.app.editor.SetEnabled(false)

	if e.app.editor.HasChanges() {
		if <-e.app.ask.Open("Save changes?") {
			e.app.save()
		}
	}

	e.app.exit()
}
