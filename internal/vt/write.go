package vt

import (
	"fmt"
	"io"
	"unicode/utf8"
)

func WriteSpaces(out io.Writer, w int) {
	fmt.Fprintf(out, " \x1b[%db", w-1)
}

func WriteText(out io.Writer, span *int, text string) {
	l := utf8.RuneCountInString(text)

	if l > *span {
		runes := []rune(text)
		text = string(runes[:*span])
		l = utf8.RuneCountInString(text)
	}

	*span -= l

	io.WriteString(out, text)
}

func WriteTextCenter(out io.Writer, span *int, text string) {
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

	fmt.Fprintf(out, "%*s%s%*s", a, " ", text, b, " ")
}
