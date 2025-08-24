package key_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/key"
)

func TestINSERT(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[2~"))
	assert.True(t, ok)
	assert.Equal(t, 4, n)
	assert.Equal(t, key.Key{
		Name:    "INSERT",
		KeyCode: 2,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[2;5~"))
	assert.True(t, ok)
	assert.Equal(t, 6, n)
	assert.Equal(t, key.Key{
		Name:    "INSERT",
		KeyCode: 2,
		Ctrl:    true,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[2;3~"))
	assert.True(t, ok)
	assert.Equal(t, 6, n)
	assert.Equal(t, key.Key{
		Name:    "INSERT",
		KeyCode: 2,
		Alt:     true,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[2;2~"))
	assert.True(t, ok)
	assert.Equal(t, 6, n)
	assert.Equal(t, key.Key{
		Name:    "INSERT",
		KeyCode: 2,
		Shift:   true,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[2;1:1~"))
	assert.True(t, ok)
	assert.Equal(t, 8, n)
	assert.Equal(t, key.Key{
		Name:    "INSERT",
		KeyCode: 2,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[2;1:2~"))
	assert.True(t, ok)
	assert.Equal(t, 8, n)
	assert.Equal(t, key.Key{
		Name:    "INSERT",
		KeyCode: 2,
		Event:   key.EventRepeat,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[2;1:3~"))
	assert.True(t, ok)
	assert.Equal(t, 8, n)
	assert.Equal(t, key.Key{
		Name:    "INSERT",
		KeyCode: 2,
		Event:   key.EventRelease,
	}, k)
}
