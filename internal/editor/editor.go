package editor

import (
	"io"
	"math"
	"os"
	"unicode/utf8"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/frame"
	"github.com/eu-ge-ne/toy2/internal/editor/history"
	"github.com/eu-ge-ne/toy2/internal/editor/syntax"
	"github.com/eu-ge-ne/toy2/internal/grammar"
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Editor struct {
	OnCursor  func(int, int, int)
	OnChanged func()

	multiLine bool
	buffer    *textbuf.TextBuf
	cursor    *cursor.Cursor
	history   *history.History
	syntax    *syntax.Syntax
	frame     *frame.Frame

	area      ui.Area
	enabled   bool
	clipboard string
	handlers  []Handler
}

func New(multiLine bool) *Editor {
	buffer := textbuf.New()

	ed := &Editor{
		multiLine: multiLine,
		buffer:    buffer,
	}

	ed.cursor = cursor.New(buffer)
	ed.cursor.OnChanged = ed.onCursorChanged

	ed.history = history.New(buffer, ed.cursor)
	ed.history.OnChanged = ed.OnChanged

	ed.syntax = syntax.New()
	ed.frame = frame.New(buffer, ed.cursor, ed.syntax)

	ed.handlers = []Handler{
		&Insert{ed},
		&Backspace{ed},
		&Bottom{ed},
		&Copy{ed},
		&Cut{ed},
		&Delete{ed},
		&Down{ed},
		&End{ed},
		&Enter{ed},
		&Home{ed},
		&Left{ed},
		&PageDown{ed},
		&PageUp{ed},
		&Paste{ed},
		&Redo{ed},
		&Right{ed},
		&SelectAll{ed},
		&Top{ed},
		&Undo{ed},
		&Up{ed},
	}

	return ed
}

func (ed *Editor) SetGrammar(grm grammar.Grammar) {
	ed.syntax.SetGrammar(grm)
}

func (ed *Editor) SetColors(t theme.Theme) {
	ed.frame.SetColors(t)
}

func (ed *Editor) SetArea(a ui.Area) {
	ed.area = a
	ed.frame.SetArea(a)
}

func (ed *Editor) SetEnabled(e bool) {
	ed.enabled = e
}

func (ed *Editor) SetIndexEnabled(e bool) {
	ed.frame.SetIndexEnabled(e)
}

func (ed *Editor) SetWrapEnabled(e bool) {
	ed.frame.SetWrapEnabled(e)
}

func (ed *Editor) ToggleWrapEnabled() {
	ed.frame.ToggleWrapEnabled()
}

func (ed *Editor) SetWhitespaceEnabled(e bool) {
	ed.frame.SetWhitespaceEnabled(e)
}

func (ed *Editor) ToggleWhitespaceEnabled() {
	ed.frame.ToggleWhitespaceEnabled()
}

func (ed *Editor) HasChanges() bool {
	return !ed.history.IsEmpty()
}

func (ed *Editor) GetText() string {
	return std.IterToStr(ed.buffer.Slice(0, math.MaxInt))
}

func (ed *Editor) SetText(text string) {
	ed.buffer.Reset(text)
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

func (ed *Editor) Render() {
	ed.frame.Render(ed.enabled)
}

func (ed *Editor) HandleKey(key key.Key) bool {
	if !ed.enabled {
		return false
	}

	for _, h := range ed.handlers {
		if h.Match(key) {
			return h.Run(key)
		}
	}

	return false
}

func (ed *Editor) Backspace() bool {
	if ed.cursor.Selecting {
		return ed.deleteSelection()
	}

	if ed.cursor.Ln == 0 && ed.cursor.Col == 0 {
		return false
	}

	var (
		change textbuf.Change
		ok     bool
	)

	if ed.cursor.Col == 0 {
		startLn := ed.cursor.Ln - 1
		startCol := max(0, ed.buffer.ColumnCount(startLn)-1)

		change, ok = ed.buffer.Delete(startLn, startCol, ed.cursor.Ln, ed.cursor.Col)
	} else {
		change, ok = ed.buffer.Delete(ed.cursor.Ln, ed.cursor.Col-1, ed.cursor.Ln, ed.cursor.Col)
	}

	if !ok {
		return false
	}

	ed.cursor.Set(change.Start.Ln, change.Start.Col, false)
	ed.history.Push()

	ed.syntax.Delete(change)

	return true
}

func (ed *Editor) Bottom(sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Bottom(sel)
}

func (ed *Editor) Copy() bool {
	cur := ed.cursor

	if cur.Selecting {
		ed.clipboard = std.IterToStr(ed.buffer.Read(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol))
		cur.Set(cur.Ln, cur.Col, false)
	} else {
		ed.clipboard = std.IterToStr(ed.buffer.Read(cur.Ln, cur.Col, cur.Ln, cur.Col+1))
	}

	vt.CopyToClipboard(vt.Sync, ed.clipboard)

	return false
}

func (ed *Editor) Cut() bool {
	cur := ed.cursor

	defer func() {
		vt.CopyToClipboard(vt.Sync, ed.clipboard)
	}()

	if cur.Selecting {
		ed.clipboard = std.IterToStr(ed.buffer.Read(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol))
		return ed.deleteSelection()
	} else {
		ed.clipboard = std.IterToStr(ed.buffer.Read(cur.Ln, cur.Col, cur.Ln, cur.Col+1))
		return ed.deleteChar()
	}
}

func (ed *Editor) Delete() bool {
	if ed.cursor.Selecting {
		return ed.deleteSelection()
	} else {
		return ed.deleteChar()
	}
}

func (ed *Editor) Down(n int, sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Down(n, sel)
}

func (ed *Editor) End(sel bool) bool {
	return ed.cursor.End(sel)
}

func (ed *Editor) Enter() bool {
	if !ed.multiLine {
		return false
	}

	return ed.insertText("\n")
}

func (ed *Editor) Home(sel bool) bool {
	return ed.cursor.Home(sel)
}

func (ed *Editor) Insert(text string) bool {
	return ed.insertText(text)
}

func (ed *Editor) Left(sel bool) bool {
	return ed.cursor.Left(sel)
}

func (ed *Editor) Paste() bool {
	if len(ed.clipboard) == 0 {
		return false
	}

	return ed.Insert(ed.clipboard)
}

func (ed *Editor) Redo() bool {
	if !ed.enabled {
		return false
	}

	return ed.history.Redo()
}

func (ed *Editor) Right(sel bool) bool {
	return ed.cursor.Right(sel)
}

func (ed *Editor) SelectAll() bool {
	if !ed.enabled {
		return false
	}

	ed.cursor.Set(0, 0, false)
	ed.cursor.Set(math.MaxInt, math.MaxInt, true)

	return true
}

func (ed *Editor) Top(sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Top(sel)
}

func (ed *Editor) Undo() bool {
	if !ed.enabled {
		return false
	}

	return ed.history.Undo()
}

func (ed *Editor) Up(n int, sel bool) bool {
	if !ed.multiLine {
		return false
	}

	return ed.cursor.Up(n, sel)
}

func (ed *Editor) insertText(text string) bool {
	if ed.cursor.Selecting {
		return ed.deleteSelection()
	}

	change := ed.buffer.Insert(ed.cursor.Ln, ed.cursor.Col, text)

	ed.cursor.Set(change.End.Ln, change.End.Col, false)
	ed.history.Push()

	ed.syntax.Insert(change)

	return true
}

func (ed *Editor) deleteSelection() bool {
	change, ok := ed.buffer.Delete(ed.cursor.StartLn, ed.cursor.StartCol, ed.cursor.EndLn, ed.cursor.EndCol)

	if !ok {
		return false
	}

	ed.cursor.Set(change.Start.Ln, change.Start.Col, false)
	ed.history.Push()

	ed.syntax.Delete(change)

	return true
}

func (ed *Editor) deleteChar() bool {
	change, ok := ed.buffer.Delete(ed.cursor.Ln, ed.cursor.Col, ed.cursor.Ln, ed.cursor.Col+1)

	if !ok {
		return false
	}

	ed.history.Push()

	ed.syntax.Delete(change)

	return true
}

func (ed *Editor) onCursorChanged() {
	if ed.OnCursor != nil {
		ed.OnCursor(ed.cursor.Ln, ed.cursor.Col, ed.buffer.LineCount())
	}
}
