package textbuf_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestDeleteFromLine(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem \nipsum \ndolor \nsit \namet")

	assert.Equal(t, 5, buf.LineCount())

	buf.Delete2(3, 0, math.MaxInt, math.MaxInt)

	assert.Equal(t, "Lorem \nipsum \ndolor \n", buf.All())
	assert.Equal(t, 21, buf.Count())
	assert.Equal(t, 4, buf.LineCount())
	buf.Validate()

	buf.Delete2(1, 0, math.MaxInt, math.MaxInt)

	assert.Equal(t, "Lorem \n", buf.All())
	assert.Equal(t, 7, buf.Count())
	assert.Equal(t, 2, buf.LineCount())
	buf.Validate()
}
