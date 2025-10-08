package textbuf_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/std"
	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestReadEmpty(t *testing.T) {
	buf := textbuf.New()

	assert.Equal(t, "", std.IterToStr(buf.Slice(0, math.MaxInt)))
	buf.Validate()
}

func TestRead(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem ipsum dolor")

	assert.Equal(t, "ipsum ", std.IterToStr(buf.Slice(6, 12)))
	buf.Validate()
}

func TestReadAtStartGTECount(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem")

	assert.Equal(t, "m", std.IterToStr(buf.Slice(4, math.MaxInt)))
	assert.Equal(t, "", std.IterToStr(buf.Slice(5, math.MaxInt)))
	assert.Equal(t, "", std.IterToStr(buf.Slice(6, math.MaxInt)))

	buf.Validate()
}

func TestReadAtStartLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem")

	assert.Equal(t, "Lorem", std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, "m", std.IterToStr(buf.Slice(buf.Count()-1, math.MaxInt)))
	assert.Equal(t, "em", std.IterToStr(buf.Slice(buf.Count()-2, math.MaxInt)))

	buf.Validate()
}

func TestReadEmptyLine(t *testing.T) {
	buf := textbuf.New()

	assert.Equal(t, 0, buf.LineCount())
	assert.Equal(t, "", std.IterToStr(buf.Read(0, 0, 1, 0)))

	buf.Validate()
}

func Test1Line(t *testing.T) {
	buf := textbuf.New()
	buf.Append("0")

	assert.Equal(t, 1, buf.LineCount())
	assert.Equal(t, "0", std.IterToStr(buf.Read(0, 0, 0, 1)))

	buf.Validate()
}

func Test2Lines(t *testing.T) {
	buf := textbuf.New()
	buf.Append("0\n")

	assert.Equal(t, 2, buf.LineCount())
	assert.Equal(t, "0\n", std.IterToStr(buf.Read(0, 0, 1, 0)))
	assert.Equal(t, "", std.IterToStr(buf.Read(1, 0, 2, 0)))

	buf.Validate()
}

func Test3Lines(t *testing.T) {
	buf := textbuf.New()
	buf.Append("0\n1\n")

	assert.Equal(t, 3, buf.LineCount())
	assert.Equal(t, "0\n", std.IterToStr(buf.Read(0, 0, 1, 0)))
	assert.Equal(t, "1\n", std.IterToStr(buf.Read(1, 0, 2, 0)))
	assert.Equal(t, "", std.IterToStr(buf.Read(2, 0, 3, 0)))

	buf.Validate()
}

func TestReadLineAtValidIndex(t *testing.T) {
	buf := textbuf.New()

	buf.Insert(0, 0, "Lorem\naliqua.")
	buf.Insert(1, 0, "ipsum\nmagna\n")
	buf.Insert(2, 0, "dolor\ndolore\n")
	buf.Insert(3, 0, "sit\net\n")
	buf.Insert(4, 0, "amet,\nlabore\n")
	buf.Insert(5, 0, "consectetur\nut\n")
	buf.Insert(6, 0, "adipiscing\nincididunt\n")
	buf.Insert(7, 0, "elit,\ntempor\n")
	buf.Insert(8, 0, "sed\neiusmod\n")
	buf.Insert(9, 0, "do\n")

	assert.Equal(t, "Lorem\n", std.IterToStr(buf.Read(0, 0, 1, 0)))
	assert.Equal(t, "ipsum\n", std.IterToStr(buf.Read(1, 0, 2, 0)))
	assert.Equal(t, "dolor\n", std.IterToStr(buf.Read(2, 0, 3, 0)))
	assert.Equal(t, "sit\n", std.IterToStr(buf.Read(3, 0, 4, 0)))
	assert.Equal(t, "amet,\n", std.IterToStr(buf.Read(4, 0, 5, 0)))
	assert.Equal(t, "consectetur\n", std.IterToStr(buf.Read(5, 0, 6, 0)))
	assert.Equal(t, "adipiscing\n", std.IterToStr(buf.Read(6, 0, 7, 0)))
	assert.Equal(t, "elit,\n", std.IterToStr(buf.Read(7, 0, 8, 0)))
	assert.Equal(t, "sed\n", std.IterToStr(buf.Read(8, 0, 9, 0)))
	assert.Equal(t, "do\n", std.IterToStr(buf.Read(9, 0, 10, 0)))
	assert.Equal(t, "eiusmod\n", std.IterToStr(buf.Read(10, 0, 11, 0)))
	assert.Equal(t, "tempor\n", std.IterToStr(buf.Read(11, 0, 12, 0)))
	assert.Equal(t, "incididunt\n", std.IterToStr(buf.Read(12, 0, 13, 0)))
	assert.Equal(t, "ut\n", std.IterToStr(buf.Read(13, 0, 14, 0)))
	assert.Equal(t, "labore\n", std.IterToStr(buf.Read(14, 0, 15, 0)))
	assert.Equal(t, "et\n", std.IterToStr(buf.Read(15, 0, 16, 0)))
	assert.Equal(t, "dolore\n", std.IterToStr(buf.Read(16, 0, 17, 0)))
	assert.Equal(t, "magna\n", std.IterToStr(buf.Read(17, 0, 18, 0)))
	assert.Equal(t, "aliqua.", std.IterToStr(buf.Read(18, 0, 18, 7)))

	buf.Validate()
}

