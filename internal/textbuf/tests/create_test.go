package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/segbuf/textbuf"
)

func TestCreateEmpty(t *testing.T) {
	buf := textbuf.New()

	assert.Equal(t, "",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, 0, buf.Count())
	assert.Equal(t, 0, buf.LineCount())

	buf.Validate()
}
