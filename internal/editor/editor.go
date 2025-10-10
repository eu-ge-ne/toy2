package editor

import (
	"io"
	"math"
	"os"
	"slices"
	"time"
	"unicode/utf8"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/handler"
	"github.com/eu-ge-ne/toy2/internal/editor/history"
	"github.com/eu-ge-ne/toy2/internal/editor/render"
	"github.com/eu-ge-ne/toy2/internal/editor/syntax"
	"github.com/eu-ge-ne/toy2/internal/key"
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
	ed.render.SetColors(t)
}

func (ed *Editor) SetSyntax() {
	ed.syntax = syntax.New(ed.buffer)

	ed.syntax.Reset()
}

func (ed *Editor) Layout(a ui.Area) {
	ed.render.SetArea(a)
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

func (ed *Editor) SetEnabled(enabled bool) {
	ed.enabled = enabled
	ed.render.SetEnabled(enabled)
}

func (ed *Editor) SetIndexEnabled(enabled bool) {
	ed.render.SetIndexEnabled(enabled)
}

func (ed *Editor) EnableWhitespace(enabled bool) {
	ed.render.SetWhitespaceEnabled(enabled)
}

func (ed *Editor) ToggleWhitespaceEnabled() {
	ed.render.ToggleWhitespaceEnabled()

	ed.cursor.Home(false)
}

func (ed *Editor) SetWrapEnabled(enabled bool) {
	ed.render.SetWrapEnabled(enabled)
}

func (ed *Editor) ToggleWrapEnabled() {
	ed.render.ToggleWrapEnabled()

	ed.cursor.Home(false)
}

func (ed *Editor) HandleKey(key key.Key) bool {
	if !ed.enabled {
		return false
	}

	t0 := time.Now()

	i := slices.IndexFunc(ed.handlers, func(h handler.Handler) bool {
		return h.Match(key)
	})

	if i < 0 {
		return false
	}

	r := ed.handlers[i].Handle(key)

	if ed.OnKeyHandled != nil {
		ed.OnKeyHandled(time.Since(t0))
	}

	return r
}

func (ed *Editor) Load(filePath string) error {
	f, err := os.Open(filePath)
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
			return err
		}

		chunk := buf[:bytesRead]

		if !utf8.Valid(chunk) {
			panic("invalid utf8 chunk")
		}

		ed.buffer.Append(string(chunk))
	}

	ed.syntax.Reset()

	return nil
}

func (ed *Editor) Save(filePath string) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	for text := range ed.buffer.Iter() {
		_, err := f.WriteString(text)
		if err != nil {
			return err
		}
	}

	return nil
}
