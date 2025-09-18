package vt

import (
	"fmt"
	"unicode/utf8"
)

func WriteSpaces(out out, w int) {
	out.Write(fmt.Appendf(nil, " \x1b[%db", w-1))
}

func WriteText(out out, span *int, text string) {
	l := utf8.RuneCountInString(text)

	if l > *span {
		runes := []rune(text)
		text = string(runes[:*span])
		l = utf8.RuneCountInString(text)
	}

	*span -= l

	out.Write([]byte(text))
}

func WriteTextCenter(out out, span *int, text string) {
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

	out.Write(fmt.Appendf(nil, "%*s%s%*s", a, " ", text, b, " "))
}
