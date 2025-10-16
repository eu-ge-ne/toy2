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
	ops     chan op
	isDirty bool

	queryHighlights *treeSitter.Query

	counter int
	hlText  []byte
}

type op struct {
	kind opKind
	ln0  int
	col0 int
	ln1  int
	col1 int
}

type opKind int

const (
	opKindScroll opKind = iota
	opKindDelete
	opKindInsert
)

func New(buffer *textbuf.TextBuf) *Syntax {
	s := Syntax{
		buffer: buffer,

		parser: treeSitter.NewParser(),
		ranges: []treeSitter.Range{{}},
		close:  make(chan struct{}),
		ops:    make(chan op, 100),
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

func (s *Syntax) Scroll(ln0, ln1 int) {
	if s != nil {
		s.ops <- op{opKindScroll, ln0, 0, ln1, 0}
	}
}

func (s *Syntax) Delete(ln0, col0, ln1, col1 int) {
	if s != nil {
		s.ops <- op{opKindDelete, ln0, col0, ln1, col1}
	}
}

func (s *Syntax) Insert(ln0, col0, ln1, col1 int) {
	if s != nil {
		s.ops <- op{opKindInsert, ln0, col0, ln1, col1}
	}
}

func (s *Syntax) run() {
	go func() {
		for {
			timeout := time.After(10 * time.Millisecond)

			select {
			case <-s.close:
				s.handleClose()
				return

			case op := <-s.ops:
				s.handleOp(op)

			case <-timeout:
				s.handleTimeout()
			}
		}
	}()
}

func (s *Syntax) handleClose() {
	s.tree.Close()
	s.tree = nil
}

func (s *Syntax) handleOp(op op) {
	if op.kind == opKindScroll {
		ln0 := min(s.buffer.LineCount(), op.ln0)
		ln1 := min(s.buffer.LineCount(), op.ln1)

		i0, _ := s.buffer.LnIndex(ln0)
		i1, _ := s.buffer.LnIndex(ln1)

		s.ranges[0].StartByte = uint(i0)
		s.ranges[0].EndByte = uint(i1)
		s.ranges[0].StartPoint.Row = uint(ln0)
		s.ranges[0].EndPoint.Row = uint(ln1)

		s.update()
		return
	}

	ed, ok := s.inputEdit(op)
	if !ok {
		panic(fmt.Sprintf("in Syntax.handleEdit: %v", op))
	}

	s.tree.Edit(&ed)
	s.isDirty = true
}

func (s *Syntax) handleTimeout() {
	if s.isDirty {
		s.update()
	}
}

func (s *Syntax) update() {
	started := time.Now()

	f, err := os.OpenFile("tmp/syntax.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "update: counter=%d\n", s.counter)
	fmt.Fprintf(f, "update: ranges=%d\n", s.ranges)

	s.updateTree()
	s.updateHighlights(f)

	fmt.Fprintf(f, "Elapsed %v\n", time.Since(started))

	s.counter += 1
	s.isDirty = false
}

func (s *Syntax) updateTree() {
	s.parser.SetIncludedRanges(s.ranges)

	maxChunkLen := int(s.ranges[0].EndByte - s.ranges[0].StartByte)

	t := s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		text := s.buffer.Chunk(i)
		if len(text) > maxChunkLen {
			text = text[0:maxChunkLen]
		}
		return []byte(text)
	}, s.tree, nil)

	s.tree.Close()
	s.tree = t
}

func (s *Syntax) updateHighlights(f *os.File) {
	rng := s.ranges[0]

	if len(s.hlText) != s.buffer.Count() {
		s.hlText = make([]byte, s.buffer.Count())
	}
	chunk := std.IterToStr(
		s.buffer.Read(int(rng.StartByte), int(rng.EndByte)),
	)
	copy(s.hlText[rng.StartByte:rng.EndByte], chunk)

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()
	qc.SetPointRange(rng.StartPoint, rng.EndPoint)

	matches := qc.Matches(s.queryHighlights, s.tree.RootNode(), s.hlText)

	for match := matches.Next(); match != nil; match = matches.Next() {
		for _, capture := range match.Captures {
			/*
				fmt.Fprintf(f,
					"highlight: Match %d, Capture %d: %s |%s| %v, %v\n",
					match.PatternIndex,
					capture.Index,
					s.queryHighlights.CaptureNames()[capture.Index],
					capture.Node.Utf8Text(s.hlText),
					capture.Node.StartPosition(),
					capture.Node.EndPosition(),
				)
			*/
			fmt.Fprintf(f,
				"highlight: [%v:%v] %s (%s)\n",
				capture.Node.StartPosition(),
				capture.Node.EndPosition(),
				capture.Node.Utf8Text(s.hlText),
				s.queryHighlights.CaptureNames()[capture.Index],
			)
		}
	}
}

func (s *Syntax) inputEdit(op op) (r treeSitter.InputEdit, ok bool) {
	i0, ok := s.buffer.Index(op.ln0, op.col0)
	if !ok {
		return
	}

	i1, ok := s.buffer.Index(op.ln1, op.col1)
	if !ok {
		return
	}

	col0i, ok := s.buffer.ColIndex(op.ln0, op.col0)
	if !ok {
		return
	}

	col1i, ok := s.buffer.ColIndex(op.ln1, op.col1)
	if !ok {
		return
	}

	switch op.kind {
	case opKindDelete:
		r.StartByte = uint(i0)
		r.OldEndByte = uint(i1)
		r.NewEndByte = uint(i0 + 1)
		r.StartPosition = treeSitter.NewPoint(uint(op.ln0), uint(col0i))
		r.OldEndPosition = treeSitter.NewPoint(uint(op.ln1), uint(col1i))
		r.NewEndPosition = treeSitter.NewPoint(uint(op.ln0), uint(col0i+1))
	case opKindInsert:
		r.StartByte = uint(i0)
		r.OldEndByte = uint(i0 + 1)
		r.NewEndByte = uint(i1)
		r.StartPosition = treeSitter.NewPoint(uint(op.ln0), uint(col0i))
		r.OldEndPosition = treeSitter.NewPoint(uint(op.ln0), uint(col0i+1))
		r.NewEndPosition = treeSitter.NewPoint(uint(op.ln1), uint(col1i))
	}

	ok = true

	return
}
