package syntax

import (
	"fmt"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type editReq struct {
	kind  editKind
	start textbuf.Pos
	end   textbuf.Pos
}

type editKind int

const (
	editKindDelete editKind = iota
	editKindInsert
)

func (s *Syntax) Delete(start, end textbuf.Pos) {
	if s != nil {
		s.edits <- editReq{editKindDelete, start, end}
	}
}

func (s *Syntax) Insert(start, end textbuf.Pos) {
	if s != nil {
		s.edits <- editReq{editKindInsert, start, end}
	}
}

func (s *Syntax) handleEdit(req editReq) {
	if s.tree == nil {
		return
	}

	switch req.kind {
	case editKindDelete:
		s.edit.StartByte = uint(req.start.Idx)
		s.edit.OldEndByte = uint(req.end.Idx)
		s.edit.NewEndByte = s.edit.StartByte

		s.edit.StartPosition.Row = uint(req.start.Ln)
		s.edit.StartPosition.Column = uint(req.start.Col)
		s.edit.OldEndPosition.Row = uint(req.end.Ln)
		s.edit.OldEndPosition.Column = uint(req.end.Col)
		s.edit.NewEndPosition = s.edit.StartPosition
	case editKindInsert:
		s.edit.StartByte = uint(req.start.Idx)
		s.edit.OldEndByte = s.edit.StartByte
		s.edit.NewEndByte = uint(req.end.Idx)

		s.edit.StartPosition.Row = uint(req.start.Ln)
		s.edit.StartPosition.Column = uint(req.start.Col)
		s.edit.OldEndPosition = s.edit.StartPosition
		s.edit.NewEndPosition.Row = uint(req.end.Ln)
		s.edit.NewEndPosition.Column = uint(req.end.Col)
	}

	fmt.Fprintf(s.log, "edit: %v\n", req)
	fmt.Fprintf(s.log, "edit: %+v\n", s.edit)

	s.tree.Edit(&s.edit)

	s.updateTree()
}
