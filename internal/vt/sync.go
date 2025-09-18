package vt

import (
	"os"
)

var bsu = []byte("\x1b[?2026h")
var esu = []byte("\x1b[?2026l")

var Sync = &syncOut{}

type syncOut struct {
	c int
}

func (o *syncOut) Bsu() {
	if o.c == 0 {
		o.Write(bsu)
	}

	o.c += 1
}

func (o *syncOut) Esu() {
	o.c -= 1

	if o.c == 0 {
		o.Write(esu)
	}
}

func (o *syncOut) Write(chunks ...[]byte) {
	for _, chunk := range chunks {
		for i := 0; i < len(chunk); {
			n, err := os.Stdout.Write(chunk[i:])
			if err != nil {
				panic(err)
			}

			i += n
		}
	}
}
