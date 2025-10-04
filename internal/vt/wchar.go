package vt

import (
	"math"
	"time"
)

func Wchar(y, x0 int, b []byte) int {
	for i := 0; i < 4; i += 1 {
		SetCursor(Sync, y, x0)
		Sync.Write(b)
		Sync.Write(cprReq)

		x1 := readCpr(time.Duration(math.Pow10(i)) * time.Millisecond)
		if x1 < 0 {
			continue
		}

		w := x1 - x0
		if w < 1 {
			panic("Wchar error")
		}

		return w
	}

	panic("Wchar timeout")
}