func TestReadLineAtIndexGTELineCount(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "amet", std.IterToStr(buf.Read(4, 0, 4, 4)))
	assert.Equal(t, "", std.IterToStr(buf.Read(5, 0, 6, 0)))
	assert.Equal(t, "", std.IterToStr(buf.Read(6, 0, 7, 0)))

	buf.Validate()
}

func TestReadLineAtIndexLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "Lorem\n", std.IterToStr(buf.Read(0, 0, 1, 0)))
	assert.Equal(t, "amet", std.IterToStr(buf.Read(buf.LineCount()-1, 0, buf.LineCount()-1, 4)))
	assert.Equal(t, "sit\n", std.IterToStr(buf.Read(buf.LineCount()-2, 0, buf.LineCount()-1, 0)))

	buf.Validate()
}

func TestInsertAddsLines(t *testing.T) {
	buf := textbuf.New()

	for i := 0; i < 10; i += 1 {
		buf.Append(fmt.Sprintf("%d\n", i))

		assert.Equal(t, i+2, buf.LineCount())
		assert.Equal(t, fmt.Sprintf("%d\n", i), std.IterToStr(buf.Read(i, 0, i+1, 0)))
		buf.Validate()
	}

	assert.Equal(t, 11, buf.LineCount())
	assert.Equal(t, "", std.IterToStr(buf.Read(11, 0, 12, 0)))
	buf.Validate()
}

func TestLineAtValidIndex(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "Lorem\nipsum\ndolor\nsit\namet",
		std.IterToStr(buf.Read(0, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "ipsum\ndolor\nsit\namet",
		std.IterToStr(buf.Read(1, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "dolor\nsit\namet",
		std.IterToStr(buf.Read(2, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "sit\namet",
		std.IterToStr(buf.Read(3, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "amet",
		std.IterToStr(buf.Read(4, 0, math.MaxInt, math.MaxInt)))

	buf.Validate()
}

func TestLineAtIndexGTELineCount(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "amet",
		std.IterToStr(buf.Read(4, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "",
		std.IterToStr(buf.Read(5, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "",
		std.IterToStr(buf.Read(6, 0, math.MaxInt, math.MaxInt)))

	buf.Validate()
}

func TestLineAtIndexLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "Lorem\nipsum\ndolor\nsit\namet",
		std.IterToStr(buf.Read(0, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "amet",
		std.IterToStr(buf.Read(buf.LineCount()-1, 0, math.MaxInt, math.MaxInt)))
	assert.Equal(t, "sit\namet",
		std.IterToStr(buf.Read(buf.LineCount()-2, 0, math.MaxInt, math.MaxInt)))

	buf.Validate()
}
