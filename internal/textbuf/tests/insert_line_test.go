package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestInsertInto0Line(t *testing.T) {
	tb := textbuf.New()

	tb.Insert2(0, 0, "Lorem ipsum")

	assert.Equal(t, "Lorem ipsum", tb.Read())
	assert.Equal(t, "Lorem ipsum", tb.ReadSlice2(0, 0, 1, 0))

	tb.Validate()
}

func TestInsertIntoALine(t *testing.T) {
	tb := textbuf.New()
	tb.Insert(0, "Lorem")

	tb.Insert2(0, 5, " ipsum")

	assert.Equal(t, "Lorem ipsum", tb.Read())
	assert.Equal(t, "Lorem ipsum", tb.ReadSlice2(0, 0, 1, 0))

	tb.Validate()
}

func TestInsertIntoALineWhichDoesNotExist(t *testing.T) {
	tb := textbuf.New()

	tb.Insert2(1, 0, "Lorem ipsum")

	assert.Equal(t, "", tb.Read())
	assert.Equal(t, "", tb.ReadSlice2(0, 0, 1, 0))

	tb.Validate()
}

func TestInsertIntoAColumnWhichDoesNotExist(t *testing.T) {
	tb := textbuf.New()

	tb.Insert2(0, 1, "Lorem ipsum")

	assert.Equal(t, "", tb.Read())
	assert.Equal(t, "", tb.ReadSlice2(0, 0, 1, 0))

	tb.Validate()
}
