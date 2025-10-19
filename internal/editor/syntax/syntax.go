package syntax

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	_ "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
	treeSitterTs "github.com/tree-sitter/tree-sitter-typescript/bindings/go"

	"github.com/eu-ge-ne/toy2/internal/std"
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
	edit       treeSitter.InputEdit
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

func (s *Syntax) Highlight(ln0, ln1 int) chan Span {
	if s == nil {
		return nil
	}

	spans := make(chan Span, 1024)

	s.highlights <- highlightReq{ln0, ln1, spans}

	return spans
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
				s.handleEdit(req)

			case req := <-s.highlights:
				s.handleHighlight(req)
			}
		}
	}()
}

func (s *Syntax) handleHighlight(req highlightReq) {
	f, err := os.OpenFile("tmp/syntax-highlight.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	started := time.Now()

	if s.tree == nil {
		s.updateTree()
	}

	ln0 := max(0, req.ln0)
	ln1 := min(s.buffer.LineCount(), req.ln1)
	startByte, _ := s.buffer.LnIndex(ln0)
	endByte, _ := s.buffer.LnIndex(ln1)
	startPoint := treeSitter.NewPoint(uint(ln0), 0)
	endPoint := treeSitter.NewPoint(uint(ln1), 0)

	if s.buffer.Count() > len(s.text) {
		s.text = make([]byte, s.buffer.Count())
	}
	copy(s.text[startByte:endByte], std.IterToStr(s.buffer.Read(startByte, endByte)))

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()

	qc.SetPointRange(startPoint, endPoint)
	capts := qc.Captures(s.query, s.tree.RootNode(), s.text)

	var span Span

	match, captIdx := capts.Next()
	if match != nil {
		span = Span{
			Start: int(match.Captures[captIdx].Node.StartByte()),
			End:   int(match.Captures[captIdx].Node.EndByte()),
		}
	}

	for ; match != nil; match, captIdx = capts.Next() {
		capt := match.Captures[captIdx]
		name := s.query.CaptureNames()[capt.Index]

		fmt.Fprintf(f,
			"%v:%v %s (%s, %d, %d)\n",
			capt.Node.StartPosition(),
			capt.Node.EndPosition(),
			capt.Node.Utf8Text(s.text),
			name,
			match.PatternIndex,
			capt.Index,
		)

		start := int(capt.Node.StartByte())
		end := int(capt.Node.EndByte())

		if span.Start != start || span.End != end {
			req.spans <- span
			span = Span{
				Start: start,
				End:   end,
			}
		}

		span.Name = name
	}

	req.spans <- span
	close(req.spans)

	fmt.Fprintf(f, "elapsed %v\n", time.Since(started))
}

func (s *Syntax) handleEdit(req editReq) {
	f, err := os.OpenFile("tmp/syntax-edit.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if s.tree == nil {
		return
	}

	i0, ok := s.buffer.Index(req.ln0, req.col0)
	if !ok {
		panic(fmt.Sprintf("in Syntax.handleEditReq: %v", req))
	}

	i1, ok := s.buffer.Index(req.ln1, req.col1)
	if !ok {
		panic(fmt.Sprintf("in Syntax.handleEditReq: %v", req))
	}

	col0, ok := s.buffer.ColIndex(req.ln0, req.col0)
	if !ok {
		panic(fmt.Sprintf("in Syntax.handleEditReq: %v", req))
	}

	col1, ok := s.buffer.ColIndex(req.ln1, req.col1)
	if !ok {
		panic(fmt.Sprintf("in Syntax.handleEditReq: %v", req))
	}

	switch req.kind {
	case editKindDelete:
		s.edit.StartByte = uint(i0)
		s.edit.OldEndByte = uint(i1)
		s.edit.NewEndByte = s.edit.StartByte

		s.edit.StartPosition.Row = uint(req.ln0)
		s.edit.StartPosition.Column = uint(col0)
		s.edit.OldEndPosition.Row = uint(req.ln1)
		s.edit.OldEndPosition.Column = uint(col1)
		s.edit.NewEndPosition = s.edit.StartPosition
	case editKindInsert:
		s.edit.StartByte = uint(i0)
		s.edit.OldEndByte = s.edit.StartByte
		s.edit.NewEndByte = uint(i1)

		s.edit.StartPosition.Row = uint(req.ln0)
		s.edit.StartPosition.Column = uint(col0)
		s.edit.OldEndPosition = s.edit.StartPosition
		s.edit.NewEndPosition.Row = uint(req.ln1)
		s.edit.NewEndPosition.Column = uint(col1)
	}

	fmt.Fprintf(f, "%v\n", req)
	fmt.Fprintf(f, "%+v\n", s.edit)

	s.tree.Edit(&s.edit)

	s.updateTree()
}

const maxChunkLen = 1024 * 64

func (s *Syntax) updateTree() {
	started := time.Now()

	f, err := os.OpenFile("tmp/syntax-tree.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	/*
		s.parser.SetIncludedRanges([]treeSitter.Range{{
			StartByte:  uint(startByte),
			EndByte:    uint(endByte),
			StartPoint: startPoint,
			EndPoint:   endPoint,
		}})
	*/

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
