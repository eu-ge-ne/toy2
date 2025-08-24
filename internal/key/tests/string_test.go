package key_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/key"
)

func TestStringa(t *testing.T) {
	k, n, ok := key.Parse([]byte("a"))
	assert.True(t, ok)
	assert.Equal(t, 1, n)
	assert.Equal(t, key.Key{
		Name: "a",
		Text: "a",
	}, k)
}
