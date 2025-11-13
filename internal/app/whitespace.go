package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Whitespace struct {
	app    *App
	option palette.Option
}

func NewWhitespace(app *App) *Whitespace {
	return &Whitespace{
		app:    app,
		option: palette.NewOption("Whitespace", "View: Toggle Render Whitespace", []key.Key{{Name: "F5"}}),
	}
}

func (w *Whitespace) Option() *palette.Option {
	return &w.option
}

func (w *Whitespace) Match(k key.Key) bool {
	return k.Name == "F5"
}

func (w *Whitespace) Run() {
	w.app.editor.ToggleWhitespaceEnabled()
}
