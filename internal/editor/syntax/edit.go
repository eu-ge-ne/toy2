package syntax

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type edit struct {
	kind editKind
	ln0  int
	col0 int
	ln1  int
	col1 int
}

type editKind int

const (
	editKindDelete editKind = iota
	editKindInsert
)

func (e *edit) index(buf *textbuf.TextBuf) (i0, i1, col0i, col1i int, ok bool) {
	i0, ok = buf.Index(e.ln0, e.col0)
	if !ok {
		return
	}

	i1, ok = buf.Index(e.ln1, e.col1)
	if !ok {
		return
	}

	col0i, ok = buf.ColIndex(e.ln0, e.col0)
	if !ok {
		return
	}

	col1i, ok = buf.ColIndex(e.ln1, e.col1)

	return
}
