package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestDeleteCharsFromTheBeginning(t *testing.T) {
	for n := 1; n <= 10; n += 1 {
		testDeleteHead(t, createTextBuf(), n)
	}
}

func TestDeleteCharsFromTheBeginningReversed(t *testing.T) {
	for n := 1; n <= 10; n += 1 {
		testDeleteHead(t, createTextBufReversed(), n)
	}
}

func TestDeleteCharsFromTheEnd(t *testing.T) {
	for n := 1; n <= 10; n += 1 {
		testDeleteTail(t, createTextBuf(), n)
	}
}

func TestDeleteCharsFromTheEndReversed(t *testing.T) {
	for n := 1; n <= 10; n += 1 {
		testDeleteTail(t, createTextBufReversed(), n)
	}
}

func TestDeleteCharsFromTheMiddle(t *testing.T) {
	for n := 1; n <= 10; n += 1 {
		testDeleteMiddle(t, createTextBuf(), n)
	}
}

func TestDeleteCharsFromTheMiddleReversed(t *testing.T) {
	for n := 1; n <= 10; n += 1 {
		testDeleteMiddle(t, createTextBufReversed(), n)
	}
}

func createTextBuf() textbuf.TextBuf {
	buf := textbuf.New()

	buf.Insert(buf.Count(), "Lorem")
	buf.Insert(buf.Count(), " ipsum")
	buf.Insert(buf.Count(), " dolor")
	buf.Insert(buf.Count(), " sit")
	buf.Insert(buf.Count(), " amet,")
	buf.Insert(buf.Count(), " consectetur")
	buf.Insert(buf.Count(), " adipiscing")
	buf.Insert(buf.Count(), " elit,")
	buf.Insert(buf.Count(), " sed")
	buf.Insert(buf.Count(), " do")
	buf.Insert(buf.Count(), " eiusmod")
	buf.Insert(buf.Count(), " tempor")
	buf.Insert(buf.Count(), " incididunt")
	buf.Insert(buf.Count(), " ut")
	buf.Insert(buf.Count(), " labore")
	buf.Insert(buf.Count(), " et")
	buf.Insert(buf.Count(), " dolore")
	buf.Insert(buf.Count(), " magna")
	buf.Insert(buf.Count(), " aliqua.")

	return buf
}

func createTextBufReversed() textbuf.TextBuf {
	buf := textbuf.New()

	buf.Insert(0, " aliqua.")
	buf.Insert(0, " magna")
	buf.Insert(0, " dolore")
	buf.Insert(0, " et")
	buf.Insert(0, " labore")
	buf.Insert(0, " ut")
	buf.Insert(0, " incididunt")
	buf.Insert(0, " tempor")
	buf.Insert(0, " eiusmod")
	buf.Insert(0, " do")
	buf.Insert(0, " sed")
	buf.Insert(0, " elit,")
	buf.Insert(0, " adipiscing")
	buf.Insert(0, " consectetur")
	buf.Insert(0, " amet,")
	buf.Insert(0, " sit")
	buf.Insert(0, " dolor")
	buf.Insert(0, " ipsum")
	buf.Insert(0, "Lorem")

	return buf
}

const text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

func testDeleteHead(t *testing.T, buf textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected, buf.Read())
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		i := min(buf.Count(), n)
		buf.DeleteSlice(0, i)
		expected = expected[i:]
	}

	assert.Equal(t, expected, buf.Read())
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}

func testDeleteTail(t *testing.T, buf textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected, buf.Read())
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		i := max(buf.Count()-n, 0)
		buf.DeleteSlice(i, buf.Count())
		expected = expected[0:i]
	}

	assert.Equal(t, expected, buf.Read())
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}

func testDeleteMiddle(t *testing.T, buf textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected, buf.Read())
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		pos := buf.Count() / 2
		i := min(buf.Count(), pos+n)
		buf.DeleteSlice(pos, i)
		expected = expected[0:pos] + expected[i:]
	}

	assert.Equal(t, expected, buf.Read())
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}
