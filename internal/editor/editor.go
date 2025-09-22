package editor

import (
	"math"
	"time"

	"github.com/eu-ge-ne/toy2/internal/cursor"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
	"github.com/eu-ge-ne/toy2/internal/theme"
	"github.com/eu-ge-ne/toy2/internal/ui"
)

type Editor struct {
	area    ui.Area
	Enabled bool

	IndexEnabled      bool
	WhitespaceEnabled bool
	WrapEnabled       bool

	indexWidth int
	textWidth  int
	wrapWidth  int
	cursorY    int
	cursorX    int
	measureY   int
	measureX   int
	scrollLn   int
	scrollCol  int

	multiLine    bool
	Buffer       *textbuf.TextBuf
	Cursor       *cursor.Cursor
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
	buffer := textbuf.New("")
	cursor := cursor.New(buffer)

	editor := Editor{
		multiLine: multiLine,
		Buffer:    buffer,
		Cursor:    cursor,
	}

	editor.handlers = append(editor.handlers,
		&TextHandler{editor: &editor},
		&BackspaceHandler{editor: &editor},
		&DeleteHandler{editor: &editor},
		&LeftHandler{editor: &editor},
		&RightHandler{editor: &editor},
		&UpHandler{editor: &editor},
		&DownHandler{editor: &editor},
		&PageUpHandler{editor: &editor},
		&PageDownHandler{editor: &editor},
		&EnterHandler{editor: &editor},
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
}
