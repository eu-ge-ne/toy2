package textbuf_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestReadEmpty(t *testing.T) {
	buf := textbuf.New()

	assert.Equal(t, "", std.IterToStr(buf.Read(0, math.MaxInt)))
	buf.Validate()
}

func TestRead(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem ipsum dolor")

	assert.Equal(t, "ipsum ", std.IterToStr(buf.Read(6, 12)))
	buf.Validate()
}

func TestReadAtStartGTECount(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem")

	assert.Equal(t, "m", std.IterToStr(buf.Read(4, math.MaxInt)))
	assert.Equal(t, "", std.IterToStr(buf.Read(5, math.MaxInt)))
	assert.Equal(t, "", std.IterToStr(buf.Read(6, math.MaxInt)))

	buf.Validate()
}

func TestReadAtStartLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem")

	assert.Equal(t, "Lorem", std.IterToStr(buf.Read(0, math.MaxInt)))
	assert.Equal(t, "m", std.IterToStr(buf.Read(buf.Count()-1, math.MaxInt)))
	assert.Equal(t, "em", std.IterToStr(buf.Read(buf.Count()-2, math.MaxInt)))

	buf.Validate()
}
