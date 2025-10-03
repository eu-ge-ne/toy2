package vt

import (
	"regexp"
)

var cprReq = []byte("\x1b[6n")
var cprRe = regexp.MustCompile(`\x1b\[\d+;(\d+)R`)

func Wchar(y, x int, b []byte) int {
	SetCursor(Sync, y, x)
	Sync.Write(b)
	Sync.Write(cprReq)

	w := readCpr() - x
	if w < 1 {
		panic("Wchar error")
	}

	return w
}
