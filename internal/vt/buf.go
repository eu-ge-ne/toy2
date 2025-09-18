package vt

var Buf = &bufOut{buf: make([]byte, 1024*64)}

type bufOut struct {
	buf []byte
	i   int
}

func (o *bufOut) Write(p []byte) (n int, err error) {
	j := o.i + len(p)

	if j > len(o.buf) {
		o.buf = make([]byte, j)
	}

	copy(o.buf[o.i:], p)

	o.i = j

	return len(p), nil
}

func (o *bufOut) Flush() {
	j := o.i
	o.i = 0

	Sync.Write(o.buf[:j])
}
