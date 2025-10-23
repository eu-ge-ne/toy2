package syntax

import (
	"fmt"
	"time"

	treeSitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/eu-ge-ne/toy2/internal/std"
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

	start, _ := s.buffer.Pos(req.startLn, 0)
	end := s.buffer.EndPos(req.endLn, 0)

	s.parser.SetIncludedRanges([]treeSitter.Range{{
		StartByte:  uint(start.Idx),
		EndByte:    uint(end.Idx),
		StartPoint: treeSitter.NewPoint(uint(start.Ln), uint(start.ColIdx)),
		EndPoint:   treeSitter.NewPoint(uint(end.Ln), uint(end.ColIdx)),
	}})

	oldTree := s.tree

	s.tree = s.parser.ParseWithOptions(func(i int, p treeSitter.Point) []byte {
		text := s.buffer.Chunk(i)
		if len(text) > maxChunkLen {
			text = text[0:maxChunkLen]
		}
		fmt.Fprintf(s.log, "highlight: reading chunk %d, %+v, %d\n", i, p, len(text))
		return []byte(text)
	}, oldTree, nil)

	fmt.Fprintf(s.log, "highlight: parsed %v\n", time.Since(started))

	if s.buffer.Count() > len(s.text) {
		s.text = make([]byte, s.buffer.Count())
	}
	copy(s.text[start.Idx:end.Idx], std.IterToStr(s.buffer.Slice(start.Idx, end.Idx)))

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()

	qc.SetByteRange(uint(start.Idx), uint(end.Idx))
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
