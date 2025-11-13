package app

import (
	"github.com/eu-ge-ne/toy2/internal/grammar/typescript"
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type SyntaxTypeScript struct {
	app    *App
	option palette.Option
}

func NewSyntaxTypeScript(app *App) *SyntaxTypeScript {
	return &SyntaxTypeScript{
		app:    app,
		option: palette.NewOption("Syntax TypeScript", "Syntax: TypeScript", []key.Key{}),
	}
}

func (s *SyntaxTypeScript) Option() *palette.Option {
	return &s.option
}

func (s *SyntaxTypeScript) Match(key.Key) bool {
	return false
}

func (s *SyntaxTypeScript) Run() {
	s.app.editor.SetGrammar(typescript.Grammar)
}
