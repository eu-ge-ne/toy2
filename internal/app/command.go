package app

import (
	"slices"

	"github.com/eu-ge-ne/toy2/internal/app/command"
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
)

func (app *App) Copy() {
	if app.editor.Actions["COPY"].Run(key.Key{}) {
		app.editor.Render()
	}
}

func (app *App) Cut() {
	if app.editor.Actions["CUT"].Run(key.Key{}) {
		app.editor.Render()
	}
}

func (app *App) Debug() {
	app.debug.ToggleEnabled()

	app.editor.Render()
}

func (app *App) Exit() {
	app.editor.SetEnabled(false)

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
	app.editor.SetEnabled(false)

	done := make(chan *palette.Option)

	go app.palette.Open(done)

	option := <-done

	app.editor.SetEnabled(true)

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
	if app.editor.Actions["PASTE"].Run(key.Key{}) {
		app.editor.Render()
	}
}

func (app *App) Redo() {
	if app.editor.Actions["REDO"].Run(key.Key{}) {
		app.editor.Render()
	}
}

func (app *App) Save() {
	app.editor.SetEnabled(false)

	if app.save() {
		//app.editor.Data.TopHome(false)
	}

	app.editor.SetEnabled(true)

	app.editor.Render()
}

func (app *App) SelectAll() {
	if app.editor.Actions["SELECTALL"].Run(key.Key{}) {
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
	if app.editor.Actions["UNDO"].Run(key.Key{}) {
		app.editor.Render()
	}
}

func (app *App) Whitespace() {
	app.editor.ToggleWhitespaceEnabled()

	app.Render()
}

func (app *App) Wrap() {
	app.editor.ToggleWrapEnabled()

	app.Render()
}

func (app *App) Zen() {
	app.enableZen(!app.zenEnabled)

	app.refresh()
}
