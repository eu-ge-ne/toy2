package vt

import (
	"bytes"
	"os"
	"strconv"

	"github.com/eu-ge-ne/toy2/internal/key"
)

var Keys = make(chan key.Key, 1_000_000)
var Pos = make(chan int)

func Read() {
	var buf []byte
	chunk := make([]byte, 1024)

	for {
		if n, err := os.Stdin.Read(chunk); err == nil {
			buf = append(buf, chunk[:n]...)
		} else {
			panic(err)
		}

		for len(buf) > 0 {
			if key, n, ok := key.Parse(buf); ok {
				Keys <- key
				buf = buf[n:]
				continue
			}

			if match := cprRe.FindSubmatch(buf); match != nil {
				x, err := strconv.Atoi(string(match[1]))
				if err != nil {
					panic(err)
				}
				Pos <- x - 1
				loc := cprRe.FindIndex(buf)
				buf = buf[loc[1]:]
				break
			}

			if n := bytes.Index(buf[1:], []byte{0x1b}); n >= 0 {
				n += 1
				buf = buf[n:]
				continue
			}

			buf = nil
		}
	}
}
