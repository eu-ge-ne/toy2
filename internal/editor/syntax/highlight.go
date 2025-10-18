package syntax

import (
	"slices"

	treeSitter "github.com/tree-sitter/go-tree-sitter"
)

type Highlight struct {
	spans   []span
	spanIdx int
	idx     int
}

type span struct {
	start    int
	end      int
	captures []int
	color    CharFgColor
}

func newHighlight() *Highlight {
	return &Highlight{
		spans: make([]span, 0, 1024),
	}
}

func (h *Highlight) AddCapture(capt treeSitter.QueryCapture) {
	start := int(capt.Node.StartByte())
	end := int(capt.Node.EndByte())

	var s *span

	if len(h.spans) > 0 {
		s = &h.spans[len(h.spans)-1]
	}

	if s == nil || s.start != start || s.end != end {
		h.spans = append(h.spans, span{
			start:    start,
			end:      end,
			captures: make([]int, 0, 2),
			color:    CharFgColorUndefined,
		})
		s = &h.spans[len(h.spans)-1]
	}

	s.captures = append(s.captures, int(capt.Index))

	if slices.Contains(s.captures, 0 /*variable*/) {
		s.color = CharFgColorVariable
	} else if slices.Contains(s.captures, 18) {
		s.color = CharFgColorKeyword
	} else if slices.Contains(s.captures, 9 /*comment*/) {
		s.color = CharFgColorComment
	} else {
		s.color = CharFgColorUndefined
	}
}

func (h *Highlight) Start(idx int) {
	if h == nil {
		return
	}

	h.idx = idx
}

func (h *Highlight) Next(l int) CharFgColor {
	if h != nil {
		for i := h.spanIdx; i < len(h.spans); i += 1 {
			span := h.spans[i]

			if h.idx < span.start {
				continue
			}

			if h.idx < span.end {
				h.spanIdx = i
				h.idx += l

				return span.color
			}
		}
	}

	h.idx += l

	return CharFgColorUndefined
}
