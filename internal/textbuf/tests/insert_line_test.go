package textbuf_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestInsertInto0Line(t *testing.T) {
	buf := textbuf.New()

	buf.Insert2(0, 0, "Lorem ipsum")

	assert.Equal(t, "Lorem ipsum", std.IterToStr(buf.Read(0, math.MaxInt)))
	assert.Equal(t, "Lorem ipsum", std.IterToStr(buf.Read2(0, 0, 0, 11)))

	buf.Validate()
}

func TestInsertIntoALine(t *testing.T) {
	buf := textbuf.New()
	buf.Insert(0, "Lorem")

	buf.Insert2(0, 5, " ipsum")

	assert.Equal(t, "Lorem ipsum", std.IterToStr(buf.Read(0, math.MaxInt)))
	assert.Equal(t, "Lorem ipsum", std.IterToStr(buf.Read2(0, 0, 0, 11)))

	buf.Validate()
}

func TestInsertIntoALineWhichDoesNotExist(t *testing.T) {
	buf := textbuf.New()

	buf.Insert2(1, 0, "Lorem ipsum")

	assert.Equal(t, "Lorem ipsum", std.IterToStr(buf.Read(0, math.MaxInt)))
	assert.Equal(t, "Lorem ipsum", std.IterToStr(buf.Read2(0, 0, 0, 11)))

	buf.Validate()
}

func TestInsertIntoAColumnWhichDoesNotExist(t *testing.T) {
	buf := textbuf.New()

	buf.Insert2(0, 1, "Lorem ipsum")

	assert.Equal(t, "Lorem ipsum", std.IterToStr(buf.Read(0, math.MaxInt)))
	assert.Equal(t, "Lorem ipsum", std.IterToStr(buf.Read2(0, 0, 0, 11)))

	buf.Validate()
}
