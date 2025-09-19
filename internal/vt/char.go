package vt

import (
	"fmt"
)

type Attr = int

const (
	Default    Attr = 0
	Bold       Attr = 1
	Faint      Attr = 2
	Italicized Attr = 3
	Underlined Attr = 4
	Blink      Attr = 5
	Inverse    Attr = 7
	Invisible  Attr = 8
	CrossedOut Attr = 9

	DoublyUnderlined Attr = 21
	Normal           Attr = 22
	NotItalicized    Attr = 23
	NotUnderlined    Attr = 24
	Steady           Attr = 25
	Positive         Attr = 27
	Visible          Attr = 28
	NotCrossedOut    Attr = 29

	FgBlack   Attr = 30
	FgRed     Attr = 31
	FgGreen   Attr = 32
	FgYellow  Attr = 33
	FgBlue    Attr = 34
	FgMagenta Attr = 35
	FgCyan    Attr = 36
	FgWhite   Attr = 37
	FgDefault Attr = 39

	BgBlack   Attr = 40
	BgRed     Attr = 41
	BgGreen   Attr = 42
	BgYellow  Attr = 43
	BgBlue    Attr = 44
	BgMagenta Attr = 45
	BgCyan    Attr = 46
	BgWhite   Attr = 47
	BgDefault Attr = 49

	FgBrightBlack   Attr = 90
	FgBrightRed     Attr = 91
	FgBrightGreen   Attr = 92
	FgBrightYellow  Attr = 93
	FgBrightBlue    Attr = 94
	FgBrightMagenta Attr = 95
	FgBrightCyan    Attr = 96
	FgBrightWhite   Attr = 97

	BgBrightBlack   Attr = 100
	BgBrightRed     Attr = 101
	BgBrightGreen   Attr = 102
	BgBrightYellow  Attr = 103
	BgBrightBlue    Attr = 104
	BgBrightMagenta Attr = 105
	BgBrightCyan    Attr = 106
	BgBrightWhite   Attr = 107
)

func CharAttr(attr Attr) []byte {
	return fmt.Appendf(nil, "\x1b[%dm", attr)
}

type RGB [3]byte

func CharFg(rgb RGB) []byte {
	return fmt.Appendf(nil, "\x1b[38;2;%d;%d;%dm", rgb[0], rgb[1], rgb[2])
}

func CharBg(rgb RGB) []byte {
	return fmt.Appendf(nil, "\x1b[48;2;%d;%d;%dm", rgb[0], rgb[1], rgb[2])
}
