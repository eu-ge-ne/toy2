package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type SyntaxOff struct {
	app    *App
	option palette.Option
}

func NewSyntaxOff(app *App) *SyntaxOff {
	return &SyntaxOff{
		app:    app,
		option: palette.NewOption("Syntax Off", "Syntax: Off", []key.Key{}),
	}
}

func (s *SyntaxOff) Option() *palette.Option {
	return &s.option
}

func (s *SyntaxOff) Match(key.Key) bool {
	return false
}

func (s *SyntaxOff) Run() {
	s.app.editor.SetGrammar(nil)
}
