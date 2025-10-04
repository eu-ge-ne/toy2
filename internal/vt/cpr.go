package vt

import (
	"regexp"
	"strconv"
)

var cprReq = []byte("\x1b[6n")
var cprRe = regexp.MustCompile(`\x1b\[\d+;(\d+)R`)

func parseCpr(b []byte) (cpr int, n int, ok bool) {
	match := cprRe.FindSubmatch(b)
	if match == nil {
		return
	}

	cpr, err := strconv.Atoi(string(match[1]))
	if err != nil {
		panic(err)
	}

	cpr -= 1
	n = cprRe.FindIndex(b)[1]
	ok = true

	return
}
