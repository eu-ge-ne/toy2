package editor

import (
	"math"
	"time"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/handler"
	"github.com/eu-ge-ne/toy2/internal/editor/history"
	"github.com/eu-ge-ne/toy2/internal/editor/syntax"
	"github.com/eu-ge-ne/toy2/internal/grapheme"
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

	colors colors
}

type colors struct {
	background []byte
	index      []byte
	void       []byte
	char       map[charColorEnum][]byte
}

func New(multiLine bool) *Editor {
	b := textbuf.New()
	c := cursor.New(&b)
	h := history.New(&b, &c)
	s := syntax.New(&b)

	ed := Editor{
		multiLine: multiLine,
		buffer:    &b,
		cursor:    &c,
		history:   &h,
		syntax:    &s,
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
	ed.colors = colors{
		background: t.MainBg(),
		index:      append(t.Light0Bg(), t.Dark0Fg()...),
		void:       t.Dark0Bg(),
		char: map[charColorEnum][]byte{
			charColorVisible:            append(t.MainBg(), t.Light1Fg()...),
			charColorWhitespace:         append(t.MainBg(), t.Dark0Fg()...),
			charColorEmpty:              append(t.MainBg(), t.MainFg()...),
			charColorVisibleSelected:    append(t.Light2Bg(), t.Light1Fg()...),
			charColorWhitespaceSelected: append(t.Light2Bg(), t.Dark1Fg()...),
			charColorEmptySelected:      append(t.Light2Bg(), t.Dark1Fg()...),
		},
	}
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

func (ed *Editor) ResetSyntax() {
	ed.syntax.SetLanguage()
}

func (ed *Editor) HasChanges() bool {
	return !ed.history.IsEmpty()
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

func (ed *Editor) SetText(text string) {
	ed.buffer.Reset(text)
}

func (ed *Editor) GetText() string {
	return ed.buffer.Read()
}

func (ed *Editor) LoadFromFile(filePath string) error {
	return ed.buffer.LoadFromFile(filePath)
}

func (ed *Editor) SaveToFile(filePath string) error {
	return ed.buffer.SaveToFile(filePath)
}

func (ed *Editor) deleteChar() {
	cur := ed.cursor

	ed.buffer.DeleteSlice2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
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
			ed.buffer.DeleteSlice2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			ed.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			cur.Left(false)
		} else {
			cur.Left(false)
			ed.buffer.DeleteSlice2(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
			ed.syntax.Delete(cur.Ln, cur.Col, cur.Ln, cur.Col+1)
		}
	} else {
		ed.buffer.DeleteSlice2(cur.Ln, cur.Col-1, cur.Ln, cur.Col)
		ed.syntax.Delete(cur.Ln, cur.Col-1, cur.Ln, cur.Col)
		cur.Left(false)
	}

	ed.history.Push()
}

func (ed *Editor) deleteSelection() {
	cur := ed.cursor

	ed.buffer.DeleteSlice2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	ed.syntax.Delete(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
	ed.cursor.Set(cur.StartLn, cur.StartCol, false)

	ed.history.Push()
}

func (ed *Editor) insertText(text string) {
	cur := ed.cursor

	if cur.Selecting {
		ed.buffer.DeleteSlice2(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		ed.syntax.Delete(cur.StartLn, cur.StartCol, cur.EndLn, cur.EndCol)
		cur.Set(cur.StartLn, cur.StartCol, false)
	}

	ed.buffer.Insert2(cur.Ln, cur.Col, text)

	dLn, dCol := grapheme.Graphemes.MeasureText(text)
	cur.Forward(dLn, dCol)

	ed.syntax.Insert(cur.Ln, cur.Col, text)

	ed.history.Push()
}
