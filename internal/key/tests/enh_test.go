package key_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/key"
)

func Test1DisambiguateEscapeCodes(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[27u"))
	assert.True(t, ok)
	assert.Equal(t, 5, n)
	assert.Equal(t, key.Key{
		Name:    "ESC",
		KeyCode: 27,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[1078;8u"))
	assert.True(t, ok)
	assert.Equal(t, 9, n)
	assert.Equal(t, key.Key{
		Name:    "ж",
		KeyCode: 1078,
		Shift:   true,
		Alt:     true,
		Ctrl:    true,
	}, k)
}

func Test14ReportAlternateKeys(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[1078:1046:59;8u"))
	assert.True(t, ok)
	assert.Equal(t, 17, n)
	assert.Equal(t, key.Key{
		Name:      "ж",
		KeyCode:   1078,
		ShiftCode: 1046,
		BaseCode:  59,
		Shift:     true,
		Alt:       true,
		Ctrl:      true,
	}, k)
}

func Test148_ReportAllKeysAsEscapeCodes(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[1078u"))
	assert.True(t, ok)
	assert.Equal(t, 7, n)
	assert.Equal(t, key.Key{
		Name:    "ж",
		KeyCode: 1078,
	}, k)
}

func Test14816ReportAssociatedText(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[1078:1046:59;2;1046u"))
	assert.True(t, ok)
	assert.Equal(t, 22, n)
	assert.Equal(t, key.Key{
		Name:      "ж",
		KeyCode:   1078,
		ShiftCode: 1046,
		BaseCode:  59,
		Text:      "Ж",
		Shift:     true,
	}, k)
}

func Test_1_4_8_16_2_ReportEventTypes(t *testing.T) {
	k, n, ok := key.Parse([]byte("\x1b[1078:1046:59;2:1;1046u"))
	assert.True(t, ok)
	assert.Equal(t, 24, n)
	assert.Equal(t, key.Key{
		Name:      "ж",
		KeyCode:   1078,
		ShiftCode: 1046,
		BaseCode:  59,
		Text:      "Ж",
		Shift:     true,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[1078:1046:59;2:2;1046u"))
	assert.True(t, ok)
	assert.Equal(t, 24, n)
	assert.Equal(t, key.Key{
		Name:      "ж",
		KeyCode:   1078,
		ShiftCode: 1046,
		BaseCode:  59,
		Event:     key.EventRepeat,
		Text:      "Ж",
		Shift:     true,
	}, k)

	k, n, ok = key.Parse([]byte("\x1b[1078::59;2:3u"))
	assert.True(t, ok)
	assert.Equal(t, 15, n)
	assert.Equal(t, key.Key{
		Name:     "ж",
		KeyCode:  1078,
		BaseCode: 59,
		Event:    key.EventRelease,
		Shift:    true,
	}, k)
}
