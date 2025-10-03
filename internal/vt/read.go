package vt

import (
	"bytes"
	"os"
	"strconv"

	"github.com/eu-ge-ne/toy2/internal/key"
)

var Keys = make(chan key.Key, 100_000)
var Cpr = make(chan int)

func Read() {
	go func() {
		var buf []byte
		chunk := make([]byte, 1024)

		for {
			n, err := os.Stdin.Read(chunk)
			if err != nil {
				panic(err)
			}

			buf = append(buf, chunk[:n]...)

			for len(buf) > 0 {
				if match := cprRe.FindSubmatch(buf); match != nil {
					x, err := strconv.Atoi(string(match[1]))
					if err != nil {
						panic(err)
					}

					Cpr <- x - 1
					loc := cprRe.FindIndex(buf)
					buf = append(buf[:loc[0]], buf[loc[1]:]...)
					continue
				}

				if key, n, ok := key.Parse(buf); ok {
					Keys <- key
					buf = buf[n:]
					continue
				}

				if n := bytes.IndexByte(buf[1:], 0x1b); n >= 0 {
					buf = buf[n+1:]
					continue
				}

				buf = nil
			}
		}
	}()
}
