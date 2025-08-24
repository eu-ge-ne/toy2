package vt

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func write(b []byte) {
	for i := 0; i < len(b); {
		n, err := os.Stdout.Write(b[i:])
		if err != nil {
			panic(err)
		}

		i += n
	}
}

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
