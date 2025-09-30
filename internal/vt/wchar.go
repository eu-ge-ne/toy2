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

	t0 := time.Now()
	buf := make([]byte, 1024)
	done := make(chan int)

	go func() {
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				panic(err)
			}

			match := re.FindStringSubmatch(string(buf[:n]))

			if match != nil {
				x1, err := strconv.Atoi(match[1])
				if err != nil {
					panic(err)
				}

				done <- x1 - 1
				return
			}

			if time.Since(t0).Milliseconds() > 10 {
				done <- -1
				return
			}
		}
	}()

	x1 := <-done
	if x1 < 0 {
		panic("Wchar timeout")
	}

	return x1 - x
}
