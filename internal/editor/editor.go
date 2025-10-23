package editor

import (
	"io"
	"math"
	"os"
	"time"
	"unicode/utf8"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/history"
	"github.com/eu-ge-ne/toy2/internal/editor/render"
	"github.com/eu-ge-ne/toy2/internal/editor/syntax"
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
)

type Editor struct {
	Handlers     map[string]Handler
	OnKeyHandled func(time.Duration)
	OnRender     func(time.Duration)
	OnCursor     func(int, int, int)
	OnChanged    func()

	multiLine bool
	buffer    *textbuf.TextBuf
	cursor    *cursor.Cursor
	history   *history.History
	syntax    *syntax.Syntax
	render    *render.Render

	enabled   bool
	pageSize  int
	clipboard string
}

func New(multiLine bool) *Editor {
	ed := &Editor{
		multiLine: multiLine,
	}

	ed.buffer = textbuf.New()
	ed.cursor = cursor.New(ed.buffer)

	ed.history = history.New(ed.buffer, ed.cursor)
	ed.history.OnChanged = ed.OnChanged

	ed.render = render.New(ed.buffer, ed.cursor)

	ed.Handlers = map[string]Handler{
		"INSERT":    &Insert{ed},
		"BACKSPACE": &Backspace{ed},
		"BOTTOM":    &Bottom{ed},
		"COPY":      &Copy{ed},
		"CUT":       &Cut{ed},
		"DELETE":    &Delete{ed},
		"DOWN":      &Down{ed},
		"END":       &End{ed},
		"ENTER":     &Enter{ed},
		"HOME":      &Home{ed},
		"LEFT":      &Left{ed},
		"PAGEDOWN":  &PageDown{ed},
		"PAGEUP":    &PageUp{ed},
		"PASTE":     &Paste{ed},
		"REDO":      &Redo{ed},
		"RIGHT":     &Right{ed},
		"SELECTALL": &SelectAll{ed},
		"TOP":       &Top{ed},
		"UNDO":      &Undo{ed},
		"UP":        &Up{ed},
	}

	return ed
}

func (ed *Editor) SetSyntax() {
	if ed.syntax != nil {
		ed.syntax.Close()
	}

	s := syntax.New(ed.buffer)

	ed.syntax = s
	ed.render.SetSyntax(s)
}

func (ed *Editor) SetColors(t theme.Theme) {
	ed.render.SetColors(t)
}

func (ed *Editor) Layout(a ui.Area) {
	ed.pageSize = a.H
	ed.render.SetArea(a)
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

	for _, h := range ed.Handlers {
		if h.Match(key) {
			r := h.Run(key)

			if ed.OnKeyHandled != nil {
				ed.OnKeyHandled(time.Since(t0))
			}

			return r
		}
	}

	return false
}

func (ed *Editor) HasChanges() bool {
	return !ed.history.IsEmpty()
}

func (ed *Editor) SetText(text string) {
	ed.buffer.Reset(text)
}

func (ed *Editor) GetText() string {
	return std.IterToStr(ed.buffer.Slice(0, math.MaxInt))
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

	return nil
}

func (ed *Editor) Save(filePath string) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	for text := range ed.buffer.Slice(0, math.MaxInt) {
		_, err := f.WriteString(text)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ed *Editor) insertText(text string) {
	change := ed.buffer.Insert(ed.cursor.Ln, ed.cursor.Col, text)

	ed.cursor.Set(change.End.Ln, change.End.Col, false)
	ed.history.Push()

	ed.syntax.Insert(change)
}

func (ed *Editor) deleteSelection() {
	change := ed.buffer.Delete(ed.cursor.StartLn, ed.cursor.StartCol, ed.cursor.EndLn, ed.cursor.EndCol)

	ed.cursor.Set(change.Start.Ln, change.Start.Col, false)
	ed.history.Push()

	ed.syntax.Delete(change)
}

func (ed *Editor) deleteChar() {
	change := ed.buffer.Delete(ed.cursor.Ln, ed.cursor.Col, ed.cursor.Ln, ed.cursor.Col+1)

	ed.history.Push()

	ed.syntax.Delete(change)
}

func (ed *Editor) deletePrevChar() {
	if ed.cursor.Ln == 0 && ed.cursor.Col == 0 {
		return
	}

	var change textbuf.Change

	if ed.cursor.Col == 0 {
		startLn := ed.cursor.Ln - 1
		startCol := max(0, ed.buffer.ColumnCount(startLn)-1)

		change = ed.buffer.Delete(startLn, startCol, ed.cursor.Ln, ed.cursor.Col)
	} else {
		change = ed.buffer.Delete(ed.cursor.Ln, ed.cursor.Col-1, ed.cursor.Ln, ed.cursor.Col)
	}

	ed.cursor.Set(change.Start.Ln, change.Start.Col, false)
	ed.history.Push()

	ed.syntax.Delete(change)
}
