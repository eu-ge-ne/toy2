package vt

import (
	"bytes"
)

var Buf = &bufOut{}

type bufOut struct {
	buf bytes.Buffer
}

func (o *bufOut) Write(p []byte) (n int, err error) {
	n, err = o.buf.Write(p)
	if err != nil {
		panic(err)
	}
	return
}

func (o *bufOut) Flush() {
	o.buf.WriteTo(Sync)
	o.buf.Reset()
}
