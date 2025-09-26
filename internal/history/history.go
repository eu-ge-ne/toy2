package history

import (
	"github.com/eu-ge-ne/toy2/internal/cursor"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type History struct {
	buffer    *textbuf.TextBuf
	cursor    *cursor.Cursor
	entries   []entry
	index     int
	OnChanged func()
}

type entry struct {
	ln   int
	col  int
	text *textbuf.Snapshot
}

func New(buffer *textbuf.TextBuf, cursor *cursor.Cursor) *History {
	h := &History{
		buffer: buffer,
		cursor: cursor,
	}

	h.Reset()

	return h
}

func (h *History) IsEmpty() bool {
	return h.index == 0
}

func (h *History) Reset() {
	h.entries = []entry{{
		ln:   h.cursor.Ln,
		col:  h.cursor.Col,
		text: h.buffer.Save(),
	}}

	h.index = 0

	if h.OnChanged != nil {
		h.OnChanged()
	}
}

func (h *History) Push() {
	h.entries = append([]entry(nil), h.entries[:h.index+1]...)
	h.entries = append(h.entries, entry{
		ln:   h.cursor.Ln,
		col:  h.cursor.Col,
		text: h.buffer.Save(),
	})

	h.index += 1

	if h.OnChanged != nil {
		h.OnChanged()
	}
}

func (h *History) Undo() bool {
	if h.index <= 0 {
		return false
	}

	h.index -= 1
	h.restore()

	if h.OnChanged != nil {
		h.OnChanged()
	}

	return true
}

func (h *History) Redo() bool {
	if h.index >= len(h.entries)-1 {
		return false
	}

	h.index += 1
	h.restore()

	if h.OnChanged != nil {
		h.OnChanged()
	}

	return true
}

func (h *History) restore() {
	entry := h.entries[h.index]

	h.buffer.Restore(entry.text)
	h.cursor.Set(entry.ln, entry.col, false)
}
