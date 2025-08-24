package vt

import "fmt"

const st = ""

func esc(code string) []byte {
	return []byte(fmt.Sprintf("\x1b%s", code))
}

func csi(code string) []byte {
	return []byte(fmt.Sprintf("\x1b[%s", code))
}

func osc(code string) []byte {
	return []byte(fmt.Sprintf("\x1b]%s", code))
}
