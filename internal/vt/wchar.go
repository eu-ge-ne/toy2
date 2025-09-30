package vt

import (
	"os"
	"regexp"
	"strconv"
	"time"
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
	t0 := time.Now()
	buf := make([]byte, 1024)

	for {
		n, err := os.Stdin.Read(buf)

		if err != nil {
			panic(err)
		}

		match := re.FindStringSubmatch(string(buf[:n]))

		if match != nil {
			x, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}

			w := x - 1
			done <- w
			return
		}

		if time.Since(t0).Milliseconds() > 10 {
			panic("Wchar timeout")
		}
	}
}
