package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/segbuf/textbuf"
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

	buf.InsertIndex(buf.Count(), "Lorem")
	buf.InsertIndex(buf.Count(), " ipsum")
	buf.InsertIndex(buf.Count(), " dolor")
	buf.InsertIndex(buf.Count(), " sit")
	buf.InsertIndex(buf.Count(), " amet,")
	buf.InsertIndex(buf.Count(), " consectetur")
	buf.InsertIndex(buf.Count(), " adipiscing")
	buf.InsertIndex(buf.Count(), " elit,")
	buf.InsertIndex(buf.Count(), " sed")
	buf.InsertIndex(buf.Count(), " do")
	buf.InsertIndex(buf.Count(), " eiusmod")
	buf.InsertIndex(buf.Count(), " tempor")
	buf.InsertIndex(buf.Count(), " incididunt")
	buf.InsertIndex(buf.Count(), " ut")
	buf.InsertIndex(buf.Count(), " labore")
	buf.InsertIndex(buf.Count(), " et")
	buf.InsertIndex(buf.Count(), " dolore")
	buf.InsertIndex(buf.Count(), " magna")
	buf.InsertIndex(buf.Count(), " aliqua.")

	return buf
}

func createTextBufReversed() textbuf.TextBuf {
	buf := textbuf.New()

	buf.InsertIndex(0, " aliqua.")
	buf.InsertIndex(0, " magna")
	buf.InsertIndex(0, " dolore")
	buf.InsertIndex(0, " et")
	buf.InsertIndex(0, " labore")
	buf.InsertIndex(0, " ut")
	buf.InsertIndex(0, " incididunt")
	buf.InsertIndex(0, " tempor")
	buf.InsertIndex(0, " eiusmod")
	buf.InsertIndex(0, " do")
	buf.InsertIndex(0, " sed")
	buf.InsertIndex(0, " elit,")
	buf.InsertIndex(0, " adipiscing")
	buf.InsertIndex(0, " consectetur")
	buf.InsertIndex(0, " amet,")
	buf.InsertIndex(0, " sit")
	buf.InsertIndex(0, " dolor")
	buf.InsertIndex(0, " ipsum")
	buf.InsertIndex(0, "Lorem")

	return buf
}

const text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

func testDeleteHead(t *testing.T, buf textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected,
			iterToStr(buf.ReadIndex(0)))
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		i := min(buf.Count(), n)
		buf.DeleteIndexRange(0, i)
		expected = expected[i:]
	}

	assert.Equal(t, expected,
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}

func testDeleteTail(t *testing.T, buf textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected,
			iterToStr(buf.ReadIndex(0)))
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		i := max(buf.Count()-n, 0)
		buf.DeleteIndexRange(i, buf.Count())
		expected = expected[0:i]
	}

	assert.Equal(t, expected,
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}

func testDeleteMiddle(t *testing.T, buf textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected,
			iterToStr(buf.ReadIndex(0)))
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		pos := buf.Count() / 2
		i := min(buf.Count(), pos+n)
		buf.DeleteIndexRange(pos, i)
		expected = expected[0:pos] + expected[i:]
	}

	assert.Equal(t, expected,
		iterToStr(buf.ReadIndex(0)))
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}
