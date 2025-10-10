package data

import (
	"github.com/eu-ge-ne/toy2/internal/key"
)

type Insert struct {
	data *Data
}

func (h *Insert) Match(k key.Key) bool {
	return len(k.Text) != 0
}

func (h *Insert) Handle(k key.Key) bool {
	return h.data.Insert(k.Text)
}
