package textbuf_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestCreateEmpty(t *testing.T) {
	buf := textbuf.New()

	assert.Equal(t, "", std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 0, buf.Count())
	assert.Equal(t, 0, buf.LineCount())

	buf.Validate()
}
