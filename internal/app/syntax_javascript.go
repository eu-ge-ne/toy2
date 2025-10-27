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

func (c *SyntaxJavaScript) Option() *palette.Option {
	return &c.option
}

func (c *SyntaxJavaScript) Match(key.Key) bool {
	return false
}

func (c *SyntaxJavaScript) Run() bool {
	c.app.editor.SetGrammar(javascript.Grammar)

	return true
}
