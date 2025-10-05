package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/segbuf/textbuf"
)

func TestDeleteFromLine(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem \nipsum \ndolor \nsit \namet")

	assert.Equal(t, 5, buf.LineCount())

	buf.DeletePos(3, 0)

	assert.Equal(t, "Lorem \nipsum \ndolor \n",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, 21, buf.Count())
	assert.Equal(t, 4, buf.LineCount())
	buf.Validate()

	buf.DeletePos(1, 0)

	assert.Equal(t, "Lorem \n",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, 7, buf.Count())
	assert.Equal(t, 2, buf.LineCount())
	buf.Validate()
}
