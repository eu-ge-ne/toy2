package app

import (
	"slices"

	"github.com/eu-ge-ne/toy2/internal/app/command"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

func (app *App) Copy() {
	if app.editor.Copy() {
		app.editor.Render()
	}
}

func (app *App) Cut() {
	if app.editor.Cut() {
		app.editor.Render()
	}
}

func (app *App) Debug() {
	app.debug.ToggleEnabled()

	app.editor.Render()
}

func (app *App) Exit() {
	app.editor.Enable(false)

	if app.editor.HasChanges() {
		save := make(chan bool)
		go app.ask.Open("Save changes?", save)

		if <-save {
			app.save()
		}
	}

	app.exit()
}

func (app *App) Palette() {
	app.editor.Enable(false)

	done := make(chan *palette.Option)

	go app.palette.Open(done)

	option := <-done

	app.editor.Enable(true)

	app.editor.Render()

	if option != nil {
		i := slices.IndexFunc(app.commands, func(c command.Command) bool {
			o := c.Option()
			return o != nil && o.Id == option.Id
		})

		if i >= 0 {
			app.commands[i].Run()
		}
	}
}

func (app *App) Paste() {
	if app.editor.Paste() {
		app.editor.Render()
	}
}

func (app *App) Redo() {
	if app.editor.Redo() {
		app.editor.Render()
	}
}

func (app *App) Save() {
	app.editor.Enable(false)

	if app.save() {
		app.editor.ResetCursor()
	}

	app.editor.Enable(true)

	app.editor.Render()
}

func (app *App) SelectAll() {
	if app.editor.SelectAll() {
		app.editor.Render()
	}
}

func (app *App) ThemeBase16() {
	app.setColors(theme.Base16{})

	app.Render()
}

func (app *App) ThemeGray() {
	app.setColors(theme.Gray{})

	app.Render()
}

func (app *App) ThemeNeutral() {
	app.setColors(theme.Neutral{})

	app.Render()
}

func (app *App) ThemeSlate() {
	app.setColors(theme.Slate{})

	app.Render()
}

func (app *App) ThemeStone() {
	app.setColors(theme.Stone{})

	app.Render()
}

func (app *App) ThemeZinc() {
	app.setColors(theme.Zinc{})

	app.Render()
}

func (app *App) Undo() {
	if app.editor.Undo() {
		app.editor.Render()
	}
}

func (app *App) Whitespace() {
	app.editor.ToggleWhitespace()

	app.Render()
}

func (app *App) Wrap() {
	app.editor.ToggleWrap()

	app.Render()
}

func (app *App) Zen() {
	app.enableZen(!app.zenEnabled)

	app.refresh()
}
