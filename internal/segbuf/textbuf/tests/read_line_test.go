package textbuf_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/segbuf/textbuf"
)

func TestReadEmptyLine(t *testing.T) {
	buf := textbuf.New()

	assert.Equal(t, 0, buf.LineCount())
	assert.Equal(t, "",
		iterToStr(buf.ReadPosRange(0, 0, 1, 0)))

	buf.Validate()
}

func Test1Line(t *testing.T) {
	buf := textbuf.New()
	buf.Append("0")

	assert.Equal(t, 1, buf.LineCount())
	assert.Equal(t, "0",
		iterToStr(buf.ReadPosRange(0, 0, 1, 0)))

	buf.Validate()
}

func Test2Lines(t *testing.T) {
	buf := textbuf.New()
	buf.Append("0\n")

	assert.Equal(t, 2, buf.LineCount())
	assert.Equal(t, "0\n",
		iterToStr(buf.ReadPosRange(0, 0, 1, 0)))
	assert.Equal(t, "",
		iterToStr(buf.ReadPosRange(1, 0, 2, 0)))

	buf.Validate()
}

func Test3Lines(t *testing.T) {
	buf := textbuf.New()
	buf.Append("0\n1\n")

	assert.Equal(t, 3, buf.LineCount())
	assert.Equal(t, "0\n",
		iterToStr(buf.ReadPosRange(0, 0, 1, 0)))
	assert.Equal(t, "1\n",
		iterToStr(buf.ReadPosRange(1, 0, 2, 0)))
	assert.Equal(t, "",
		iterToStr(buf.ReadPosRange(2, 0, 3, 0)))

	buf.Validate()
}

func TestReadLineAtValidIndex(t *testing.T) {
	buf := textbuf.New()

	buf.InsertIndex(0, "Lorem\naliqua.")
	buf.InsertIndex(6, "ipsum\nmagna\n")
	buf.InsertIndex(12, "dolor\ndolore\n")
	buf.InsertIndex(18, "sit\net\n")
	buf.InsertIndex(22, "amet,\nlabore\n")
	buf.InsertIndex(28, "consectetur\nut\n")
	buf.InsertIndex(40, "adipiscing\nincididunt\n")
	buf.InsertIndex(51, "elit,\ntempor\n")
	buf.InsertIndex(57, "sed\neiusmod\n")
	buf.InsertIndex(61, "do\n")

	assert.Equal(t, "Lorem\n",
		iterToStr(buf.ReadPosRange(0, 0, 1, 0)))
	assert.Equal(t, "ipsum\n",
		iterToStr(buf.ReadPosRange(1, 0, 2, 0)))
	assert.Equal(t, "dolor\n",
		iterToStr(buf.ReadPosRange(2, 0, 3, 0)))
	assert.Equal(t, "sit\n",
		iterToStr(buf.ReadPosRange(3, 0, 4, 0)))
	assert.Equal(t, "amet,\n",
		iterToStr(buf.ReadPosRange(4, 0, 5, 0)))
	assert.Equal(t, "consectetur\n",
		iterToStr(buf.ReadPosRange(5, 0, 6, 0)))
	assert.Equal(t, "adipiscing\n",
		iterToStr(buf.ReadPosRange(6, 0, 7, 0)))
	assert.Equal(t, "elit,\n",
		iterToStr(buf.ReadPosRange(7, 0, 8, 0)))
	assert.Equal(t, "sed\n",
		iterToStr(buf.ReadPosRange(8, 0, 9, 0)))
	assert.Equal(t, "do\n",
		iterToStr(buf.ReadPosRange(9, 0, 10, 0)))
	assert.Equal(t, "eiusmod\n",
		iterToStr(buf.ReadPosRange(10, 0, 11, 0)))
	assert.Equal(t, "tempor\n",
		iterToStr(buf.ReadPosRange(11, 0, 12, 0)))
	assert.Equal(t, "incididunt\n",
		iterToStr(buf.ReadPosRange(12, 0, 13, 0)))
	assert.Equal(t, "ut\n",
		iterToStr(buf.ReadPosRange(13, 0, 14, 0)))
	assert.Equal(t, "labore\n",
		iterToStr(buf.ReadPosRange(14, 0, 15, 0)))
	assert.Equal(t, "et\n",
		iterToStr(buf.ReadPosRange(15, 0, 16, 0)))
	assert.Equal(t, "dolore\n",
		iterToStr(buf.ReadPosRange(16, 0, 17, 0)))
	assert.Equal(t, "magna\n",
		iterToStr(buf.ReadPosRange(17, 0, 18, 0)))
	assert.Equal(t, "aliqua.",
		iterToStr(buf.ReadPosRange(18, 0, 19, 0)))

	buf.Validate()
}

func TestReadLineAtIndexGTELineCount(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "amet",
		iterToStr(buf.ReadPosRange(4, 0, 5, 0)))
	assert.Equal(t, "",
		iterToStr(buf.ReadPosRange(5, 0, 6, 0)))
	assert.Equal(t, "",
		iterToStr(buf.ReadPosRange(6, 0, 7, 0)))

	buf.Validate()
}

func TestReadLineAtIndexLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "Lorem\n",
		iterToStr(buf.ReadPosRange(0, 0, 1, 0)))
	assert.Equal(t, "amet",
		iterToStr(buf.ReadPosRange(buf.LineCount()-1, 0, buf.LineCount(), 0)))
	assert.Equal(t, "sit\n",
		iterToStr(buf.ReadPosRange(buf.LineCount()-2, 0, buf.LineCount()-1, 0)))

	buf.Validate()
}

func TestInsertAddsLines(t *testing.T) {
	buf := textbuf.New()

	for i := 0; i < 10; i += 1 {
		buf.InsertIndex(buf.Count(), fmt.Sprintf("%d\n", i))

		assert.Equal(t, i+2, buf.LineCount())
		assert.Equal(t, fmt.Sprintf("%d\n", i),
			iterToStr(buf.ReadPosRange(i, 0, i+1, 0)))
		buf.Validate()
	}

	assert.Equal(t, 11, buf.LineCount())
	assert.Equal(t, "",
		iterToStr(buf.ReadPosRange(11, 0, 12, 0)))
	buf.Validate()
}
