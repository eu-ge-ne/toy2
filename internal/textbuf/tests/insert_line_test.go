package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestInsertInto0Line(t *testing.T) {
	buf := textbuf.New("")

	buf.Insert2(0, 0, "Lorem ipsum")

	assert.Equal(t, "Lorem ipsum",
		iterToStr(buf.ReadToEnd(0)))
	assert.Equal(t, "Lorem ipsum",
		iterToStr(buf.Read2(0, 0, 1, 0)))

	buf.Validate()
}

func TestInsertIntoALine(t *testing.T) {
	buf := textbuf.New("")
	buf.Insert(0, "Lorem")

	buf.Insert2(0, 5, " ipsum")

	assert.Equal(t, "Lorem ipsum",
		iterToStr(buf.ReadToEnd(0)))
	assert.Equal(t, "Lorem ipsum",
		iterToStr(buf.Read2(0, 0, 1, 0)))

	buf.Validate()
}

func TestInsertIntoALineWhichDoesNotExist(t *testing.T) {
	buf := textbuf.New("")

	buf.Insert2(1, 0, "Lorem ipsum")

	assert.Equal(t, "",
		iterToStr(buf.ReadToEnd(0)))
	assert.Equal(t, "",
		iterToStr(buf.Read2(0, 0, 1, 0)))

	buf.Validate()
}

func TestInsertIntoAColumnWhichDoesNotExist(t *testing.T) {
	buf := textbuf.New("")

	buf.Insert2(0, 1, "Lorem ipsum")

	assert.Equal(t, "",
		iterToStr(buf.ReadToEnd(0)))
	assert.Equal(t, "",
		iterToStr(buf.Read2(0, 0, 1, 0)))

	buf.Validate()
}
