package key_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/key"
)

func TestInvalidCSIflagsu(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[?31u"))
	assert.False(t, ok)
	assert.Equal(t, 0, n)
	assert.Equal(t, key.Key{}, k)
}

func TestInvalidCSI1z(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[1z"))
	assert.False(t, ok)
	assert.Equal(t, 0, n)
	assert.Equal(t, key.Key{}, k)
}

func TestInvalidCSI1(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[1"))
	assert.False(t, ok)
	assert.Equal(t, 0, n)
	assert.Equal(t, key.Key{}, k)
}

func TestInvalidCSI(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b["))
	assert.False(t, ok)
	assert.Equal(t, 0, n)
	assert.Equal(t, key.Key{}, k)
}
