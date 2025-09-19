package app

import (
	"flag"
	"io"
	"os"
	"os/signal"
	"runtime/pprof"
	"slices"
	"strings"
	"syscall"
	"unicode/utf8"

	"github.com/eu-ge-ne/toy2/internal/debug"
	"github.com/eu-ge-ne/toy2/internal/editor"
	"github.com/eu-ge-ne/toy2/internal/footer"
	"github.com/eu-ge-ne/toy2/internal/header"
	"github.com/eu-ge-ne/toy2/internal/palette"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type App struct {
	area      ui.Area
	header    *header.Header
	editor    *editor.Editor
	footer    *footer.Footer
	debug     *debug.Debug
	palette   *palette.Palette
	commands  []Command
	restoreVt func()
}

func New() *App {
	app := App{}

	app.commands = []Command{
		NewBase16ThemeCommand(&app),
		NewDebugCommand(&app),
		NewExitCommand(&app),
		NewGrayThemeCommand(&app),
		NewNeutralThemeCommand(&app),
		NewPaletteCommand(&app),
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

	app.header = header.New()
	app.editor = editor.New(true)
	app.footer = footer.New()
	app.debug = debug.New()
	app.palette = palette.New(&app, options)

	app.editor.SetEnabled(true)
	app.editor.OnCursor = app.footer.SetCursorStatus
	app.editor.OnKeyHandled = app.debug.SetInputTime
	app.editor.OnRender = app.debug.SetRenderTime

	return &app
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

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

	app.SetColors(theme.Neutral{})

	app.refresh()

	go app.listenSigwinch()

	if flag.NArg() > 0 {
		app.openFile(flag.Arg(0))
	}

	app.processInput()
}

func (app *App) SetColors(t theme.Tokens) {
	app.header.SetColors(t)
	app.footer.SetColors(t)
	app.editor.SetColors(t)
	app.debug.SetColors(t)
	app.palette.SetColors(t)

	//set_alert_colors(tokens)
	//set_ask_colors(tokens)
	//set_save_as_colors(tokens)
}

func (app *App) exit() {
	app.restoreVt()

	os.Exit(0)
}

func (app *App) openFile(filePath string) {
	err := app.load(filePath)

	if err != nil {
		panic(err)
	}

	app.editor.Reset(true)
	app.editor.Render()

	app.setFilePath(filePath)
}

func (app *App) load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	buf := make([]byte, 1024*1024*64)

	for {
		bytesRead, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if bytesRead == 0 {
			break
		}

		chunk := buf[:bytesRead]

		if !utf8.Valid(chunk) {
			panic("invalid utf8 chunk")
		}

		app.editor.Buffer.Append(string(chunk))
	}

	return nil
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

			if app.editor.IsEnabled() {
				if app.editor.HandleKey(key) {
					app.editor.Render()
				}
			}
		}
	}
}

func (app *App) setFilePath(filePath string) {
	app.header.SetFilePath(filePath)
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

	app.Layout(ui.Area{X: 0, Y: 0, W: w, H: h})
	app.Render()
}
