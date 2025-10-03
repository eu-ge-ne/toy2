package vt

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
