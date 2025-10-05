package editor

import (
	"math"
	"time"

	"github.com/eu-ge-ne/toy2/internal/cursor"
	"github.com/eu-ge-ne/toy2/internal/history"
	"github.com/eu-ge-ne/toy2/internal/segbuf"
	"github.com/eu-ge-ne/toy2/internal/syntax"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Editor struct {
	multiLine bool
	area      ui.Area
	Enabled   bool
	clipboard string

	IndexEnabled      bool
	WhitespaceEnabled bool
	WrapEnabled       bool

	indexWidth int
	textWidth  int
	cursorY    int
	cursorX    int
	scrollLn   int
	scrollCol  int

	Buffer       *segbuf.SegBuf
	Cursor       *cursor.Cursor
	history      *history.History
	syntax       *syntax.Syntax
	handlers     []Handler
	OnKeyHandled func(time.Duration)
	OnRender     func(time.Duration)
	OnCursor     func(int, int, int)
	OnChanged    func()

	colors colors
}

type colors struct {
	background []byte
	index      []byte
	void       []byte
	char       map[charColorEnum][]byte
}

func New(multiLine bool) *Editor {
	b := segbuf.New()
	c := cursor.New(&b)
	h := history.New(&b, &c)
	s := syntax.New(&b)

	ed := Editor{
		multiLine: multiLine,
		Buffer:    &b,
		Cursor:    &c,
		history:   &h,
		syntax:    &s,
	}

	ed.history.OnChanged = ed.OnChanged

	ed.handlers = append(ed.handlers,
		&TextHandler{editor: &ed},
		&BackspaceHandler{editor: &ed},
		&BottomHandler{editor: &ed},
		&CopyHandler{editor: &ed},
		&CutHandler{editor: &ed},
		&DeleteHandler{editor: &ed},
		&DownHandler{editor: &ed},
		&EndHandler{editor: &ed},
		&EnterHandler{editor: &ed},
		&HomeHandler{editor: &ed},
		&LeftHandler{editor: &ed},
		&PageDownHandler{editor: &ed},
		&PageUpHandler{editor: &ed},
		&PasteHandler{editor: &ed},
		&RedoHandler{editor: &ed},
		&RightHandler{editor: &ed},
		&SelectAllHandler{editor: &ed},
		&TopHandler{editor: &ed},
		&UndoHandler{editor: &ed},
		&UpHandler{editor: &ed},
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

func (ed *Editor) Reset(resetCursor bool) {
	if resetCursor {
		if ed.multiLine {
			ed.Cursor.Set(0, 0, false)
		} else {
			ed.Cursor.Set(math.MaxInt, math.MaxInt, false)
		}
	}

	ed.history.Reset()
	ed.syntax.Reset()
}

func (ed *Editor) Copy() bool {
	c := ed.Cursor

	if c.Selecting {
		ed.clipboard = ed.Buffer.Read(c.FromLn, c.FromCol, c.ToLn, c.ToCol+1)
		c.Set(c.Ln, c.Col, false)
	} else {
		ed.clipboard = ed.Buffer.Read(c.Ln, c.Col, c.Ln, c.Col+1)
	}

	vt.CopyToClipboard(vt.Sync, ed.clipboard)

	return false
}

func (ed *Editor) Cut() bool {
	c := ed.Cursor

	if c.Selecting {
		ed.clipboard = ed.Buffer.Read(c.FromLn, c.FromCol, c.ToLn, c.ToCol+1)
		ed.deleteSelection()
	} else {
		ed.clipboard = ed.Buffer.Read(c.Ln, c.Col, c.Ln, c.Col+1)
		ed.deleteChar()
	}

	vt.CopyToClipboard(vt.Sync, ed.clipboard)

	return true
}

func (ed *Editor) Paste() bool {
	if len(ed.clipboard) == 0 {
		return false
	}

	ed.insert(ed.clipboard)

	return true
}

func (ed *Editor) Undo() bool {
	return ed.history.Undo()
}

func (ed *Editor) Redo() bool {
	return ed.history.Redo()
}

func (ed *Editor) HasChanges() bool {
	return !ed.history.IsEmpty()
}
