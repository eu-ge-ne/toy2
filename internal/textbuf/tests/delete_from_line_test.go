package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestDeleteFromLine(t *testing.T) {
	buf := textbuf.New("Lorem \nipsum \ndolor \nsit \namet")

	assert.Equal(t, 5, buf.LineCount())

	buf.Delete2ToEnd(3, 0)

	assert.Equal(t, "Lorem \nipsum \ndolor \n",
		iterToStr(buf.ReadToEnd(0)))
	assert.Equal(t, 21, buf.Count())
	assert.Equal(t, 4, buf.LineCount())
	buf.Validate()

	buf.Delete2ToEnd(1, 0)

	assert.Equal(t, "Lorem \n",
		iterToStr(buf.ReadToEnd(0)))
	assert.Equal(t, 7, buf.Count())
	assert.Equal(t, 2, buf.LineCount())
	buf.Validate()
}
