package editor

import (
	"math"
	"time"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/handler"
	"github.com/eu-ge-ne/toy2/internal/editor/history"
	"github.com/eu-ge-ne/toy2/internal/editor/render"
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
	enabled   bool
	clipboard string
	pageSize  int

	buffer   *textbuf.TextBuf
	cursor   *cursor.Cursor
	history  *history.History
	syntax   *syntax.Syntax
	render   *render.Render
	handlers []handler.Handler
}

func New(multiLine bool) *Editor {
	b := textbuf.New()
	c := cursor.New(&b)
	h := history.New(&b, &c)
	r := render.New(&b, &c)

	ed := Editor{
		multiLine: multiLine,
		buffer:    &b,
		cursor:    &c,
		history:   &h,
		render:    &r,
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
	ed.render.Colors = render.NewColors(t)
}

func (ed *Editor) SetSyntax() {
	ed.syntax = syntax.New(ed.buffer)

	ed.syntax.Reset()
}

func (ed *Editor) Layout(a ui.Area) {
	ed.render.Area = a
	ed.pageSize = a.H
}

func (ed *Editor) Render() {
	started := time.Now()

	if ed.OnCursor != nil {
		ed.OnCursor(ed.cursor.Ln, ed.cursor.Col, ed.buffer.LineCount())
	}

	ed.render.Render()

	if ed.OnRender != nil {
		ed.OnRender(time.Since(started))
	}
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
	ed.render.Enabled = enable
}

func (ed *Editor) EnableIndex(enable bool) {
	ed.render.IndexEnabled = enable
}

func (ed *Editor) EnableWhitespace(enable bool) {
	ed.render.WhitespaceEnabled = enable
}

func (ed *Editor) ToggleWhitespace() {
	ed.render.WhitespaceEnabled = !ed.render.WhitespaceEnabled

	ed.cursor.Home(false)
}

func (ed *Editor) EnableWrap(enable bool) {
	ed.render.WrapEnabled = enable
}

func (ed *Editor) ToggleWrap() {
	ed.render.WrapEnabled = !ed.render.WrapEnabled

	ed.cursor.Home(false)
}
