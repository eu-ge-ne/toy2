package syntax

type highlightReq struct {
	startLn int
	endLn   int
	spans   chan Span
}

type Span struct {
	StartByte int
	EndByte   int
	Name      string
}
