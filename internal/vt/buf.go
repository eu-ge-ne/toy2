package vt

var Buf = &bufOut{buf: make([]byte, 1024*64)}

type bufOut struct {
	buf []byte
	i   int
}

func (o *bufOut) Write(chunks ...[]byte) {
	for _, chunk := range chunks {
		j := o.i + len(chunk)

		if j > len(o.buf) {
			o.buf = make([]byte, j)
		}

		copy(o.buf[o.i:], chunk)

		o.i = j
	}
}

func (o *bufOut) Flush() {
	j := o.i
	o.i = 0

	Sync.Write(o.buf[:j])
}
