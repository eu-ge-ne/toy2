package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/segbuf/textbuf"
)

func TestLineAtValidIndex(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "Lorem\nipsum\ndolor\nsit\namet",
		iterToStr(buf.ReadPos(0, 0)))
	assert.Equal(t, "ipsum\ndolor\nsit\namet",
		iterToStr(buf.ReadPos(1, 0)))
	assert.Equal(t, "dolor\nsit\namet",
		iterToStr(buf.ReadPos(2, 0)))
	assert.Equal(t, "sit\namet",
		iterToStr(buf.ReadPos(3, 0)))
	assert.Equal(t, "amet",
		iterToStr(buf.ReadPos(4, 0)))

	buf.Validate()
}

func TestLineAtIndexGTELineCount(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "amet",
		iterToStr(buf.ReadPos(4, 0)))
	assert.Equal(t, "",
		iterToStr(buf.ReadPos(5, 0)))
	assert.Equal(t, "",
		iterToStr(buf.ReadPos(6, 0)))

	buf.Validate()
}

func TestLineAtIndexLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "Lorem\nipsum\ndolor\nsit\namet",
		iterToStr(buf.ReadPos(0, 0)))
	assert.Equal(t, "amet",
		iterToStr(buf.ReadPos(buf.LineCount()-1, 0)))
	assert.Equal(t, "sit\namet",
		iterToStr(buf.ReadPos(buf.LineCount()-2, 0)))

	buf.Validate()
}
