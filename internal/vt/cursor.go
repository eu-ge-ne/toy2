package vt

import (
	"fmt"
	"os"
	"io"
	"regexp"
	"strconv"
)

var SaveCursor = esc("7")
var RestoreCursor = esc("8")

var HideCursor = csi("?25l")
var ShowCursor = csi("?25h")

var CursorDown = csi("B")

func SetCursor(out io.Writer, y int, x int) {
	out.Write(fmt.Appendf(nil, "\x1b[%d;%dH", y+1, x+1))
}

var cprReq = csi("6n")
var re = regexp.MustCompile(`\x1b\[\d+;(\d+)R`)

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
