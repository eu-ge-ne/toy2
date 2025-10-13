package editor

import (
	"io"
	"os"
	"time"
	"unicode/utf8"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/history"
	"github.com/eu-ge-ne/toy2/internal/editor/render"
	"github.com/eu-ge-ne/toy2/internal/editor/syntax"
	"github.com/eu-ge-ne/toy2/internal/grapheme"
	"github.com/eu-ge-ne/toy2/internal/key"
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

	ed.syntax = syntax.New(ed.buffer)
}

func (ed *Editor) SetColors(t theme.Tokens) {
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
	ed.syntax.Reset()
}

func (ed *Editor) GetText() string {
	return ed.buffer.All()
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

func (ed *Editor) deleteChar() {
	cur := ed.cursor

	ed.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
	ed.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)

	ed.history.Push()
}

func (ed *Editor) deletePrevChar() {
	cur := ed.cursor

	if cur.Ln > 0 && cur.Col == 0 {
		l := 0
		for range ed.buffer.IterLine(cur.Ln, false) {
			l += 1
			if l == 2 {
				break
			}
		}

		if l == 1 {
			ed.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			ed.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			cur.Left(false)
		} else {
			cur.Left(false)
			ed.buffer.Delete2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			ed.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
		}
	} else {
		ed.buffer.Delete2(cur.Ln, cur.Col-1, cur.Ln, cur.Col)
		ed.syntax.Delete(cur.Ln, cur.Col-1, cur.Ln, cur.Col)
		cur.Left(false)
	}

	ed.history.Push()
}

func (ed *Editor) deleteSelection() {
	cur := ed.cursor

	ed.buffer.Delete2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	ed.syntax.Delete(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	ed.cursor.Set(cur.StartLn, cur.StartCol, false)

	ed.history.Push()
}

func (ed *Editor) insertText(text string) {
	cur := ed.cursor

	if cur.Selecting {
		ed.buffer.Delete2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		ed.syntax.Delete(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		cur.Set(cur.StartLn, cur.StartCol, false)
	}

	ed.buffer.Insert2(cur.Ln, cur.Col, text)

	startLn := cur.Ln
	startCol := cur.Col

	dLn, dCol := grapheme.Graphemes.MeasureText(text)
	cur.Forward(dLn, dCol)

	endLn := cur.Ln
	endCol := cur.Col

	ed.syntax.Insert(startLn, startCol, endLn, endCol)

	ed.history.Push()
}
