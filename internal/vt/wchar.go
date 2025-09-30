package vt

import (
	"os"
	"regexp"
	"strconv"
)

var cprReq = []byte("\x1b[6n")
var re = regexp.MustCompile(`\x1b\[\d+;(\d+)R`)

func Wchar(y, x int, b []byte) int {
	SetCursor(Sync, y, x)
	Sync.Write(b)
	Sync.Write(cprReq)

	done := make(chan int)
	go measure(done)

	return <-done - x
}

func measure(done chan<- int) {
	buf := make([]byte, 1024)

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}

		match := re.FindStringSubmatch(string(buf[:n]))
		if match == nil {
			continue
		}

		x, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		done <- x - 1
		return
	}
}
