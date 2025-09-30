package vt

import (
	"bytes"
	"os"
	"strconv"

	"github.com/eu-ge-ne/toy2/internal/key"
)

var Keys = make(chan key.Key, 100)
var Pos = make(chan int)

func Read() {
	var buf = make([]byte, 1024)

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}

		raw := buf[:n]

		for len(raw) > 0 {
			if key, bytesParsed, ok := key.Parse(raw); ok {
				Keys <- key
				raw = raw[bytesParsed:]
				continue
			}

			if match := cprRe.FindStringSubmatch(string(raw)); match != nil {
				x, err := strconv.Atoi(match[1])
				if err != nil {
					panic(err)
				}
				Pos <- x - 1
				break
			}

			if next_esc_i := bytes.Index(raw[1:], []byte{0x1b}); next_esc_i >= 0 {
				next_esc_i += 1
				raw = raw[next_esc_i:]
				continue
			}

			break
		}
	}
}
