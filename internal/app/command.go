package app

import (
	"github.com/eu-ge-ne/toy2/internal/key"
	"github.com/eu-ge-ne/toy2/internal/palette"
)

type Command interface {
	Option() *palette.Option
	Match(key *key.Key) bool
	Run()
}
