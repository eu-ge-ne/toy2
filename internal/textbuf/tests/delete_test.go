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

	buf.Insert(buf.Count(), []byte("Lorem"))
	buf.Insert(buf.Count(), []byte(" ipsum"))
	buf.Insert(buf.Count(), []byte(" dolor"))
	buf.Insert(buf.Count(), []byte(" sit"))
	buf.Insert(buf.Count(), []byte(" amet,"))
	buf.Insert(buf.Count(), []byte(" consectetur"))
	buf.Insert(buf.Count(), []byte(" adipiscing"))
	buf.Insert(buf.Count(), []byte(" elit,"))
	buf.Insert(buf.Count(), []byte(" sed"))
	buf.Insert(buf.Count(), []byte(" do"))
	buf.Insert(buf.Count(), []byte(" eiusmod"))
	buf.Insert(buf.Count(), []byte(" tempor"))
	buf.Insert(buf.Count(), []byte(" incididunt"))
	buf.Insert(buf.Count(), []byte(" ut"))
	buf.Insert(buf.Count(), []byte(" labore"))
	buf.Insert(buf.Count(), []byte(" et"))
	buf.Insert(buf.Count(), []byte(" dolore"))
	buf.Insert(buf.Count(), []byte(" magna"))
	buf.Insert(buf.Count(), []byte(" aliqua."))

	return buf
}

func createTextBufReversed() *textbuf.TextBuf {
	buf := textbuf.New()

	buf.Insert(0, []byte(" aliqua."))
	buf.Insert(0, []byte(" magna"))
	buf.Insert(0, []byte(" dolore"))
	buf.Insert(0, []byte(" et"))
	buf.Insert(0, []byte(" labore"))
	buf.Insert(0, []byte(" ut"))
	buf.Insert(0, []byte(" incididunt"))
	buf.Insert(0, []byte(" tempor"))
	buf.Insert(0, []byte(" eiusmod"))
	buf.Insert(0, []byte(" do"))
	buf.Insert(0, []byte(" sed"))
	buf.Insert(0, []byte(" elit,"))
	buf.Insert(0, []byte(" adipiscing"))
	buf.Insert(0, []byte(" consectetur"))
	buf.Insert(0, []byte(" amet,"))
	buf.Insert(0, []byte(" sit"))
	buf.Insert(0, []byte(" dolor"))
	buf.Insert(0, []byte(" ipsum"))
	buf.Insert(0, []byte("Lorem"))

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
