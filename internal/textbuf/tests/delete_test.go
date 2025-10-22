package textbuf_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/std"
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

	buf.Append("Lorem")
	buf.Append(" ipsum")
	buf.Append(" dolor")
	buf.Append(" sit")
	buf.Append(" amet,")
	buf.Append(" consectetur")
	buf.Append(" adipiscing")
	buf.Append(" elit,")
	buf.Append(" sed")
	buf.Append(" do")
	buf.Append(" eiusmod")
	buf.Append(" tempor")
	buf.Append(" incididunt")
	buf.Append(" ut")
	buf.Append(" labore")
	buf.Append(" et")
	buf.Append(" dolore")
	buf.Append(" magna")
	buf.Append(" aliqua.")

	return buf
}

func createTextBufReversed() *textbuf.TextBuf {
	buf := textbuf.New()

	buf.Insert(0, 0, " aliqua.")
	buf.Insert(0, 0, " magna")
	buf.Insert(0, 0, " dolore")
	buf.Insert(0, 0, " et")
	buf.Insert(0, 0, " labore")
	buf.Insert(0, 0, " ut")
	buf.Insert(0, 0, " incididunt")
	buf.Insert(0, 0, " tempor")
	buf.Insert(0, 0, " eiusmod")
	buf.Insert(0, 0, " do")
	buf.Insert(0, 0, " sed")
	buf.Insert(0, 0, " elit,")
	buf.Insert(0, 0, " adipiscing")
	buf.Insert(0, 0, " consectetur")
	buf.Insert(0, 0, " amet,")
	buf.Insert(0, 0, " sit")
	buf.Insert(0, 0, " dolor")
	buf.Insert(0, 0, " ipsum")
	buf.Insert(0, 0, "Lorem")

	return buf
}

const text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

func testDeleteHead(t *testing.T, buf *textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected, std.IterToStr(buf.Read(0, math.MaxInt)))
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		i := min(buf.Count(), n)
		buf.Delete(0, i)
		expected = expected[i:]
	}

	assert.Equal(t, expected, std.IterToStr(buf.Read(0, math.MaxInt)))
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}

func testDeleteTail(t *testing.T, buf *textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected, std.IterToStr(buf.Read(0, math.MaxInt)))
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		i := max(buf.Count()-n, 0)
		buf.Delete(i, buf.Count())
		expected = expected[0:i]
	}

	assert.Equal(t, expected, std.IterToStr(buf.Read(0, math.MaxInt)))
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}

func testDeleteMiddle(t *testing.T, buf *textbuf.TextBuf, n int) {
	expected := text

	for len(expected) > 0 {
		assert.Equal(t, expected, std.IterToStr(buf.Read(0, math.MaxInt)))
		assert.Equal(t, len(expected), buf.Count())
		buf.Validate()

		pos := buf.Count() / 2
		i := min(buf.Count(), pos+n)
		buf.Delete(pos, i)
		expected = expected[0:pos] + expected[i:]
	}

	assert.Equal(t, expected, std.IterToStr(buf.Read(0, math.MaxInt)))
	assert.Equal(t, len(expected), buf.Count())
	buf.Validate()
}
