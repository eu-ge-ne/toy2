package textbuf_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestLineAtValidIndex(t *testing.T) {
	buf := textbuf.New()
	buf.Append([]byte("Lorem\nipsum\ndolor\nsit\namet"))

	assert.Equal(t, "Lorem\nipsum\ndolor\nsit\namet",
		std.IterToStr(buf.Read2(0, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "ipsum\ndolor\nsit\namet",
		std.IterToStr(buf.Read2(1, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "dolor\nsit\namet",
		std.IterToStr(buf.Read2(2, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "sit\namet",
		std.IterToStr(buf.Read2(3, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "amet",
		std.IterToStr(buf.Read2(4, 0, math.MaxInt, math.MaxInt)))

	buf.Validate()
}

func TestLineAtIndexGTELineCount(t *testing.T) {
	buf := textbuf.New()
	buf.Append([]byte("Lorem\nipsum\ndolor\nsit\namet"))

	assert.Equal(t, "amet",
		std.IterToStr(buf.Read2(4, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "",
		std.IterToStr(buf.Read2(5, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "",
		std.IterToStr(buf.Read2(6, 0, math.MaxInt, math.MaxInt)))

	buf.Validate()
}

func TestLineAtIndexLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append([]byte("Lorem\nipsum\ndolor\nsit\namet"))

	assert.Equal(t, "Lorem\nipsum\ndolor\nsit\namet",
		std.IterToStr(buf.Read2(0, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "amet",
		std.IterToStr(buf.Read2(buf.LineCount()-1, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "sit\namet",
		std.IterToStr(buf.Read2(buf.LineCount()-2, 0, math.MaxInt, math.MaxInt)))

	buf.Validate()
}
