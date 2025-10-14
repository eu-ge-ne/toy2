package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestInsertInto0Line(t *testing.T) {
	buf := textbuf.New()

	buf.InsertString2(0, 0, "Lorem ipsum")

	assert.Equal(t, "Lorem ipsum", buf.All())
	assert.Equal(t, "Lorem ipsum", buf.Read2(0, 0, 1, 0))

	buf.Validate()
}

func TestInsertIntoALine(t *testing.T) {
	buf := textbuf.New()
	buf.InsertString(0, "Lorem")

	buf.InsertString2(0, 5, " ipsum")

	assert.Equal(t, "Lorem ipsum", buf.All())
	assert.Equal(t, "Lorem ipsum", buf.Read2(0, 0, 1, 0))

	buf.Validate()
}

func TestInsertIntoALineWhichDoesNotExist(t *testing.T) {
	buf := textbuf.New()

	buf.InsertString2(1, 0, "Lorem ipsum")

	assert.Equal(t, "Lorem ipsum", buf.All())
	assert.Equal(t, "Lorem ipsum", buf.Read2(0, 0, 1, 0))

	buf.Validate()
}

func TestInsertIntoAColumnWhichDoesNotExist(t *testing.T) {
	buf := textbuf.New()

	buf.InsertString2(0, 1, "Lorem ipsum")

	assert.Equal(t, "Lorem ipsum", buf.All())
	assert.Equal(t, "Lorem ipsum", buf.Read2(0, 0, 1, 0))

	buf.Validate()
}
