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

func (o *syncOut) Write(p []byte) (n int, err error) {
	n, err = os.Stdout.Write(p)
	if err != nil {
		panic(err)
	}
	return
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
