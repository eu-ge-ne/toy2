package vt

import (
	"fmt"
	"io"
)

func ECH(out io.Writer, n int) {
	fmt.Fprintf(out, "\x1b[%dX", n)
}
