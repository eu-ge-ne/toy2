package syntax

type CharFgColor int

const (
	CharFgColorUndefined CharFgColor = iota
	CharFgColorVisible
	CharFgColorWhitespace
	CharFgColorEmpty
	CharFgColorVariable
	CharFgColorKeyword
	CharFgColorComment
)
