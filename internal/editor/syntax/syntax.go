package syntax

import (
	_ "embed"
	"fmt"
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
	buffer *textbuf.TextBuf

	parser  *treeSitter.Parser
	ranges  []treeSitter.Range
	tree    *treeSitter.Tree
	close   chan struct{}
	edits   chan edit
	isDirty bool

	queryHighlights *treeSitter.Query

	parseCounter int
	hlCounter    int
}

func New(buffer *textbuf.TextBuf) *Syntax {
	s := Syntax{
		buffer: buffer,

		parser: treeSitter.NewParser(),
		ranges: []treeSitter.Range{{}},
		close:  make(chan struct{}),
		edits:  make(chan edit, 100),
	}

	//Log(s.parser)

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

func (s *Syntax) Scroll(startLn, endLn int) {
	if s == nil {
		return
	}

	i0, ok := s.buffer.LnIndex(startLn)
	if !ok {
		return
		panic("in Syntax.SetArea")
	}

	i1, ok := s.buffer.LnIndex(endLn)
	if !ok {
		return
		panic("in Syntax.SetArea")
	}

	s.ranges[0].StartByte = uint(i0)
	s.ranges[0].EndByte = uint(i1)
	s.ranges[0].StartPoint.Row = uint(startLn)
	s.ranges[0].EndPoint.Row = uint(endLn)

	s.parseTree()
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

func (s *Syntax) Highlight() {
	if s == nil || s.tree == nil {
		return
	}

	started := time.Now()

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()

	f, err := os.OpenFile("tmp/highlight.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	qc.SetPointRange(s.ranges[0].StartPoint, s.ranges[0].EndPoint)
	text := []byte(std.IterToStr(s.buffer.Read2(int(s.ranges[0].StartPoint.Row), 0, int(s.ranges[0].EndPoint.Row), 0)))
	matches := qc.Matches(s.queryHighlights, s.tree.RootNode(), text)

	for match := matches.Next(); match != nil; match = matches.Next() {
		for _, capture := range match.Captures {
			fmt.Fprintf(f,
				"Match %d, Capture %d: %s |%s| %v, %v\n",
				match.PatternIndex,
				capture.Index,
				s.queryHighlights.CaptureNames()[capture.Index],
				capture.Node.Utf8Text(text),
				capture.Node.StartPosition(),
				capture.Node.EndPosition(),
			)
		}
	}

	fmt.Fprintf(f, "%d: Elapsed %v\n", s.hlCounter, time.Since(started))
	s.hlCounter += 1
}

func (s *Syntax) run() {
	go func() {
		for {
			timeout := time.After(100 * time.Millisecond)

			select {
			case <-s.close:
				s.tree.Close()
				s.tree = nil
				return

			case p := <-s.edits:
				ed, ok := p.index(s.buffer)
				if !ok {
					panic("in Syntax.edits")
				}
				s.tree.Edit(&ed)
				s.isDirty = true

			case <-timeout:
				if s.isDirty {
					s.parseTree()
					s.isDirty = false
				}
			}
		}
	}()
}

func (s *Syntax) parseTree() {
	started := time.Now()

	f, err := os.OpenFile("tmp/parse.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "%d: Ranges %v\n", s.parseCounter, s.ranges)
	s.parser.SetIncludedRanges(s.ranges)

	newTree := s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		text := s.buffer.Chunk(i)

		if len(text) < 1024 {
			return []byte(text)
		}

		return []byte(text[0:1024])
	}, s.tree, nil)

	s.tree.Close()
	s.tree = newTree

	fmt.Fprintf(f, "%d: Elapsed %v\n", s.parseCounter, time.Since(started))
	s.parseCounter += 1
}
