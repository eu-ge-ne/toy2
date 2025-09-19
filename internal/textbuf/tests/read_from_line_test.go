package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestLineAtValidIndex(t *testing.T) {
	buf := textbuf.New("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "Lorem\nipsum\ndolor\nsit\namet",
		iterToStr(buf.Read2ToEnd(0, 0)))
	assert.Equal(t, "ipsum\ndolor\nsit\namet",
		iterToStr(buf.Read2ToEnd(1, 0)))
	assert.Equal(t, "dolor\nsit\namet",
		iterToStr(buf.Read2ToEnd(2, 0)))
	assert.Equal(t, "sit\namet",
		iterToStr(buf.Read2ToEnd(3, 0)))
	assert.Equal(t, "amet",
		iterToStr(buf.Read2ToEnd(4, 0)))

	buf.Validate()
}

func TestLineAtIndexGTELineCount(t *testing.T) {
	buf := textbuf.New("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "amet",
		iterToStr(buf.Read2ToEnd(4, 0)))
	assert.Equal(t, "",
		iterToStr(buf.Read2ToEnd(5, 0)))
	assert.Equal(t, "",
		iterToStr(buf.Read2ToEnd(6, 0)))

	buf.Validate()
}

func TestLineAtIndexLT0(t *testing.T) {
	buf := textbuf.New("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "Lorem\nipsum\ndolor\nsit\namet",
		iterToStr(buf.Read2ToEnd(0, 0)))
	assert.Equal(t, "amet",
		iterToStr(buf.Read2ToEnd(buf.LineCount()-1, 0)))
	assert.Equal(t, "sit\namet",
		iterToStr(buf.Read2ToEnd(buf.LineCount()-2, 0)))

	buf.Validate()
}
