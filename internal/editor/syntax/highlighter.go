package syntax

import (
	"slices"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
)

type Highlighter struct {
	spans   []HighlightSpan
	spanIdx int
	idx     int
}

type HighlightSpan struct {
	Start int
	End   int
	Color CharFgColor

	captures []int
}

func newHighlighter() *Highlighter {
	return &Highlighter{
		spans: make([]HighlightSpan, 0, 1024),
	}
}

func (h *Highlighter) AddCapture(capt treeSitter.QueryCapture) {
	start := int(capt.Node.StartByte())
	end := int(capt.Node.EndByte())

	var span *HighlightSpan

	if len(h.spans) > 0 {
		span = &h.spans[len(h.spans)-1]
	}

	if span == nil || span.Start != start || span.End != end {
		h.spans = append(h.spans, HighlightSpan{
			Start:    start,
			End:      end,
			Color:    CharFgColorUndefined,
			captures: make([]int, 0, 2),
		})
		span = &h.spans[len(h.spans)-1]
	}

	span.captures = append(span.captures, int(capt.Index))

	if slices.Contains(span.captures, 0 /*variable*/) {
		span.Color = CharFgColorVariable
	} else if slices.Contains(span.captures, 18 /*keyword*/) {
		span.Color = CharFgColorKeyword
	} else if slices.Contains(span.captures, 9 /*comment*/) {
		span.Color = CharFgColorComment
	} else {
		span.Color = CharFgColorUndefined
	}
}

func (h *Highlighter) Start(idx int) {
	if h == nil {
		return
	}

	h.idx = idx
}

func (h *Highlighter) Next(l int) CharFgColor {
	if h != nil {
		for i := h.spanIdx; i < len(h.spans); i += 1 {
			span := h.spans[i]

			if h.idx < span.Start {
				continue
			}

			if h.idx < span.End {
				h.spanIdx = i
				h.idx += l

				return span.Color
			}
		}
	}

	h.idx += l

	return CharFgColorUndefined
}
