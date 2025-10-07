package textbuf_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestLineAtValidIndex(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "Lorem\nipsum\ndolor\nsit\namet", buf.ReadSlice2(0, 0, math.MaxInt, math.MaxInt))
	assert.Equal(t, "ipsum\ndolor\nsit\namet", buf.ReadSlice2(1, 0, math.MaxInt, math.MaxInt))
	assert.Equal(t, "dolor\nsit\namet", buf.ReadSlice2(2, 0, math.MaxInt, math.MaxInt))
	assert.Equal(t, "sit\namet", buf.ReadSlice2(3, 0, math.MaxInt, math.MaxInt))
	assert.Equal(t, "amet", buf.ReadSlice2(4, 0, math.MaxInt, math.MaxInt))

	buf.Validate()
}

func TestLineAtIndexGTELineCount(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "amet", buf.ReadSlice2(4, 0, math.MaxInt, math.MaxInt))
	assert.Equal(t, "", buf.ReadSlice2(5, 0, math.MaxInt, math.MaxInt))
	assert.Equal(t, "", buf.ReadSlice2(6, 0, math.MaxInt, math.MaxInt))

	buf.Validate()
}

func TestLineAtIndexLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "Lorem\nipsum\ndolor\nsit\namet", buf.ReadSlice2(0, 0, math.MaxInt, math.MaxInt))
	assert.Equal(t, "amet", buf.ReadSlice2(buf.LineCount()-1, 0, math.MaxInt, math.MaxInt))
	assert.Equal(t, "sit\namet", buf.ReadSlice2(buf.LineCount()-2, 0, math.MaxInt, math.MaxInt))

	buf.Validate()
}
