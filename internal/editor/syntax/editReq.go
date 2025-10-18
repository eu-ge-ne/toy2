package syntax

import (
	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type editReq struct {
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

func (op editReq) inputEdit(buffer *textbuf.TextBuf) (r treeSitter.InputEdit, ok bool) {
	i0, ok := buffer.Index(op.ln0, op.col0)
	if !ok {
		return
	}

	i1, ok := buffer.Index(op.ln1, op.col1)
	if !ok {
		return
	}

	col0i, ok := buffer.ColIndex(op.ln0, op.col0)
	if !ok {
		return
	}

	col1i, ok := buffer.ColIndex(op.ln1, op.col1)
	if !ok {
		return
	}

	switch op.kind {
	case editKindDelete:
		r.StartByte = uint(i0)
		r.OldEndByte = uint(i1)
		r.NewEndByte = uint(i0 + 1)
		r.StartPosition = treeSitter.NewPoint(uint(op.ln0), uint(col0i))
		r.OldEndPosition = treeSitter.NewPoint(uint(op.ln1), uint(col1i))
		r.NewEndPosition = treeSitter.NewPoint(uint(op.ln0), uint(col0i+1))
	case editKindInsert:
		r.StartByte = uint(i0)
		r.OldEndByte = uint(i0 + 1)
		r.NewEndByte = uint(i1)
		r.StartPosition = treeSitter.NewPoint(uint(op.ln0), uint(col0i))
		r.OldEndPosition = treeSitter.NewPoint(uint(op.ln0), uint(col0i+1))
		r.NewEndPosition = treeSitter.NewPoint(uint(op.ln1), uint(col1i))
	}

	ok = true

	return
}
