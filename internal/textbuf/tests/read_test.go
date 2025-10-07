package textbuf_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestReadEmpty(t *testing.T) {
	buf := textbuf.New()

	assert.Equal(t, "",
		buf.Read())
	buf.Validate()
}

func TestRead(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem ipsum dolor")

	assert.Equal(t, "ipsum ",
		iterToStr(buf.ReadSlice(6, 12)))
	buf.Validate()
}

func TestReadAtStartGTECount(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem")

	assert.Equal(t, "m",
		iterToStr(buf.ReadSlice(4, math.MaxInt)))
	assert.Equal(t, "",
		iterToStr(buf.ReadSlice(5, math.MaxInt)))
	assert.Equal(t, "",
		iterToStr(buf.ReadSlice(6, math.MaxInt)))

	buf.Validate()
}

func TestReadAtStartLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem")

	assert.Equal(t, "Lorem",
		iterToStr(buf.ReadSlice(0, math.MaxInt)))
	assert.Equal(t, "m",
		iterToStr(buf.ReadSlice(buf.Count()-1, math.MaxInt)))
	assert.Equal(t, "em",
		iterToStr(buf.ReadSlice(buf.Count()-2, math.MaxInt)))

	buf.Validate()
}
