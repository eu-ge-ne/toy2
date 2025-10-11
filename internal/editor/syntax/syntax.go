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
	close   chan struct{}
	reset   chan struct{}
	edits   chan edit
	isDirty bool
}

type edit struct {
	kind     editKind
	startLn  int
	startCol int
	endLn    int
	endCol   int
}

type editKind int

const (
	editDelete editKind = iota
	editInsert
)

func New(buffer *textbuf.TextBuf) *Syntax {
	s := Syntax{
		buffer: buffer,
		parser: treeSitter.NewParser(),
		close:  make(chan struct{}),
		reset:  make(chan struct{}),
		edits:  make(chan edit, 100),
	}

	err := s.parser.SetLanguage(treeSitter.NewLanguage(treeSitterTs.LanguageTypescript()))
	if err != nil {
		panic(err)
	}

	go s.run()

	s.Reset()

	return &s
}

func (s *Syntax) Close() {
	if s != nil {
		s.close <- struct{}{}
	}
}

func (s *Syntax) Reset() {
	if s != nil {
		s.reset <- struct{}{}
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
		case <-s.close:
			return
		case <-timeout:
			s.parseTree()
		case <-s.reset:
			s.resetTree()
		case p := <-s.edits:
			s.editTree(p)
		}
	}
}

func (s *Syntax) resetTree() {
	s.tree = nil
	s.isDirty = true
	s.parseTree()
}

func (s *Syntax) parseTree() {
	if !s.isDirty {
		return
	}

	s.tree = s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		return []byte(s.buffer.Chunk(i))
	}, s.tree, nil)

	s.isDirty = false
}

func (s *Syntax) editTree(p edit) {
	var start, oldEnd, newEnd, startLn, startCol, oldEndLn, oldEndCol, newEndLn, newEndCol int

	a, b, ok := s.buffer.Index2(p.startLn, p.startCol, p.endLn, p.endCol)
	if !ok {
		panic("in Syntax.editTree")
	}

	switch p.kind {
	case editDelete:
		start = a
		oldEnd = b
		newEnd = start + 1

		startLn = p.startLn
		startCol = p.startCol

		oldEndLn = p.endLn
		oldEndCol = p.endCol

		newEndLn = p.startLn
		newEndCol = p.startCol + 1

	case editInsert:
		start = a
		oldEnd = start + 1
		newEnd = b

		startLn = p.startLn
		startCol = p.startCol

		oldEndLn = p.startLn
		oldEndCol = p.startCol + 1

		newEndLn = p.endLn
		newEndCol = p.endCol
	}

	s.tree.Edit(&treeSitter.InputEdit{
		StartByte:      uint(start),
		OldEndByte:     uint(oldEnd),
		NewEndByte:     uint(newEnd),
		StartPosition:  treeSitter.NewPoint(uint(startLn), uint(startCol)),
		OldEndPosition: treeSitter.NewPoint(uint(oldEndLn), uint(oldEndCol)),
		NewEndPosition: treeSitter.NewPoint(uint(newEndLn), uint(newEndCol)),
	})

	s.isDirty = true
}
