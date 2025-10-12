package syntax

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/eu-ge-ne/toy2/internal/textbuf"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

//go:embed highlights.scm
var scmHighlights string

type Syntax struct {
	buffer          *textbuf.TextBuf
	parser          *treeSitter.Parser
	queryHighlights *treeSitter.Query
	tree            *treeSitter.Tree
	close           chan struct{}
	reset           chan struct{}
	edits           chan edit
	isDirty         bool
}

func New(buffer *textbuf.TextBuf) *Syntax {
	s := Syntax{
		buffer: buffer,
		parser: treeSitter.NewParser(),
		close:  make(chan struct{}),
		reset:  make(chan struct{}),
		edits:  make(chan edit, 100),
	}

	s.log()

	lang := treeSitter.NewLanguage(treeSitterTs.LanguageTypescript())

	err := s.parser.SetLanguage(lang)
	if err != nil {
		panic(err)
	}

	queryHighlights, qerr := treeSitter.NewQuery(lang, scmHighlights)
	if qerr != nil {
		panic(qerr)
	}

	s.queryHighlights = queryHighlights

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

func (s *Syntax) editTree(e edit) {
	p, ok := e.index(s.buffer)
	if !ok {
		panic("in Syntax.editTree")
	}

	s.tree.Edit(&p)

	s.isDirty = true
}
