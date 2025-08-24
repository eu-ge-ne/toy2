package vt

import "fmt"

func ECH(n int) []byte {
	return fmt.Appendf(nil, "\x1b[%dX", n)
}
