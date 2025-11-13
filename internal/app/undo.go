package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Undo struct {
	app    *App
	option palette.Option
}

func NewUndo(app *App) *Undo {
	return &Undo{
		app: app,
		option: palette.NewOption(
			"Undo",
			"Edit: Undo",
			[]key.Key{{Name: "z", Mods: key.Ctrl}, {Name: "z", Mods: key.Super}},
		),
	}
}

func (u *Undo) Option() *palette.Option {
	return &u.option
}

func (u *Undo) Match(key.Key) bool {
	return false
}

func (u *Undo) Run() {
	u.app.editor.Undo()
}
