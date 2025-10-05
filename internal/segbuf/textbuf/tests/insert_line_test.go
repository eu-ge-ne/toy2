package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/segbuf/textbuf"
)

func TestInsertInto0Line(t *testing.T) {
	buf := textbuf.New()

	buf.InsertPos(0, 0, "Lorem ipsum")

	assert.Equal(t, "Lorem ipsum",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, "Lorem ipsum",
		iterToStr(buf.ReadPosRange(0, 0, 1, 0)))

	buf.Validate()
}

func TestInsertIntoALine(t *testing.T) {
	buf := textbuf.New()
	buf.InsertIndex(0, "Lorem")

	buf.InsertPos(0, 5, " ipsum")

	assert.Equal(t, "Lorem ipsum",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, "Lorem ipsum",
		iterToStr(buf.ReadPosRange(0, 0, 1, 0)))

	buf.Validate()
}

func TestInsertIntoALineWhichDoesNotExist(t *testing.T) {
	buf := textbuf.New()

	buf.InsertPos(1, 0, "Lorem ipsum")

	assert.Equal(t, "",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, "",
		iterToStr(buf.ReadPosRange(0, 0, 1, 0)))

	buf.Validate()
}

func TestInsertIntoAColumnWhichDoesNotExist(t *testing.T) {
	buf := textbuf.New()

	buf.InsertPos(0, 1, "Lorem ipsum")

	assert.Equal(t, "",
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, "",
		iterToStr(buf.ReadPosRange(0, 0, 1, 0)))

	buf.Validate()
}
