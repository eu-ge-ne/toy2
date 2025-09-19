package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/eu-ge-ne/toy2/internal/key"
	"golang.org/x/term"
)

func main() {
	stdinState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), stdinState)

	os.Stdout.Write(
		key.SetFlags(
			key.FLAG_DISAMBIGUATE+key.FLAG_EVENTS+key.FLAG_ALTERNATES+key.FLAG_ALLKEYS+key.FLAG_TEXT,
			key.MODE_ALL,
		),
	)

	defer os.Stdout.Write(key.SetFlags(0, key.MODE_ALL))

	buf := make([]byte, 1024)
	j := 0

	for {
		bytesRead, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}

		if bytesRead == 0 {
			continue
		}

		raw := buf[:bytesRead]

		for i := 0; i < bytesRead; {
			key, bytesParsed, ok := key.Parse(raw[i:])

			if !ok {
				next_esc_i := bytes.Index(raw[1:], []byte{0x1b})
				if next_esc_i < 0 {
					next_esc_i = len(raw)
				} else {
					next_esc_i += 1
				}

				fmt.Printf("%d: %x\r\n", j, raw[i:next_esc_i])

				j += 1
				i = next_esc_i

				continue
			}

			fmt.Printf("%d: %+v\r\n", j, key)

			if key.Name == "c" && key.Ctrl {
				return
			}

			j += 1
			i += bytesParsed
		}
	}
}
