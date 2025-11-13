package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Paste struct {
	app    *App
	option palette.Option
}

func NewPaste(app *App) *Paste {
	return &Paste{
		app: app,
		option: palette.NewOption(
			"Paste",
			"Edit: Paste",
			[]key.Key{{Name: "v", Mods: key.Ctrl}, {Name: "v", Mods: key.Super}},
		),
	}
}

func (p *Paste) Option() *palette.Option {
	return &p.option
}

func (p *Paste) Match(key.Key) bool {
	return false
}

func (p *Paste) Run() {
	p.app.editor.Paste()
}
