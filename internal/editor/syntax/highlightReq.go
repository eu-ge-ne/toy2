package syntax

type highlightReq struct {
	ln0   int
	ln1   int
	spans chan *HighlightSpan
}

type HighlightSpan struct {
	Start int
	End   int
	Color CharFgColor

	captures []int
}
