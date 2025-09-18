package vt

import (
	"fmt"
	"unicode/utf8"
)

func WriteSpaces(w int) []byte {
	return fmt.Appendf(nil, " \x1b[%db", w-1)
}

func WriteText(span *int, text string) []byte {
	l := utf8.RuneCountInString(text)

	if l > *span {
		runes := []rune(text)
		text = string(runes[:*span])
		l = utf8.RuneCountInString(text)
	}

	*span -= l

	return []byte(text)
}

func WriteTextCenter(span *int, text string) []byte {
	l := utf8.RuneCountInString(text)

	if l > *span {
		runes := []rune(text)
		text = string(runes[:*span])
		l = utf8.RuneCountInString(text)
	}

	ab := *span - l
	a := ab / 2
	b := ab - a

	*span -= l

	return fmt.Appendf(nil, "%*s%s%*s", a, " ", text, b, " ")
}
