package editor

import (
	"slices"
	"time"

	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/editor/data"
	"github.com/eu-ge-ne/toy2/internal/editor/history"
	"github.com/eu-ge-ne/toy2/internal/editor/render"
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

	Data   *data.Data
	render *render.Render
}

func New(multiLine bool) *Editor {
	ed := Editor{
		multiLine: multiLine,
	}

	buffer := textbuf.New()
	cursor := cursor.New(buffer)
	history := history.New(buffer, cursor)
	history.OnChanged = ed.OnChanged
	ed.Data = data.New(multiLine, buffer, cursor, history)
	ed.render = render.New(buffer, cursor)

	return &ed
}

func (ed *Editor) SetColors(t theme.Tokens) {
	ed.render.SetColors(t)
}

func (ed *Editor) Layout(a ui.Area) {
	ed.render.SetArea(a)
	ed.Data.SetPageSize(a.H)
}

func (ed *Editor) Render() {
	started := time.Now()

	if ed.OnCursor != nil {
		ed.OnCursor(ed.Data.CursorStatus())
	}

	ed.render.Render()

	if ed.OnRender != nil {
		ed.OnRender(time.Since(started))
	}
}

func (ed *Editor) SetEnabled(enabled bool) {
	ed.enabled = enabled

	ed.Data.SetEnabled(enabled)
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

	ed.Data.GoHome(false)
}

func (ed *Editor) SetWrapEnabled(enabled bool) {
	ed.render.SetWrapEnabled(enabled)
}

func (ed *Editor) ToggleWrapEnabled() {
	ed.render.ToggleWrapEnabled()

	ed.Data.GoHome(false)
}

func (ed *Editor) HandleKey(key key.Key) bool {
	if !ed.enabled {
		return false
	}

	t0 := time.Now()

	i := slices.IndexFunc(ed.Data.Handlers, func(h data.Handler) bool {
		return h.Match(key)
	})

	if i < 0 {
		return false
	}

	r := ed.Data.Handlers[i].Handle(key)

	if ed.OnKeyHandled != nil {
		ed.OnKeyHandled(time.Since(t0))
	}

	return r
}
