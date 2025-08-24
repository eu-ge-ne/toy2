package key_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/key"
)

func TestESC(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[27u"))
	assert.True(t, ok)
	assert.Equal(t, 5, n)
	assert.Equal(t, key.Key{
		Name:    "ESC",
		KeyCode: 27,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[27;5u"))
	assert.True(t, ok)
	assert.Equal(t, 7, n)
	assert.Equal(t, key.Key{
		Name:    "ESC",
		KeyCode: 27,
		Ctrl:    true,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[27;3u"))
	assert.True(t, ok)
	assert.Equal(t, 7, n)
	assert.Equal(t, key.Key{
		Name:    "ESC",
		KeyCode: 27,
		Alt:     true,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[27;2u"))
	assert.True(t, ok)
	assert.Equal(t, 7, n)
	assert.Equal(t, key.Key{
		Name:    "ESC",
		KeyCode: 27,
		Shift:   true,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[27;1:1u"))
	assert.True(t, ok)
	assert.Equal(t, 9, n)
	assert.Equal(t, key.Key{
		Name:    "ESC",
		KeyCode: 27,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[27;1:2u"))
	assert.True(t, ok)
	assert.Equal(t, 9, n)
	assert.Equal(t, key.Key{
		Name:    "ESC",
		KeyCode: 27,
		Event:   key.EventRepeat,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[27;1:3u"))
	assert.True(t, ok)
	assert.Equal(t, 9, n)
	assert.Equal(t, key.Key{
		Name:    "ESC",
		KeyCode: 27,
		Event:   key.EventRelease,
	}, k)
}
