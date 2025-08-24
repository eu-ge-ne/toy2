package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestReadEmpty(t *testing.T) {
	buf := textbuf.New("")

	assert.Equal(t, "",
		iterToStr(buf.ReadToEnd(0)))
	buf.Validate()
}

func TestRead(t *testing.T) {
	buf := textbuf.New("Lorem ipsum dolor")

	assert.Equal(t, "ipsum ",
		iterToStr(buf.Read(6, 12)))
	buf.Validate()
}

func TestReadAtStartGTECount(t *testing.T) {
	buf := textbuf.New("Lorem")

	assert.Equal(t, "m",
		iterToStr(buf.ReadToEnd(4)))
	assert.Equal(t, "",
		iterToStr(buf.ReadToEnd(5)))
	assert.Equal(t, "",
		iterToStr(buf.ReadToEnd(6)))

	buf.Validate()
}

func TestReadAtStartLT0(t *testing.T) {
	buf := textbuf.New("Lorem")

	assert.Equal(t, "Lorem",
		iterToStr(buf.ReadToEnd(0)))
	assert.Equal(t, "m",
		iterToStr(buf.ReadToEnd(buf.Count()-1)))
	assert.Equal(t, "em",
		iterToStr(buf.ReadToEnd(buf.Count()-2)))

	buf.Validate()
}
