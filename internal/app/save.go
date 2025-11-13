package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Save struct {
	app    *App
	option palette.Option
}

func NewSave(app *App) *Save {
	return &Save{
		app:    app,
		option: palette.NewOption("Save", "Global: Save", []key.Key{{Name: "F2"}}),
	}
}

func (s *Save) Option() *palette.Option {
	return &s.option
}

func (s *Save) Match(k key.Key) bool {
	return k.Name == "F2"
}

func (s *Save) Run() {
	s.app.editor.SetEnabled(false)

	if s.app.save() {
		//app.editor.Data.TopHome(false)
	}

	s.app.editor.SetEnabled(true)
}
