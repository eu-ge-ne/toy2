package vt

import (
	"fmt"
)

func ECH(out out, n int) {
	out.Write(fmt.Appendf(nil, "\x1b[%dX", n))
}
