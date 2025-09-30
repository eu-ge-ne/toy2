package vt

var cprReq = []byte("\x1b[6n")

func Wchar(y, x int, b []byte) int {
	SetCursor(Sync, y, x)
	Sync.Write(b)
	Sync.Write(cprReq)

	w := <-Pos - x
	if w < 1 {
		panic("Wchar error")
	}

	return w
}
