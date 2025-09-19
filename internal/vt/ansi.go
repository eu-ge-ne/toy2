package vt

import "fmt"

func esc(code string) []byte {
	return []byte(fmt.Sprintf("\x1b%s", code))
}
