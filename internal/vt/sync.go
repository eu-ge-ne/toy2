package vt

var bsuBytes = csi("?2026h")
var esuBytes = csi("?2026l")

var c = 0

func Bsu() {
	if c == 0 {
		write(bsuBytes)
	}

	c += 1
}

func Esu() {
	c -= 1

	if c == 0 {
		write(esuBytes)
	}
}
