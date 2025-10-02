package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Handler interface {
	Match(key.Key) bool
	Handle(key.Key) bool
}
