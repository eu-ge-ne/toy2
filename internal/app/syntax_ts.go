package app

import (
	"github.com/eu-ge-ne/toy2/internal/grammar/ts"
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type SyntaxTS struct {
	app    *App
	option palette.Option
}

func NewSyntaxTS(app *App) *SyntaxTS {
	return &SyntaxTS{
		app:    app,
		option: palette.NewOption("Syntax TS", "Syntax: TS", []key.Key{}),
	}
}

func (c *SyntaxTS) Option() *palette.Option {
	return &c.option
}

func (c *SyntaxTS) Match(key.Key) bool {
	return false
}

func (c *SyntaxTS) Run() bool {
	c.app.editor.SetGrammar(ts.TS)

	return true
}
