package syntax

import (
	_ "embed"
	"fmt"
	"os"
	"time"

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
	buffer     *textbuf.TextBuf
	parser     *treeSitter.Parser
	query      *treeSitter.Query
	close      chan struct{}
	edits      chan editReq
	highlights chan highlightReq

	tree *treeSitter.Tree
	text []byte
}

func New(buffer *textbuf.TextBuf) *Syntax {
	s := Syntax{
		buffer:     buffer,
		parser:     treeSitter.NewParser(),
		close:      make(chan struct{}),
		edits:      make(chan editReq),
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

func (s *Syntax) Delete(ln0, col0, ln1, col1 int) {
	if s != nil {
		s.edits <- editReq{editKindDelete, ln0, col0, ln1, col1}
	}
}

func (s *Syntax) Insert(ln0, col0, ln1, col1 int) {
	if s != nil {
		s.edits <- editReq{editKindInsert, ln0, col0, ln1, col1}
	}
}

func (s *Syntax) Highlight(ln0, ln1 int) chan HighlightSpan {
	if s == nil {
		return nil
	}

	hls := make(chan HighlightSpan, 1024)

	s.highlights <- highlightReq{ln0, ln1, hls}

	return hls
}

func (s *Syntax) run() {
	go func() {
		for {

			select {
			case <-s.close:
				s.tree.Close()
				s.tree = nil
				return

			case req := <-s.edits:
				s.handleEditReq(req)

			case req := <-s.highlights:
				s.handleHighlightReq(req)
			}
		}
	}()
}

const maxChunkLen = 1024 * 16

func (s *Syntax) updateTree() {
	started := time.Now()

	f, err := os.OpenFile("tmp/syntax-update.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//fmt.Fprintf(f, "counter %d\n", s.counter)
	//fmt.Fprintf(f, "ranges %d\n", s.ranges)

	//s.parser.SetIncludedRanges(s.ranges)
	//maxChunkLen := int(s.ranges[0].EndByte - s.ranges[0].StartByte)

	t := s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		text := s.buffer.Chunk(i)
		if len(text) > maxChunkLen {
			text = text[0:maxChunkLen]
		}
		return []byte(text)
	}, s.tree, nil)

	s.tree.Close()
	s.tree = t

	fmt.Fprintf(f, "elapsed %v\n", time.Since(started))
}
