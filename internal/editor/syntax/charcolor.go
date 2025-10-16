package syntax

type CharColor int

const (
	CharColorUndefined CharColor = iota
	CharColorVisible
	CharColorWhitespace
	CharColorEmpty
	CharColorDelimiter
)
