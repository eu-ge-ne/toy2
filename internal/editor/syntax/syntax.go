package syntax

import (
	"fmt"
	"os"
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

func New(buffer *textbuf.TextBuf) *Syntax {
	s := Syntax{
		buffer: buffer,
		parser: treeSitter.NewParser(),
		close:  make(chan struct{}),
		reset:  make(chan struct{}),
		edits:  make(chan edit, 100),
	}

	s.log()

	err := s.parser.SetLanguage(treeSitter.NewLanguage(treeSitterTs.LanguageTypescript()))
	if err != nil {
		panic(err)
	}

	s.resetTree()

	go s.run()

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
		s.edits <- edit{editKindDelete, startLn, startCol, endLn, endCol}
	}
}

func (s *Syntax) Insert(startLn, startCol, endLn, endCol int) {
	if s != nil {
		s.edits <- edit{editKindInsert, startLn, startCol, endLn, endCol}
	}
}

func (s *Syntax) log() {
	f, err := os.OpenFile("tmp/syntax.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	i := 0

	s.parser.SetLogger(func(t treeSitter.LogType, msg string) {
		var tp string

		switch t {
		case treeSitter.LogTypeParse:
			tp = "Parse"
		case treeSitter.LogTypeLex:
			tp = "Lex"
		}

		fmt.Fprintf(f, "%d: %s: %s\n", i, tp, msg)

		i += 1
	})
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
	i0, ok := s.buffer.Index(p.ln0, p.col0)
	if !ok {
		panic("in Syntax.editTree")
	}

	i1, ok := s.buffer.Index(p.ln1, p.col1)
	if !ok {
		panic("in Syntax.editTree")
	}

	col0i, ok := s.buffer.ColIndex(p.ln0, p.col0)
	if !ok {
		panic("in Syntax.editTree")
	}

	col1i, ok := s.buffer.ColIndex(p.ln1, p.col1)
	if !ok {
		panic("in Syntax.editTree")
	}

	var start, oldEnd, newEnd, startLn, startCol, oldEndLn, oldEndCol, newEndLn, newEndCol int

	switch p.kind {
	case editKindDelete:
		start = i0
		oldEnd = i1
		newEnd = start + 1

		startLn = p.ln0
		startCol = col0i

		oldEndLn = p.ln1
		oldEndCol = col1i

		newEndLn = p.ln0
		newEndCol = col0i + 1

	case editKindInsert:
		start = i0
		oldEnd = start + 1
		newEnd = i1

		startLn = p.ln0
		startCol = col0i

		oldEndLn = p.ln0
		oldEndCol = col0i + 1

		newEndLn = p.ln1
		newEndCol = col1i
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
