package editor

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Handler interface {
	Match(key key.Key) bool
	Handle(key key.Key) bool
}
