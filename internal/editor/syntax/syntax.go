package syntax

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	treeSitterJs "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

//go:embed js/highlights.scm
var scmJsHighlights string

//go:embed ts/highlights.scm
var scmTsHighlights string

type Syntax struct {
	buffer          *textbuf.TextBuf
	parser          *treeSitter.Parser
	queryHighlights *treeSitter.Query
	tree            *treeSitter.Tree
	close           chan struct{}
	reset           chan struct{}
	edits           chan edit
	isDirty         bool
	hlCounter       int
}

func New(buffer *textbuf.TextBuf) *Syntax {
	s := Syntax{
		buffer: buffer,
		close:  make(chan struct{}),
		reset:  make(chan struct{}),
		edits:  make(chan edit, 100),
	}

	s.parser = treeSitter.NewParser()
	log(s.parser)

	treeSitter.NewLanguage(treeSitterJs.Language())
	langTs := treeSitter.NewLanguage(treeSitterTs.LanguageTypescript())

	err := s.parser.SetLanguage(langTs)
	if err != nil {
		panic(err)
	}

	queryHighlights, err0 := treeSitter.NewQuery(langTs, scmJsHighlights+scmTsHighlights)
	if err0 != nil {
		panic(err0)
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

func log(parser *treeSitter.Parser) {
	f, err := os.OpenFile("tmp/syntax.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	i := 0

	parser.SetLogger(func(t treeSitter.LogType, msg string) {
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

		case <-s.reset:
			s.resetTree()

		case p := <-s.edits:
			s.editTree(p)

		case <-timeout:
			s.parseTree()
		}
	}
}

func (s *Syntax) resetTree() {
	s.tree = nil
	s.isDirty = true
	s.parseTree()
}

func (s *Syntax) editTree(e edit) {
	p, ok := e.index(s.buffer)
	if !ok {
		panic("in Syntax.editTree")
	}

	s.tree.Edit(&p)

	s.isDirty = true
}

func (s *Syntax) parseTree() {
	if !s.isDirty {
		return
	}

	newTree := s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		return s.buffer.Chunk(i)
	}, s.tree, nil)

	s.tree.Close()

	s.tree = newTree
	s.isDirty = false

	s.highlight()
}

func (s *Syntax) highlight() {
	started := time.Now()

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()

	f, err := os.OpenFile("tmp/highlight.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	text := std.IterToBytes(s.buffer.Read(0, math.MaxInt))
	matches := qc.Matches(s.queryHighlights, s.tree.RootNode(), text)

	for match := matches.Next(); match != nil; match = matches.Next() {
		for _, capture := range match.Captures {
			fmt.Fprintf(f,
				"Match %d, Capture %d (%s): %s\n",
				match.PatternIndex,
				capture.Index,
				s.queryHighlights.CaptureNames()[capture.Index],
				capture.Node.Utf8Text(text),
			)
		}
	}

	fmt.Fprintf(f, "%d: Elapsed %v\n", s.hlCounter, time.Since(started))
	s.hlCounter += 1
}
