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

func createTextBuf() *textbuf.TextBuf {
	buf := textbuf.New()

	buf.InsertString(buf.Count(), "Lorem")
	buf.InsertString(buf.Count(), " ipsum")
	buf.InsertString(buf.Count(), " dolor")
	buf.InsertString(buf.Count(), " sit")
	buf.InsertString(buf.Count(), " amet,")
	buf.InsertString(buf.Count(), " consectetur")
	buf.InsertString(buf.Count(), " adipiscing")
	buf.InsertString(buf.Count(), " elit,")
	buf.InsertString(buf.Count(), " sed")
	buf.InsertString(buf.Count(), " do")
	buf.InsertString(buf.Count(), " eiusmod")
	buf.InsertString(buf.Count(), " tempor")
	buf.InsertString(buf.Count(), " incididunt")
	buf.InsertString(buf.Count(), " ut")
	buf.InsertString(buf.Count(), " labore")
	buf.InsertString(buf.Count(), " et")
	buf.InsertString(buf.Count(), " dolore")
	buf.InsertString(buf.Count(), " magna")
	buf.InsertString(buf.Count(), " aliqua.")

	return buf
}

func createTextBufReversed() *textbuf.TextBuf {
	buf := textbuf.New()

	buf.InsertString(0, " aliqua.")
	buf.InsertString(0, " magna")
	buf.InsertString(0, " dolore")
	buf.InsertString(0, " et")
	buf.InsertString(0, " labore")
	buf.InsertString(0, " ut")
	buf.InsertString(0, " incididunt")
	buf.InsertString(0, " tempor")
	buf.InsertString(0, " eiusmod")
	buf.InsertString(0, " do")
	buf.InsertString(0, " sed")
	buf.InsertString(0, " elit,")
	buf.InsertString(0, " adipiscing")
	buf.InsertString(0, " consectetur")
	buf.InsertString(0, " amet,")
	buf.InsertString(0, " sit")
	buf.InsertString(0, " dolor")
	buf.InsertString(0, " ipsum")
	buf.InsertString(0, "Lorem")

	return buf
}

const text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

func testDeleteHead(t *testing.T, buf *textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected, buf.All())
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		i := min(buf.Count(), n)
		buf.Delete(0, i)
		expected = expected[i:]
	}

	assert.Equal(t, expected, buf.All())
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}

func testDeleteTail(t *testing.T, buf *textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected, buf.All())
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		i := max(buf.Count()-n, 0)
		buf.Delete(i, buf.Count())
		expected = expected[0:i]
	}

	assert.Equal(t, expected, buf.All())
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}

func testDeleteMiddle(t *testing.T, buf *textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected, buf.All())
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		pos := buf.Count() / 2
		i := min(buf.Count(), pos+n)
		buf.Delete(pos, i)
		expected = expected[0:pos] + expected[i:]
	}

	assert.Equal(t, expected, buf.All())
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}
