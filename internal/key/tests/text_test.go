package key_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/key"
)

func Test_a(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[97;;97u"))
	assert.True(t, ok)
	assert.Equal(t, 9, n)
	assert.Equal(t, key.Key{
		Name:    "a",
		KeyCode: 97,
		Text:    "a",
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[97;3u"))
	assert.True(t, ok)
	assert.Equal(t, 7, n)
	assert.Equal(t, key.Key{
		Name:    "a",
		KeyCode: 97,
		Mods:    key.Alt,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[97;5u"))
	assert.True(t, ok)
	assert.Equal(t, 7, n)
	assert.Equal(t, key.Key{
		Name:    "a",
		KeyCode: 97,
		Mods:    key.Ctrl,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[97;9u"))
	assert.True(t, ok)
	assert.Equal(t, 7, n)
	assert.Equal(t, key.Key{
		Name:    "a",
		KeyCode: 97,
		Mods:    key.Super,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[97;1:1u"))
	assert.True(t, ok)
	assert.Equal(t, 9, n)
	assert.Equal(t, key.Key{
		Name:    "a",
		KeyCode: 97,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[97;1:2u"))
	assert.True(t, ok)
	assert.Equal(t, 9, n)
	assert.Equal(t, key.Key{
		Name:    "a",
		KeyCode: 97,
		Event:   key.Repeat,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[97;1:3u"))
	assert.True(t, ok)
	assert.Equal(t, 9, n)
	assert.Equal(t, key.Key{
		Name:    "a",
		KeyCode: 97,
		Event:   key.Release,
	}, k)
}

func Test_A(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[97:65;2;65u"))
	assert.True(t, ok)
	assert.Equal(t, 13, n)
	assert.Equal(t, key.Key{
		Name:      "a",
		KeyCode:   97,
		ShiftCode: 65,
		Text:      "A",
		Mods:      key.Shift,
	}, k)
}
