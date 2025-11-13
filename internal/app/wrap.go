package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Wrap struct {
	app    *App
	option palette.Option
}

func NewWrap(app *App) *Wrap {
	return &Wrap{
		app:    app,
		option: palette.NewOption("Wrap", "View: Toggle Line Wrap", []key.Key{{Name: "F6"}}),
	}
}

func (w *Wrap) Option() *palette.Option {
	return &w.option
}

func (w *Wrap) Match(k key.Key) bool {
	return k.Name == "F6"
}

func (w *Wrap) Run() {
	w.app.editor.ToggleWrapEnabled()
}
