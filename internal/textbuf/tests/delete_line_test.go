package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestDeleteLine(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem \nipsum \ndolor \nsit \namet ")

	assert.Equal(t, 5, buf.LineCount())

	buf.DeletePosRange(4, 0, 5, 0)

	assert.Equal(t, "Lorem \nipsum \ndolor \nsit \n",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, 26, buf.Count())
	assert.Equal(t, 5, buf.LineCount())
	buf.Validate()

	buf.DeletePosRange(3, 0, 4, 0)

	assert.Equal(t, "Lorem \nipsum \ndolor \n",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, 21, buf.Count())
	assert.Equal(t, 4, buf.LineCount())
	buf.Validate()

	buf.DeletePosRange(2, 0, 3, 0)

	assert.Equal(t, "Lorem \nipsum \n",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, 14, buf.Count())
	assert.Equal(t, 3, buf.LineCount())
	buf.Validate()

	buf.DeletePosRange(1, 0, 2, 0)

	assert.Equal(t, "Lorem \n",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, 7, buf.Count())
	assert.Equal(t, 2, buf.LineCount())
	buf.Validate()
}
