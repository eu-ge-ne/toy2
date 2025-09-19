package vt

import (
	"bytes"
	"iter"
	"os"

	"github.com/eu-ge-ne/toy2/internal/key"
)

var readBuf = make([]byte, 1024)

func Read() iter.Seq[key.Key] {
	return func(yield func(key.Key) bool) {
		bytesRead, err := os.Stdin.Read(readBuf)

		if err != nil {
			panic(err)
		}

		if bytesRead == 0 {
			return
		}

		raw := readBuf[:bytesRead]

		for i := 0; i < bytesRead; {
			key, bytesParsed, ok := key.Parse(raw[i:])

			if !ok {
				next_esc_i := bytes.Index(raw[1:], []byte{0x1b})
				if next_esc_i < 0 {
					next_esc_i = len(raw)
				} else {
					next_esc_i += 1
				}
				i = next_esc_i
				continue
			}

			if !yield(key) {
				return
			}

			i += bytesParsed
		}
	}
}
