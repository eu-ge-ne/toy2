package syntax

type highlightReq struct {
	ln0   int
	ln1   int
	spans chan Span
}

type Span struct {
	Start int
	End   int
	Kind  SpanKind
}

type SpanKind int

const (
	SpanKindNone SpanKind = 1000 + iota
	SpanKindVariable
	SpanKindKeyword
	SpanKindComment
)
