package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/segbuf/textbuf"
)

func TestReadEmpty(t *testing.T) {
	buf := textbuf.New()

	assert.Equal(t, "",
		iterToStr(buf.ReadIndex(0)))
	buf.Validate()
}

func TestRead(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem ipsum dolor")

	assert.Equal(t, "ipsum ",
		iterToStr(buf.ReadIndexRange(6, 12)))
	buf.Validate()
}

func TestReadAtStartGTECount(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem")

	assert.Equal(t, "m",
		iterToStr(buf.ReadIndex(4)))
	assert.Equal(t, "",
		iterToStr(buf.ReadIndex(5)))
	assert.Equal(t, "",
		iterToStr(buf.ReadIndex(6)))

	buf.Validate()
}

func TestReadAtStartLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem")

	assert.Equal(t, "Lorem",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, "m",
		iterToStr(buf.ReadIndex(buf.Count()-1)))
	assert.Equal(t, "em",
		iterToStr(buf.ReadIndex(buf.Count()-2)))

	buf.Validate()
}
