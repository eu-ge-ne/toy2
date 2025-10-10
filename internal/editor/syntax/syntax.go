package syntax

import (
	"time"

	"github.com/eu-ge-ne/toy2/internal/textbuf"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

type Syntax struct {
	buffer  *textbuf.TextBuf
	parser  *treeSitter.Parser
	tree    *treeSitter.Tree
	edits   chan edit
	isDirty bool
}

func New(buffer *textbuf.TextBuf) *Syntax {
	s := Syntax{
		buffer: buffer,
		parser: treeSitter.NewParser(),
		edits:  make(chan edit, 100),
	}

	err := s.parser.SetLanguage(treeSitter.NewLanguage(treeSitterTs.LanguageTypescript()))
	if err != nil {
		panic(err)
	}

	go s.run()

	return &s
}

func (s *Syntax) Reset() {
	if s != nil {
		s.edits <- edit{editReset, 0, 0, 0, 0}
	}
}

func (s *Syntax) Delete(startLn, startCol, endLn, endCol int) {
	if s != nil {
		s.edits <- edit{editDelete, startLn, startCol, endLn, endCol}
	}
}

func (s *Syntax) Insert(startLn, startCol, endLn, endCol int) {
	if s != nil {
		s.edits <- edit{editInsert, startLn, startCol, endLn, endCol}
	}
}

func (s *Syntax) run() {
	for {
		timeout := time.After(100 * time.Millisecond)

		select {
		case <-timeout:
			if s.isDirty {
				s.parseTree()
			}
		case pos := <-s.edits:
			switch pos.edit {
			case editReset:
				s.tree = nil
				s.parseTree()
			case editDelete:
				start, oldEnd, ok := s.buffer.Index2(pos.startLn, pos.startCol, pos.endLn, pos.endCol)
				if !ok {
					panic("in Syntax.Delete")
				}
				s.editTree(start, oldEnd, start+1, pos.startLn, pos.startCol, pos.endLn, pos.endCol, pos.startLn, pos.startCol+1)
				s.isDirty = true
			case editInsert:
				start, newEnd, ok := s.buffer.Index2(pos.startLn, pos.startCol, pos.endLn, pos.endCol)
				if !ok {
					panic("in Syntax.Insert")
				}
				s.editTree(start, start+1, newEnd, pos.startLn, pos.startCol, pos.startLn, pos.startCol+1, pos.endLn, pos.endCol)
				s.isDirty = true
			}
		}
	}
}

func (s *Syntax) editTree(start, oldEnd, newEnd, startLn, startCol, oldEndLn, oldEndCol, newEndLn, newEndCol int) {
	s.tree.Edit(&treeSitter.InputEdit{
		StartByte:      uint(start),
		OldEndByte:     uint(oldEnd),
		NewEndByte:     uint(newEnd),
		StartPosition:  treeSitter.NewPoint(uint(startLn), uint(startCol)),
		OldEndPosition: treeSitter.NewPoint(uint(oldEndLn), uint(oldEndCol)),
		NewEndPosition: treeSitter.NewPoint(uint(newEndLn), uint(newEndCol)),
	})
}

func (s *Syntax) parseTree() {
	s.tree = s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		return []byte(s.buffer.Chunk(i))
	}, s.tree, nil)

	s.isDirty = false
}
