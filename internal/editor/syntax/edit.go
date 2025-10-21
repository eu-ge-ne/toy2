package syntax

import (
	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func (s *Syntax) Delete(start, end textbuf.Pos) {
	if s == nil || s.tree == nil {
		return
	}

	var e treeSitter.InputEdit

	e.StartByte = uint(start.Idx)
	e.OldEndByte = uint(end.Idx)
	e.NewEndByte = e.StartByte

	e.StartPosition.Row = uint(start.Ln)
	e.StartPosition.Column = uint(start.Col)
	e.OldEndPosition.Row = uint(end.Ln)
	e.OldEndPosition.Column = uint(end.Col)
	e.NewEndPosition = e.StartPosition

	s.tree.Edit(&e)
	s.dirty = true
}

func (s *Syntax) Insert(start, end textbuf.Pos) {
	if s == nil || s.tree == nil {
		return
	}

	var e treeSitter.InputEdit

	e.StartByte = uint(start.Idx)
	e.OldEndByte = e.StartByte
	e.NewEndByte = uint(end.Idx)

	e.StartPosition.Row = uint(start.Ln)
	e.StartPosition.Column = uint(start.Col)
	e.OldEndPosition = e.StartPosition
	e.NewEndPosition.Row = uint(end.Ln)
	e.NewEndPosition.Column = uint(end.Col)

	s.tree.Edit(&e)
	s.dirty = true
}
