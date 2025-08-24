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
	enabled bool

	indexEnabled      bool
	whitespaceEnabled bool
	wrapEnabled       bool

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
	cursor       *cursor.Cursor
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
		cursor:    cursor,
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

	editor.indexEnabled = true
	editor.whitespaceEnabled = true
	editor.wrapEnabled = true

	return &editor
}

func (ed *Editor) Reset(resetCursor bool) {
	if resetCursor {
		if ed.multiLine {
			ed.cursor.Set(0, 0, false)
		} else {
			ed.cursor.Set(math.MaxInt, math.MaxInt, false)
		}
	}
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
