package syntax

type Highlight struct {
	spans chan span
	span  span
	idx   int
}

type span struct {
	startIdx int
	endIdx   int
	name     string
}

func (h *Highlight) Next(l int) string {
	if h == nil {
		return ""
	}

	defer func() { h.idx += l }()

	if h.idx >= h.span.endIdx {
		if spn, ok := <-h.spans; ok {
			h.span = spn
		}
	}

	if h.idx >= h.span.startIdx && h.idx < h.span.endIdx {
		return h.span.name
	}

	return ""
}
