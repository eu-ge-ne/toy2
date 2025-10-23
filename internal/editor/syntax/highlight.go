package syntax

import (
	"fmt"
	"time"

	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

type highlightReq struct {
	startLn int
	endLn   int
	spans   chan Span
}

type Span struct {
	StartIdx int
	EndIdx   int
	Name     string
}

const maxChunkLen = 1024 * 64

func (s *Syntax) Highlight(startLn, endLn int) <-chan Span {
	if s == nil {
		return nil
	}

	spans := make(chan Span, 1024)

	s.highlight <- highlightReq{startLn, endLn, spans}

	return spans
}

func (s *Syntax) handleHighlight(req highlightReq) {
	started := time.Now()

	fmt.Fprintln(s.log, "highlight: started")

	startPos, _ := s.buffer.Pos(req.startLn, 0)
	endPos := s.buffer.EndPos(req.endLn, 0)

	s.parse(startPos, endPos)

	fmt.Fprintf(s.log, "highlight: parsed %v\n", time.Since(started))

	s.prepareText(startPos, endPos)

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()

	qc.SetByteRange(uint(startPos.Idx), uint(endPos.Idx))

	var span Span

	capts := qc.Captures(s.query, s.tree.RootNode(), s.text)

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

		/*
			fmt.Fprintf(s.log,
				"highlight: %v:%v %s (%s)\n",
				capt.Node.StartPosition(),
				capt.Node.EndPosition(),
				capt.Node.Utf8Text(s.text),
				name,
				//match.PatternIndex,
				//capt.Index,
			)
		*/

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

func (s *Syntax) parse(startPos, endPos textbuf.Pos) {
	s.parser.SetIncludedRanges([]treeSitter.Range{{
		StartByte:  uint(startPos.Idx),
		EndByte:    uint(endPos.Idx),
		StartPoint: treeSitter.NewPoint(uint(startPos.Ln), uint(startPos.ColIdx)),
		EndPoint:   treeSitter.NewPoint(uint(endPos.Ln), uint(endPos.ColIdx)),
	}})

	oldTree := s.tree

	s.tree = s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		text := s.buffer.Chunk(i)

		if len(text) > maxChunkLen {
			text = text[0:maxChunkLen]
		}

		fmt.Fprintf(s.log, "parse: reading chunk %d, %+v, %d\n", i, p, len(text))

		return []byte(text)
	}, oldTree, nil)
}

func (s *Syntax) prepareText(startPos, endPos textbuf.Pos) {
	if s.buffer.Count() > len(s.text) {
		s.text = make([]byte, s.buffer.Count())
	}

	copy(s.text[startPos.Idx:endPos.Idx], std.IterToStr(s.buffer.Slice(startPos.Idx, endPos.Idx)))
}
