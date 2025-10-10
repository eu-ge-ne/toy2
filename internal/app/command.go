package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

type Command interface {
	Option() *palette.Option
	Match(key.Key) bool
	Run()
}

func (app *App) ThemeSlate() {
	app.setColors(theme.Slate{})

	app.Render()
}

func (app *App) ThemeStone() {
	app.setColors(theme.Stone{})

	app.Render()
}

func (app *App) ThemeZinc() {
	app.setColors(theme.Zinc{})

	app.Render()
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
