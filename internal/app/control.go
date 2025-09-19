package app

import (
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

func (*App) IsEnabled() bool {
	return true
}

func (*App) ToggleEnabled() {
}

func (*App) SetEnabled(bool) {
}

func (app *App) Area() ui.Area {
	return app.area
}

func (app *App) Layout(a ui.Area) {
	app.area = a

	app.header.Layout(app.area)
	app.footer.Layout(app.area)
	if app.zenEnabled {
		app.editor.Layout(app.area)
	} else {
		app.editor.Layout(ui.Area{Y: a.Y + 1, X: a.X, W: a.W, H: a.H - 2})
	}
	app.debug.Layout(app.editor.Area())
	app.palette.Layout(app.editor.Area())
}

func (app *App) Render() {
	vt.Sync.Bsu()

	app.header.Render()
	app.footer.Render()
	app.editor.Render()
	app.debug.Render()
	app.palette.Render()

	vt.Sync.Esu()
}
