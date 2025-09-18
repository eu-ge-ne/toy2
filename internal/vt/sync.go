package vt

import (
	"os"
)

var bsuBytes = csi("?2026h")
var esuBytes = csi("?2026l")

var Sync = &syncOut{}

type syncOut struct {
	c int
}

func (o *syncOut) Bsu() {
	if o.c == 0 {
		o.Write(bsuBytes)
	}

	o.c += 1
}

func (o *syncOut) Esu() {
	o.c -= 1

	if o.c == 0 {
		o.Write(esuBytes)
	}
}

func (o *syncOut) Write(b []byte) {
	for i := 0; i < len(b); {
		n, err := os.Stdout.Write(b[i:])
		if err != nil {
			panic(err)
		}

		i += n
	}
}
