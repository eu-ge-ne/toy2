package vt

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var SaveCursor = []byte("\x1b7")
var RestoreCursor = []byte("\x1b8")

var HideCursor = []byte("\x1b[?25l")
var ShowCursor = []byte("\x1b[?25h")

var CursorDown = []byte("\x1b[B")

func SetCursor(out io.Writer, y int, x int) {
	fmt.Fprintf(out, "\x1b[%d;%dH", y+1, x+1)
}

var cprReq = []byte("\x1b[6n")
var re = regexp.MustCompile(`\x1b\[\d+;(\d+)R`)

// TODO
func MeasureCursor(y, x int, b []byte) int {
	buf := make([]byte, 1024)

	SetCursor(Sync, y, x)
	Sync.Write(b)
	Sync.Write(cprReq)

	for i := 0; i < 4; i += 1 {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}

		if n == 0 {
			continue
		}

		match := re.FindStringSubmatch(string(buf[:n]))
		if match == nil {
			continue
		}

		x2, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		return x2 - 1 - x
	}

	panic("cursor.measure(): timeout")
}
