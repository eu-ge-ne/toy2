package vt

import (
	"fmt"
	"io"
)

func ECH(out io.Writer, n int) {
	out.Write(fmt.Appendf(nil, "\x1b[%dX", n))
}
