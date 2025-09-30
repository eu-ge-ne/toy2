package vt

import (
	"fmt"
	"io"
)

var SaveCursor = []byte("\x1b7")
var RestoreCursor = []byte("\x1b8")

var HideCursor = []byte("\x1b[?25l")
var ShowCursor = []byte("\x1b[?25h")

var CursorDown = []byte("\x1b[B")

func SetCursor(out io.Writer, y int, x int) {
	fmt.Fprintf(out, "\x1b[%d;%dH", y+1, x+1)
}
