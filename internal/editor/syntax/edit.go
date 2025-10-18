package syntax

import (
	"fmt"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
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

func (s *Syntax) handleEditReq(req editReq) {
	i0, ok := s.buffer.Index(req.ln0, req.col0)
	if !ok {
		panic(fmt.Sprintf("in Syntax.handleOp: %v", req))
	}

	i1, ok := s.buffer.Index(req.ln1, req.col1)
	if !ok {
		panic(fmt.Sprintf("in Syntax.handleOp: %v", req))
	}

	col0i, ok := s.buffer.ColIndex(req.ln0, req.col0)
	if !ok {
		panic(fmt.Sprintf("in Syntax.handleOp: %v", req))
	}

	col1i, ok := s.buffer.ColIndex(req.ln1, req.col1)
	if !ok {
		panic(fmt.Sprintf("in Syntax.handleOp: %v", req))
	}

	var ed treeSitter.InputEdit

	switch req.kind {
	case editKindDelete:
		ed.StartByte = uint(i0)
		ed.OldEndByte = uint(i1)
		ed.NewEndByte = uint(i0 + 1)
		ed.StartPosition = treeSitter.NewPoint(uint(req.ln0), uint(col0i))
		ed.OldEndPosition = treeSitter.NewPoint(uint(req.ln1), uint(col1i))
		ed.NewEndPosition = treeSitter.NewPoint(uint(req.ln0), uint(col0i+1))
	case editKindInsert:
		ed.StartByte = uint(i0)
		ed.OldEndByte = uint(i0 + 1)
		ed.NewEndByte = uint(i1)
		ed.StartPosition = treeSitter.NewPoint(uint(req.ln0), uint(col0i))
		ed.OldEndPosition = treeSitter.NewPoint(uint(req.ln0), uint(col0i+1))
		ed.NewEndPosition = treeSitter.NewPoint(uint(req.ln1), uint(col1i))
	}

	s.tree.Edit(&ed)

	s.updateTree()
}
