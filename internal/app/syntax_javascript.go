package app

import (
	"github.com/eu-ge-ne/toy2/internal/grammar/javascript"
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type SyntaxJavaScript struct {
	app    *App
	option palette.Option
}

func NewSyntaxJavaScript(app *App) *SyntaxJavaScript {
	return &SyntaxJavaScript{
		app:    app,
		option: palette.NewOption("Syntax JavaScript", "Syntax: JavaScript", []key.Key{}),
	}
}

func (s *SyntaxJavaScript) Option() *palette.Option {
	return &s.option
}

func (s *SyntaxJavaScript) Match(key.Key) bool {
	return false
}

func (s *SyntaxJavaScript) Run() {
	s.app.editor.SetGrammar(javascript.Grammar)
}
