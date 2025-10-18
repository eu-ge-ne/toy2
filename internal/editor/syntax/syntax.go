package syntax

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	_ "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

//go:embed js/highlights.scm
var scmJsHighlights string

//go:embed ts/highlights.scm
var scmTsHighlights string

type Syntax struct {
	Highlighter *Highlighter

	buffer *textbuf.TextBuf
	parser *treeSitter.Parser
	close  chan struct{}
	ops    chan op

	ranges  []treeSitter.Range
	isDirty bool
	tree    *treeSitter.Tree
	query   *treeSitter.Query

	text    []byte
	counter int
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
		close:  make(chan struct{}),
		ops:    make(chan op, 100),

		ranges: []treeSitter.Range{{}},
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
				s.tree.Close()
				s.tree = nil
				return

			case op := <-s.ops:
				s.handleOp(op)

			case <-timeout:
				if s.isDirty {
					s.update()
				}
			}
		}
	}()
}

func (s *Syntax) handleOp(op op) {
	if op.kind == opKindScroll {
		h := op.ln1 - op.ln0
		ln0 := max(0, op.ln0-h)
		ln1 := min(s.buffer.LineCount(), op.ln1+h)

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
		panic(fmt.Sprintf("in Syntax.handleOp: %v", op))
	}

	s.tree.Edit(&ed)
	s.isDirty = true
}

func (s *Syntax) update() {
	s.updateTree()

	s.updateHighlight()

	s.counter += 1
	s.isDirty = false
}

const maxChunkLen = 1024 * 16

func (s *Syntax) updateTree() {
	started := time.Now()

	f, err := os.OpenFile("tmp/syntax-update.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "counter %d\n", s.counter)
	fmt.Fprintf(f, "ranges %d\n", s.ranges)

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

func (s *Syntax) updateHighlight() {
	started := time.Now()

	f, err := os.OpenFile("tmp/syntax-highlight.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "counter %d\n", s.counter)
	fmt.Fprintf(f, "ranges %d\n", s.ranges)

	if s.buffer.Count() > len(s.text) {
		s.text = make([]byte, s.buffer.Count())
	}
	start := int(s.ranges[0].StartByte)
	end := int(s.ranges[0].EndByte)
	copy(s.text[start:end], std.IterToStr(s.buffer.Read(start, end)))

	highlighter := newHighlighter()

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()

	qc.SetPointRange(s.ranges[0].StartPoint, s.ranges[0].EndPoint)
	capts := qc.Captures(s.query, s.tree.RootNode(), s.text)

	for match, captIdx := capts.Next(); match != nil; match, captIdx = capts.Next() {
		capt := match.Captures[captIdx]

		highlighter.AddCapture(capt)

		fmt.Fprintf(f,
			"hl: %v:%v %s (%s, %d, %d)\n",
			capt.Node.StartPosition(),
			capt.Node.EndPosition(),
			capt.Node.Utf8Text(s.text),
			s.query.CaptureNames()[capt.Index],
			match.PatternIndex,
			capt.Index,
		)
	}

	s.Highlighter = highlighter

	fmt.Fprintf(f, "hl: elapsed %v\n", time.Since(started))
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
