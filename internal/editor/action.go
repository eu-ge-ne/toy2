package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Action interface {
	Match(key.Key) bool
	Run(key.Key) bool
}
