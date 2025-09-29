package app

import (
	"flag"
	"os"
	"os/signal"
	"runtime/pprof"
	"slices"
	"strings"
	"syscall"

	"github.com/eu-ge-ne/toy2/internal/alert"
	"github.com/eu-ge-ne/toy2/internal/ask"
	"github.com/eu-ge-ne/toy2/internal/debug"
	"github.com/eu-ge-ne/toy2/internal/editor"
	"github.com/eu-ge-ne/toy2/internal/file"
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
	commands   []Command
	restoreVt  func()
	zenEnabled bool
	filePath   string
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func New() *App {
	app := App{}

	app.commands = []Command{
		NewDebugCommand(&app),
		NewExitCommand(&app),
		NewPaletteCommand(&app),
		NewRedoCommand(&app),
		NewSaveCommand(&app),
		NewSelectAllCommand(&app),
		NewUndoCommand(&app),
		NewWhitespaceCommand(&app),
		NewWrapCommand(&app),
		NewZenCommand(&app),
		NewBase16ThemeCommand(&app),
		NewGrayThemeCommand(&app),
		NewNeutralThemeCommand(&app),
		NewSlateThemeCommand(&app),
		NewStoneThemeCommand(&app),
		NewZincThemeCommand(&app),
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
	app.editor = editor.New(true)
	app.footer = footer.New()
	app.debug = debug.New()
	app.palette = palette.New(&app, options)
	app.saveas = saveas.New()

	app.editor.Enabled = true
	app.editor.OnCursor = app.footer.SetCursorStatus
	app.editor.OnKeyHandled = app.debug.SetInputTime
	app.editor.OnRender = app.debug.SetRenderTime
	app.editor.History.OnChanged = func() {
		app.header.SetFlag(!app.editor.History.IsEmpty())
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

	app.setColors(theme.Neutral{})
	app.enableZen(false)
	app.editor.WhitespaceEnabled = true
	app.editor.WrapEnabled = true

	app.debug.Enabled = true

	app.refresh()

	go app.listenSigwinch()

	if flag.NArg() > 0 {
		app.open(flag.Arg(0))
	}

	app.editor.Reset(true)
	app.editor.Render()

	app.processInput()
}

func (app *App) Area() ui.Area {
	return app.area
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
	if app.zenEnabled {
		app.editor.Layout(app.area)
	} else {
		app.editor.Layout(ui.Area{
			Y: a.Y + 1,
			X: a.X,
			W: a.W,
			H: a.H - 2,
		})
	}
	app.debug.Layout(app.editor.Area())
	app.palette.Layout(app.editor.Area())
	app.ask.Layout(app.editor.Area())
	app.alert.Layout(app.editor.Area())
	app.saveas.Layout(app.editor.Area())
}

func (app *App) setColors(t theme.Tokens) {
	app.ask.SetColors(t)
	app.alert.SetColors(t)
	app.debug.SetColors(t)
	app.header.SetColors(t)
	app.footer.SetColors(t)
	app.editor.SetColors(t)
	app.palette.SetColors(t)
	app.saveas.SetColors(t)
}

func (app *App) enableZen(enabled bool) {
	app.zenEnabled = enabled

	app.header.Enabled = !enabled
	app.footer.Enabled = !enabled
	app.editor.IndexEnabled = !enabled
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
	for {
		for key := range vt.Read() {
			i := slices.IndexFunc(app.commands, func(c Command) bool {
				return c.Match(key)
			})

			if i >= 0 {
				app.commands[i].Run()
				continue
			}

			if app.editor.Enabled {
				if app.editor.HandleKey(key) {
					app.editor.Render()
				}
			}
		}
	}
}

func (app *App) open(filePath string) {
	err := file.Load(app.editor.Buffer, filePath)

	if os.IsNotExist(err) {
		return
	}

	if err != nil {
		done := make(chan struct{})
		go app.alert.Open(err.Error(), done)
		<-done

		app.exit()
	}

	app.setFilePath(filePath)
}

func (app *App) save() bool {
	if len(app.filePath) != 0 {
		return app.saveFile()
	} else {
		return app.saveFileAs()
	}
}

func (app *App) saveFile() bool {
	err := file.Save(app.editor.Buffer, app.filePath)
	if err == nil {
		return true
	}

	done := make(chan struct{})
	go app.alert.Open(err.Error(), done)
	<-done

	return app.saveFileAs()
}

func (app *App) saveFileAs() bool {
	for {
		filePathResult := make(chan string)
		go app.saveas.Open(app.filePath, filePathResult)

		filePath := <-filePathResult
		if len(filePath) == 0 {
			return false
		}

		err := file.Save(app.editor.Buffer, filePath)
		if err == nil {
			app.setFilePath(filePath)
			return true
		}

		done := make(chan struct{})
		go app.alert.Open(err.Error(), done)
		<-done
	}
}

func (app *App) setFilePath(filePath string) {
	app.filePath = filePath

	app.header.SetFilePath(filePath)
}
