package vt

import (
	"io"
	"os"

	"golang.org/x/term"

	"github.com/eu-ge-ne/toy2/internal/key"
)

func Init() func() {
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	io.WriteString(Sync, "\x1b[?1049h")
	Sync.Write(
		key.SetFlags(
			key.FLAG_DISAMBIGUATE+key.FLAG_ALTERNATES+key.FLAG_ALLKEYS+key.FLAG_TEXT,
			key.MODE_ALL,
		),
	)

	return func() {
		Sync.Write(key.SetFlags(0, key.MODE_ALL))
		io.WriteString(Sync, "\x1b[?1049l")
		Sync.Write(ShowCursor)

		term.Restore(int(os.Stdin.Fd()), state)
	}
}

func GetSize() (int, int, error) {
	return term.GetSize(int(os.Stdin.Fd()))
}
