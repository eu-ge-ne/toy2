package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestCreateEmpty(t *testing.T) {
	buf := textbuf.New()

	assert.Equal(t, "", buf.All())
	assert.Equal(t, 0, buf.Count())
	assert.Equal(t, 0, buf.LineCount())

	buf.Validate()
}
