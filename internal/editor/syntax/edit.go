package syntax

import (
	treeSitter "github.com/tree-sitter/go-tree-sitter"

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

func (e *edit) index(buf *textbuf.TextBuf) (r treeSitter.InputEdit, ok bool) {
	i0, ok := buf.Index(e.ln0, e.col0)
	if !ok {
		return
	}

	i1, ok := buf.Index(e.ln1, e.col1)
	if !ok {
		return
	}

	col0i, ok := buf.ColIndex(e.ln0, e.col0)
	if !ok {
		return
	}

	col1i, ok := buf.ColIndex(e.ln1, e.col1)
	if !ok {
		return
	}

	switch e.kind {
	case editKindDelete:
		r.StartByte = uint(i0)
		r.OldEndByte = uint(i1)
		r.NewEndByte = uint(i0 + 1)
		r.StartPosition = treeSitter.NewPoint(uint(e.ln0), uint(col0i))
		r.OldEndPosition = treeSitter.NewPoint(uint(e.ln1), uint(col1i))
		r.NewEndPosition = treeSitter.NewPoint(uint(e.ln0), uint(col0i+1))
	case editKindInsert:
		r.StartByte = uint(i0)
		r.OldEndByte = uint(i0 + 1)
		r.NewEndByte = uint(i1)
		r.StartPosition = treeSitter.NewPoint(uint(e.ln0), uint(col0i))
		r.OldEndPosition = treeSitter.NewPoint(uint(e.ln0), uint(col0i+1))
		r.NewEndPosition = treeSitter.NewPoint(uint(e.ln1), uint(col1i))
	}

	ok = true

	return
}
