package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Redo struct {
	app    *App
	option palette.Option
}

func NewRedo(app *App) *Redo {
	return &Redo{
		app: app,
		option: palette.NewOption(
			"Redo",
			"Edit: Redo",
			[]key.Key{{Name: "y", Mods: key.Ctrl}, {Name: "y", Mods: key.Super}},
		),
	}
}

func (r *Redo) Option() *palette.Option {
	return &r.option
}

func (r *Redo) Match(key.Key) bool {
	return false
}

func (r *Redo) Run() {
	r.app.editor.Redo()
}
