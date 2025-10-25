package app

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"slices"
	"strings"
	"syscall"

	"github.com/eu-ge-ne/toy2/internal/alert"
	"github.com/eu-ge-ne/toy2/internal/ask"
	"github.com/eu-ge-ne/toy2/internal/debug"
	"github.com/eu-ge-ne/toy2/internal/editor"
	"github.com/eu-ge-ne/toy2/internal/footer"
	"github.com/eu-ge-ne/toy2/internal/header"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/saveas"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type App struct {
	area       ui.Area
	alert      *alert.Alert
	ask        *ask.Ask
	debug      *debug.Debug
	editor     *editor.Editor
	footer     *footer.Footer
	header     *header.Header
	palette    *palette.Palette
	saveas     *saveas.SaveAs
	commands   map[string]Command
	restoreVt  func()
	zenEnabled bool
	filePath   string
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func New() *App {
	app := App{}

	app.commands = map[string]Command{
		"COPY":          NewCopy(&app),
		"CUT":           NewCut(&app),
		"DEBUG":         NewDebug(&app),
		"EXIT":          NewExit(&app),
		"PALETTE":       NewPalette(&app),
		"PASTE":         NewPaste(&app),
		"REDO":          NewRedo(&app),
		"SAVE":          NewSave(&app),
		"SELECTALL":     NewSelectAll(&app),
		"SYNTAX_OFF":    NewSyntaxOff(&app),
		"SYNTAX_TS":     NewSyntaxTS(&app),
		"THEME_BASE16":  NewThemeBase16(&app),
		"THEME_GRAY":    NewThemeGray(&app),
		"THEME_NEUTRAL": NewThemeNeutral(&app),
		"THEME_SLATE":   NewThemeSlate(&app),
		"THEME_STONE":   NewThemeStone(&app),
		"THEME_ZINC":    NewThemeZinc(&app),
		"UNDO":          NewUndo(&app),
		"WHITESPACE":    NewWhitespace(&app),
		"WRAP":          NewWrap(&app),
		"ZEN":           NewZen(&app),
	}

	options := []*palette.Option{}
	for _, c := range app.commands {
		opt := c.Option()
		if opt != nil {
			options = append(options, opt)
		}
	}
	slices.SortFunc(options, func(a, b *palette.Option) int {
		return strings.Compare(strings.ToLower(a.Description), strings.ToLower(b.Description))
	})

	app.ask = ask.New()
	app.alert = alert.New()
	app.header = header.New()
	app.footer = footer.New()
	app.debug = debug.New()
	app.palette = palette.New(&app, options)
	app.saveas = saveas.New()

	app.editor = editor.New(true)
	app.editor.OnCursor = app.footer.SetCursorStatus
	app.editor.OnKeyHandled = app.debug.SetInputTime
	app.editor.OnRender = app.debug.SetRenderTime
	app.editor.OnChanged = func() {
		app.header.SetFlag(app.editor.HasChanges())
	}

	return &app
}

func (app *App) Run() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	app.restoreVt = vt.Init()

	app.setColors(theme.Zinc{})
	app.setZenEnabled(false)

	app.editor.SetEnabled(true)
	app.editor.EnableWhitespace(true)
	app.editor.SetWrapEnabled(true)
	app.editor.SetSyntax(nil)

	app.debug.Enable(true)

	app.refresh()

	go app.listenSigwinch()

	if flag.NArg() > 0 {
		app.open(flag.Arg(0))
	}

	app.editor.Render()

	vt.ListenStdin()

	app.processInput()
}

func (app *App) Render() {
	vt.Sync.Bsu()

	app.header.Render()
	app.footer.Render()
	app.editor.Render()
	app.debug.Render()
	app.palette.Render()
	app.ask.Render()
	app.alert.Render()
	app.saveas.Render()

	vt.Sync.Esu()
}

func (app *App) layout(a ui.Area) {
	app.area = a

	app.header.Layout(app.area)
	app.footer.Layout(app.area)

	var editorArea ui.Area
	if app.zenEnabled {
		editorArea = app.area
	} else {
		editorArea = ui.Area{
			Y: a.Y + 1,
			X: a.X,
			W: a.W,
			H: a.H - 2,
		}
	}

	app.editor.Layout(editorArea)
	app.debug.Layout(editorArea)
	app.palette.Layout(editorArea)
	app.ask.Layout(editorArea)
	app.alert.Layout(editorArea)
	app.saveas.Layout(editorArea)
}

func (app *App) setColors(t theme.Theme) {
	app.ask.SetColors(t)
	app.alert.SetColors(t)
	app.debug.SetColors(t)
	app.header.SetColors(t)
	app.footer.SetColors(t)
	app.editor.SetColors(t)
	app.palette.SetColors(t)
	app.saveas.SetColors(t)
}

func (app *App) setZenEnabled(enabled bool) {
	app.zenEnabled = enabled

	app.header.Enable(!enabled)
	app.footer.Enable(!enabled)
	app.editor.SetIndexEnabled(!enabled)
}

func (app *App) exit() {
	app.restoreVt()

	os.Exit(0)
}

func (app *App) listenSigwinch() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGWINCH)

	for {
		<-c

		app.refresh()
	}
}

func (app *App) refresh() {
	w, h, err := vt.GetSize()
	if err != nil {
		panic(err)
	}

	app.layout(ui.Area{X: 0, Y: 0, W: w, H: h})
	app.Render()
}

func (app *App) processInput() {
outer:
	for {
		key := vt.ReadKey()

		for _, c := range app.commands {
			if c.Match(key) {
				if c.Run() {
					app.Render()
				}
				continue outer
			}
		}

		if app.editor.HandleKey(key) {
			app.editor.Render()
		}
	}
}

func (app *App) open(filePath string) {
	err := app.editor.Load(filePath)

	if os.IsNotExist(err) {
		return
	}

	if err != nil {
		<-app.alert.Open(err.Error())
		app.exit()
	}

	app.setFilePath(filePath)

	runtime.GC()
}

func (app *App) save() bool {
	if len(app.filePath) != 0 {
		return app.saveFile()
	} else {
		return app.saveFileAs()
	}
}

func (app *App) saveFile() bool {
	err := app.editor.Save(app.filePath)
	if err == nil {
		return true
	}

	<-app.alert.Open(err.Error())

	return app.saveFileAs()
}

func (app *App) saveFileAs() bool {
	for {
		filePath := <-app.saveas.Open(app.filePath)
		if len(filePath) == 0 {
			return false
		}

		err := app.editor.Save(filePath)
		if err == nil {
			app.setFilePath(filePath)
			return true
		}

		<-app.alert.Open(err.Error())
	}
}

func (app *App) setFilePath(filePath string) {
	app.filePath = filePath

	app.header.SetFilePath(filePath)
}
