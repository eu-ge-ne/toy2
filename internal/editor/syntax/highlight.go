package syntax

import (
	"fmt"
	"time"

	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type highlightReq struct {
	start textbuf.Pos
	end   textbuf.Pos
	spans chan Span
}

type Span struct {
	StartIdx int
	EndIdx   int
	Name     string
}

func (s *Syntax) Highlight(start, end textbuf.Pos) chan Span {
	if s == nil {
		return nil
	}

	spans := make(chan Span, 1024)

	s.highlights <- highlightReq{start, end, spans}

	return spans
}

func (s *Syntax) handleHighlight(req highlightReq) {
	started := time.Now()

	if s.tree == nil {
		s.updateTree()
	}

	if s.buffer.Count() > len(s.text) {
		s.text = make([]byte, s.buffer.Count())
	}
	copy(
		s.text[req.start.Idx:req.end.Idx],
		std.IterToStr(s.buffer.Read(req.start.Idx, req.end.Idx)),
	)

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()

	qc.SetByteRange(uint(req.start.Idx), uint(req.end.Idx))
	capts := qc.Captures(s.query, s.tree.RootNode(), s.text)

	var span Span

	match, captIdx := capts.Next()
	if match != nil {
		capt := match.Captures[captIdx]
		span = Span{
			StartIdx: int(capt.Node.StartByte()),
			EndIdx:   int(capt.Node.EndByte()),
			Name:     s.query.CaptureNames()[capt.Index],
		}
	}

	for ; match != nil; match, captIdx = capts.Next() {
		capt := match.Captures[captIdx]
		name := s.query.CaptureNames()[capt.Index]

		fmt.Fprintf(s.log,
			"highlight: %v:%v %s (%s)\n",
			capt.Node.StartPosition(),
			capt.Node.EndPosition(),
			capt.Node.Utf8Text(s.text),
			name,
			//match.PatternIndex,
			//capt.Index,
		)

		startIdx := int(capt.Node.StartByte())
		endIdx := int(capt.Node.EndByte())

		if span.StartIdx != startIdx || span.EndIdx != endIdx {
			req.spans <- span
			span = Span{StartIdx: startIdx, EndIdx: endIdx}
		}

		span.Name = name
	}

	req.spans <- span
	close(req.spans)

	fmt.Fprintf(s.log, "highlight: elapsed %v\n", time.Since(started))
}
