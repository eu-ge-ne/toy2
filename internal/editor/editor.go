package editor

import (
	"math"
	"time"

	"github.com/eu-ge-ne/toy2/internal/cursor"
	"github.com/eu-ge-ne/toy2/internal/history"
	"github.com/eu-ge-ne/toy2/internal/segbuf"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
	"github.com/eu-ge-ne/toy2/internal/vt"
)

type Editor struct {
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

	multiLine    bool
	Buffer       *segbuf.SegBuf
	Cursor       *cursor.Cursor
	History      *history.History
	handlers     []Handler
	OnKeyHandled func(time.Duration)
	OnRender     func(time.Duration)
	OnCursor     func(int, int, int)

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
	c := cursor.New(b)
	h := history.New(b, c)

	editor := Editor{
		multiLine: multiLine,
		Buffer:    b,
		Cursor:    c,
		History:   h,
	}

	editor.handlers = append(editor.handlers,
		&TextHandler{editor: &editor},
		&BackspaceHandler{editor: &editor},
		&BottomHandler{editor: &editor},
		&CopyHandler{editor: &editor},
		&CutHandler{editor: &editor},
		&DeleteHandler{editor: &editor},
		&DownHandler{editor: &editor},
		&EndHandler{editor: &editor},
		&EnterHandler{editor: &editor},
		&HomeHandler{editor: &editor},
		&LeftHandler{editor: &editor},
		&PageDownHandler{editor: &editor},
		&PageUpHandler{editor: &editor},
		&PasteHandler{editor: &editor},
		&RedoHandler{editor: &editor},
		&RightHandler{editor: &editor},
		&SelectAllHandler{editor: &editor},
		&TopHandler{editor: &editor},
		&UndoHandler{editor: &editor},
		&UpHandler{editor: &editor},
	)

	return &editor
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

func (ed *Editor) Area() ui.Area {
	return ed.area
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

	ed.History.Reset()
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
