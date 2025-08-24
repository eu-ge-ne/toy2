package vt

import (
	"os"

	"golang.org/x/term"

	"github.com/eu-ge-ne/toy2/internal/key"
)

func Init() func() {
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	write(csi("?1049h"))
	write(
		key.SetFlags(
			key.FLAG_DISAMBIGUATE+key.FLAG_ALTERNATES+key.FLAG_ALLKEYS+key.FLAG_TEXT,
			key.MODE_ALL,
		),
	)

	return func() {
		write(key.SetFlags(0, key.MODE_ALL))
		write(csi("?1049l"))
		write(ShowCursor)

		term.Restore(int(os.Stdin.Fd()), state)
	}
}

func GetSize() (int, int, error) {
	return term.GetSize(int(os.Stdin.Fd()))
}
