package syntax

import (
	_ "embed"
	"fmt"
	"os"
	"slices"
	"time"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
	_ "github.com/tree-sitter/tree-sitter-javascript/bindings/go"

	"github.com/eu-ge-ne/toy2/internal/std"
)

type highlightReq struct {
	ln0   int
	ln1   int
	spans chan HighlightSpan
}

type HighlightSpan struct {
	Start int
	End   int
	Color CharFgColor
}

func (s *Syntax) handleHighlightReq(req highlightReq) {
	started := time.Now()

	f, err := os.OpenFile("tmp/syntax-highlight.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if s.tree == nil {
		s.updateTree()
	}

	ln0 := max(0, req.ln0)
	ln1 := min(s.buffer.LineCount(), req.ln1)
	start, _ := s.buffer.LnIndex(ln0)
	end, _ := s.buffer.LnIndex(ln1)
	startPoint := treeSitter.NewPoint(uint(ln0), 0)
	endPoint := treeSitter.NewPoint(uint(ln1), 0)

	if s.buffer.Count() > len(s.text) {
		s.text = make([]byte, s.buffer.Count())
	}
	copy(s.text[start:end], std.IterToStr(s.buffer.Read(start, end)))

	qc := treeSitter.NewQueryCursor()
	defer qc.Close()

	qc.SetPointRange(startPoint, endPoint)
	capts := qc.Captures(s.query, s.tree.RootNode(), s.text)

	var (
		span         HighlightSpan
		spanCaptures []int
	)

	match, captIdx := capts.Next()
	if match != nil {
		span = HighlightSpan{
			Start: int(match.Captures[captIdx].Node.StartByte()),
			End:   int(match.Captures[captIdx].Node.EndByte()),
			Color: CharFgColorUndefined,
		}
		spanCaptures = make([]int, 0, 5)
	}

	for ; match != nil; match, captIdx = capts.Next() {
		capt := match.Captures[captIdx]

		/*
			fmt.Fprintf(f,
				"%v:%v %s (%s, %d, %d)\n",
				capt.Node.StartPosition(),
				capt.Node.EndPosition(),
				capt.Node.Utf8Text(s.text),
				s.query.CaptureNames()[capt.Index],
				match.PatternIndex,
				capt.Index,
			)
		*/

		start := int(capt.Node.StartByte())
		end := int(capt.Node.EndByte())

		if span.Start != start || span.End != end {
			req.spans <- span
			span = HighlightSpan{
				Start: start,
				End:   end,
				Color: CharFgColorUndefined,
			}
			spanCaptures = make([]int, 0, 5)
		}

		spanCaptures = append(spanCaptures, int(capt.Index))

		if slices.Contains(spanCaptures, 0 /*variable*/) {
			span.Color = CharFgColorVariable
		} else if slices.Contains(spanCaptures, 18 /*keyword*/) {
			span.Color = CharFgColorKeyword
		} else if slices.Contains(spanCaptures, 9 /*comment*/) {
			span.Color = CharFgColorComment
		} else {
			span.Color = CharFgColorUndefined
		}
	}

	req.spans <- span
	close(req.spans)

	fmt.Fprintf(f, "elapsed %v\n", time.Since(started))
}
