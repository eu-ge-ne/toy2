package history

import (
	"github.com/eu-ge-ne/toy2/internal/cursor"
	"github.com/eu-ge-ne/toy2/internal/segbuf"
)

type History struct {
	buffer    *segbuf.SegBuf
	cursor    *cursor.Cursor
	entries   []entry
	index     int
	OnChanged func()
}

type entry struct {
	ln       int
	col      int
	snapshot segbuf.Snapshot
}

func New(buffer *segbuf.SegBuf, cursor *cursor.Cursor) History {
	h := History{
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
		ln:       h.cursor.Ln,
		col:      h.cursor.Col,
		snapshot: h.buffer.Save(),
	}}

	h.index = 0

	if h.OnChanged != nil {
		h.OnChanged()
	}
}

func (h *History) Push() {
	h.entries = append([]entry(nil), h.entries[:h.index+1]...)
	h.entries = append(h.entries, entry{
		ln:       h.cursor.Ln,
		col:      h.cursor.Col,
		snapshot: h.buffer.Save(),
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

	h.buffer.Restore(entry.snapshot)
	h.cursor.Set(entry.ln, entry.col, false)
}
