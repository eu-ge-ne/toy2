package syntax

type CharFgColor int

const (
	CharFgColorUndefined CharFgColor = iota
	CharFgColorVisible
	CharFgColorWhitespace
	CharFgColorEmpty
	CharFgColorDelimiter
)
