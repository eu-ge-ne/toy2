package editor

import (
	"io"
	"math"
	"os"
	"slices"
	"time"
	"unicode/utf8"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/data"
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

	buffer  *textbuf.TextBuf
	cursor  *cursor.Cursor
	history *history.History
	syntax  *syntax.Syntax
	data    *data.Data
	render  *render.Render
}

func New(multiLine bool) *Editor {
	b := textbuf.New()
	c := cursor.New(&b)
	h := history.New(&b, c)
	d := data.New(multiLine, &b, c, h)
	r := render.New(&b, c)

	ed := Editor{
		multiLine: multiLine,
		buffer:    &b,
		cursor:    c,
		history:   h,
		data:      d,
		render:    r,
	}

	ed.history.OnChanged = ed.OnChanged

	return &ed
}

func (ed *Editor) SetColors(t theme.Tokens) {
	ed.render.SetColors(t)
}

func (ed *Editor) SetSyntax() {
	ed.syntax = syntax.New(ed.buffer)
	ed.syntax.Reset()

	ed.data.SetSyntax(ed.syntax)
}

func (ed *Editor) Layout(a ui.Area) {
	ed.render.SetArea(a)
	ed.data.SetPageSize(a.H)
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

	ed.data.SetEnabled(enabled)
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

	i := slices.IndexFunc(ed.data.Handlers, func(h data.Handler) bool {
		return h.Match(key)
	})

	if i < 0 {
		return false
	}

	r := ed.data.Handlers[i].Handle(key)

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

func (ed *Editor) HasChanges() bool {
	return !ed.history.IsEmpty()
}

func (ed *Editor) SetText(text string) {
	ed.buffer.Reset(text)
	ed.syntax.Reset()
}

func (ed *Editor) GetText() string {
	return ed.buffer.All()
}

func (ed *Editor) Copy() bool {
	return ed.data.Copy()
}

func (ed *Editor) Cut() bool {
	return ed.data.Cut()
}

func (ed *Editor) Paste() bool {
	return ed.data.Paste()
}

func (ed *Editor) Redo() bool {
	return ed.data.Redo()
}

func (ed *Editor) Undo() bool {
	return ed.data.Undo()
}

func (ed *Editor) SelectAll() bool {
	return ed.data.SelectAll()
}
