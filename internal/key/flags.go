package key

import (
	"fmt"
)

type Flags int

const (
	FLAG_DISAMBIGUATE Flags = 1
	FLAG_EVENTS       Flags = 2
	FLAG_ALTERNATES   Flags = 4
	FLAG_ALLKEYS      Flags = 8
	FLAG_TEXT         Flags = 16
)

type FlagsMode int

const (
	MODE_ALL FlagsMode = 1 + iota
	MODE_SET
	MODE_RESET
)

func SetFlags(flags Flags, mode FlagsMode) []byte {
	return []byte(fmt.Sprintf("\x1b[=%d;%du", flags, mode))
}
