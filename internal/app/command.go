package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Command interface {
	Option() *palette.Option
	Match(key.Key) bool
	Run()
}

func (app *App) Undo() {
	if app.editor.Handlers["UNDO"].Run(key.Key{}) {
		app.editor.Render()
	}
}

func (app *App) Whitespace() {
	app.editor.ToggleWhitespaceEnabled()

	app.Render()
}

func (app *App) Wrap() {
	app.editor.ToggleWrapEnabled()

	app.Render()
}

func (app *App) Zen() {
	app.enableZen(!app.zenEnabled)

	app.refresh()
}
