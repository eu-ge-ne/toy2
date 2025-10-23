package syntax

import (
	_ "embed"
	"os"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	_ "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

//go:embed js/highlights.scm
var scmJsHighlights string

//go:embed ts/highlights.scm
var scmTsHighlights string

type Syntax struct {
	buffer *textbuf.TextBuf
	parser *treeSitter.Parser
	query  *treeSitter.Query
	tree   *treeSitter.Tree

	close      chan struct{}
	highlights chan highlightReq

	text []byte
	log  *os.File
}

func New(buffer *textbuf.TextBuf) *Syntax {
	s := Syntax{
		buffer:     buffer,
		parser:     treeSitter.NewParser(),
		close:      make(chan struct{}),
		highlights: make(chan highlightReq),
	}

	//Log(s.parser)

	lang := treeSitter.NewLanguage(treeSitterTs.LanguageTypescript())
	err := s.parser.SetLanguage(lang)
	if err != nil {
		panic(err)
	}

	query, err0 := treeSitter.NewQuery(lang, scmJsHighlights+scmTsHighlights)
	if err0 != nil {
		panic(err0)
	}
	s.query = query

	s.run()

	return &s
}

func (s *Syntax) Close() {
	if s != nil {
		s.close <- struct{}{}
	}
}

func (s *Syntax) Restart() {
	if s != nil {
		s.Close()
		s.run()
	}
}

func (s *Syntax) run() {
	f, err := os.OpenFile("tmp/syntax.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	s.log = f

	go func() {
		for {
			select {
			case <-s.close:
				s.handleClose()
				return

			case req := <-s.highlights:
				s.handleHighlight(req)
			}
		}
	}()
}

func (s *Syntax) handleClose() {
	if s.log != nil {
		s.log.Close()
		s.log = nil
	}

	s.tree.Close()
	s.tree = nil
}
