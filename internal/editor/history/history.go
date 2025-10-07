package history

import (
	"github.com/eu-ge-ne/toy2/internal/editor/cursor"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type History struct {
	OnChanged func()

	buffer  *textbuf.TextBuf
	cursor  *cursor.Cursor
	entries []entry
	index   int
}

type entry struct {
	ln       int
	col      int
	snapshot textbuf.Snapshot
}

func New(buffer *textbuf.TextBuf, cursor *cursor.Cursor) History {
	return History{
		buffer: buffer,
		cursor: cursor,
		entries: []entry{{
			ln:       cursor.Ln,
			col:      cursor.Col,
			snapshot: buffer.Save(),
		}},
	}
}

func (h *History) IsEmpty() bool {
	return h.index == 0
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
