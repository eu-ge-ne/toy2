package key_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/key"
)

func TestLEFT(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[1D"))
	assert.True(t, ok)
	assert.Equal(t, n, 4)
	assert.Equal(t, key.Key{
		Name:    "LEFT",
		KeyCode: 1,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[1;5D"))
	assert.True(t, ok)
	assert.Equal(t, n, 6)
	assert.Equal(t, key.Key{
		Name:    "LEFT",
		KeyCode: 1,
		Ctrl:    true,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[1;3D"))
	assert.True(t, ok)
	assert.Equal(t, n, 6)
	assert.Equal(t, key.Key{
		Name:    "LEFT",
		KeyCode: 1,
		Alt:     true,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[1;2D"))
	assert.True(t, ok)
	assert.Equal(t, n, 6)
	assert.Equal(t, key.Key{
		Name:    "LEFT",
		KeyCode: 1,
		Shift:   true,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[1;1:1D"))
	assert.True(t, ok)
	assert.Equal(t, n, 8)
	assert.Equal(t, key.Key{
		Name:    "LEFT",
		KeyCode: 1,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[1;1:2D"))
	assert.True(t, ok)
	assert.Equal(t, n, 8)
	assert.Equal(t, key.Key{
		Name:    "LEFT",
		KeyCode: 1,
		Event:   key.EventRepeat,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[1;1:3D"))
	assert.True(t, ok)
	assert.Equal(t, n, 8)
	assert.Equal(t, key.Key{
		Name:    "LEFT",
		KeyCode: 1,
		Event:   key.EventRelease,
	}, k)
}
