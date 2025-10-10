package editor

import (
	"math"
	"time"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/handler"
	"github.com/eu-ge-ne/toy2/internal/editor/history"
	"github.com/eu-ge-ne/toy2/internal/editor/syntax"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
)

type Editor struct {
	OnKeyHandled func(time.Duration)
	OnRender     func(time.Duration)
	OnCursor     func(int, int, int)
	OnChanged    func()

	multiLine bool
	area      ui.Area
	enabled   bool
	clipboard string

	indexEnabled      bool
	whitespaceEnabled bool
	wrapEnabled       bool

	indexWidth int
	textWidth  int
	cursorY    int
	cursorX    int
	scrollLn   int
	scrollCol  int

	buffer   *textbuf.TextBuf
	cursor   *cursor.Cursor
	history  *history.History
	syntax   *syntax.Syntax
	handlers []handler.Handler

	colors Colors
}

func New(multiLine bool, withSyntax bool) *Editor {
	b := textbuf.New()
	c := cursor.New(&b)
	h := history.New(&b, &c)

	ed := Editor{
		multiLine: multiLine,
		buffer:    &b,
		cursor:    &c,
		history:   &h,
	}

	if withSyntax {
		ed.syntax = syntax.New(&b)
	}

	ed.history.OnChanged = ed.OnChanged

	ed.handlers = append(ed.handlers,
		&handler.Insert{Editor: &ed},
		&handler.Backspace{Editor: &ed},
		&handler.Bottom{Editor: &ed},
		&handler.Copy{Editor: &ed},
		&handler.Cut{Editor: &ed},
		&handler.Delete{Editor: &ed},
		&handler.Down{Editor: &ed},
		&handler.End{Editor: &ed},
		&handler.Enter{Editor: &ed},
		&handler.Home{Editor: &ed},
		&handler.Left{Editor: &ed},
		&handler.PageDown{Editor: &ed},
		&handler.PageUp{Editor: &ed},
		&handler.Paste{Editor: &ed},
		&handler.Redo{Editor: &ed},
		&handler.Right{Editor: &ed},
		&handler.SelectAll{Editor: &ed},
		&handler.Top{Editor: &ed},
		&handler.Undo{Editor: &ed},
		&handler.Up{Editor: &ed},
	)

	return &ed
}

func (ed *Editor) SetColors(t theme.Tokens) {
	ed.colors = NewColors(t)
}

func (ed *Editor) Layout(a ui.Area) {
	ed.area = a
}

// TODO: why needed?
func (ed *Editor) ResetCursor() {
	if ed.multiLine {
		ed.cursor.Set(0, 0, false)
	} else {
		ed.cursor.Set(math.MaxInt, math.MaxInt, false)
	}
}

func (ed *Editor) Enable(enable bool) {
	ed.enabled = enable
}

func (ed *Editor) EnableIndex(enable bool) {
	ed.indexEnabled = enable
}

func (ed *Editor) EnableWhitespace(enable bool) {
	ed.whitespaceEnabled = enable
}

func (ed *Editor) ToggleWhitespace() {
	ed.whitespaceEnabled = !ed.whitespaceEnabled

	ed.cursor.Home(false)
}

func (ed *Editor) EnableWrap(enable bool) {
	ed.wrapEnabled = enable
}

func (ed *Editor) ToggleWrap() {
	ed.wrapEnabled = !ed.wrapEnabled

	ed.cursor.Home(false)
}
