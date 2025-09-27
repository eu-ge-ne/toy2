package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestCreateEmpty(t *testing.T) {
	buf := textbuf.New("")

	assert.Equal(t, "",
		iterToStr(buf.Read(0)))
	assert.Equal(t, 0, buf.Count())
	assert.Equal(t, 0, buf.LineCount())

	buf.Validate()
}

func TestCreateNonEmpty(t *testing.T) {
	buf := textbuf.New("Lorem ipsum")

	assert.Equal(t, "Lorem ipsum",
		iterToStr(buf.Read(0)))
	assert.Equal(t, 11, buf.Count())
	assert.Equal(t, 1, buf.LineCount())

	buf.Validate()
}
