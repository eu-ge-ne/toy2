package syntax

import (
	"fmt"

	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func (s *Syntax) Delete(change textbuf.Change) {
	if s == nil || s.tree == nil {
		return
	}

	var e treeSitter.InputEdit

	e.StartByte = uint(change.Start.Idx)
	e.OldEndByte = uint(change.End.Idx)
	e.NewEndByte = e.StartByte

	e.StartPosition.Row = uint(change.Start.Ln)
	e.StartPosition.Column = uint(change.Start.ColIdx)
	e.OldEndPosition.Row = uint(change.End.Ln)
	e.OldEndPosition.Column = uint(change.End.ColIdx)
	e.NewEndPosition = e.StartPosition

	s.tree.Edit(&e)

	fmt.Fprintf(s.log, "delete: change %+v\n", change)
	fmt.Fprintf(s.log, "delete: e %+v\n", e)

	s.changed = true
}

func (s *Syntax) Insert(change textbuf.Change) {
	if s == nil || s.tree == nil {
		return
	}

	var e treeSitter.InputEdit

	e.StartByte = uint(change.Start.Idx)
	e.OldEndByte = e.StartByte
	e.NewEndByte = uint(change.End.Idx)

	e.StartPosition.Row = uint(change.Start.Ln)
	e.StartPosition.Column = uint(change.Start.ColIdx)
	e.OldEndPosition = e.StartPosition
	e.NewEndPosition.Row = uint(change.End.Ln)
	e.NewEndPosition.Column = uint(change.End.ColIdx)

	s.tree.Edit(&e)

	fmt.Fprintf(s.log, "insert: change %+v\n", change)
	fmt.Fprintf(s.log, "insert: e %+v\n", e)

	s.changed = true
}
