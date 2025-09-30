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

	io.WriteString(out, text)

	*span -= l
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

	WriteSpaces(out, a)
	io.WriteString(out, text)

	*span -= a + l
}
