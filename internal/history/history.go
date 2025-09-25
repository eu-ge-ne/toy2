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
	OnChanged func(int)
}

type entry struct {
	ln       int
	col      int
	snapshot *textbuf.Snapshot
}

func New(buffer *textbuf.TextBuf, cursor *cursor.Cursor) *History {
	h := &History{
		buffer: buffer,
		cursor: cursor,
	}

	h.Reset()

	return h
}

func (h *History) Reset() {
	snapshot := h.buffer.Save()

	h.entries = []entry{{
		ln:       h.cursor.Ln,
		col:      h.cursor.Col,
		snapshot: snapshot,
	}}

	h.index = 0

	if h.OnChanged != nil {
		h.OnChanged(h.index)
	}
}

func (h *History) Push() {
	snapshot := h.buffer.Save()

	h.entries = append([]entry(nil), h.entries[:h.index+1]...)
	h.entries = append(h.entries, entry{
		ln:       h.cursor.Ln,
		col:      h.cursor.Col,
		snapshot: snapshot,
	})

	h.index += 1

	if h.OnChanged != nil {
		h.OnChanged(h.index)
	}
}

func (h *History) Undo() bool {
	if h.index > 0 {
		h.index -= 1
		h.restore()

		if h.OnChanged != nil {
			h.OnChanged(h.index)
		}

		return true
	}

	return false
}

func (h *History) Redo() bool {
	if h.index < (len(h.entries) - 1) {
		h.index += 1
		h.restore()

		if h.OnChanged != nil {
			h.OnChanged(h.index)
		}

		return true
	}

	return false
}

func (h *History) restore() {
	entry := h.entries[h.index]

	h.buffer.Restore(entry.snapshot)
	h.cursor.Set(entry.ln, entry.col, false)
}
