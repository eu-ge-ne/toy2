package vt

var wBuf = make([]byte, 1024*64)
var i int = 0

func WriteBuf(chunks ...[]byte) {
	for _, chunk := range chunks {
		j := i + len(chunk)

		if j > len(wBuf) {
			wBuf = make([]byte, j)
		}

		copy(wBuf[i:], chunk)

		i = j
	}
}

func FlushBuf() {
	j := i
	i = 0

	write(wBuf[:j])
}
