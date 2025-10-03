package vt

import (
	"bytes"
	"os"
	"strconv"

	"github.com/eu-ge-ne/toy2/internal/key"
)

var keys = make(chan key.Key)
var cprs = make(chan int)

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
				if key, n, ok := key.Parse(buf); ok {
					keys <- key
					buf = buf[n:]
					continue
				}

				if match := cprRe.FindSubmatch(buf); match != nil {
					x, err := strconv.Atoi(string(match[1]))
					if err != nil {
						panic(err)
					}

					cprs <- x - 1
					loc := cprRe.FindIndex(buf)
					buf = buf[loc[1]:]
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

func readCpr() int {
	for {
		select {
		case <-keys:
		case cpr := <-cprs:
			return cpr
		}
	}
}

func ReadKey() key.Key {
	for {
		select {
		case <-cprs:
		case key := <-keys:
			return key
		}
	}
}
