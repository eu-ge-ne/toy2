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
		iterToStr(buf.Read2Range(0, 0, 1, 0)))

	buf.Validate()
}

func Test1Line(t *testing.T) {
	buf := textbuf.New()
	buf.Append("0")

	assert.Equal(t, 1, buf.LineCount())
	assert.Equal(t, "0",
		iterToStr(buf.Read2Range(0, 0, 1, 0)))

	buf.Validate()
}

func Test2Lines(t *testing.T) {
	buf := textbuf.New()
	buf.Append("0\n")

	assert.Equal(t, 2, buf.LineCount())
	assert.Equal(t, "0\n",
		iterToStr(buf.Read2Range(0, 0, 1, 0)))
	assert.Equal(t, "",
		iterToStr(buf.Read2Range(1, 0, 2, 0)))

	buf.Validate()
}

func Test3Lines(t *testing.T) {
	buf := textbuf.New()
	buf.Append("0\n1\n")

	assert.Equal(t, 3, buf.LineCount())
	assert.Equal(t, "0\n",
		iterToStr(buf.Read2Range(0, 0, 1, 0)))
	assert.Equal(t, "1\n",
		iterToStr(buf.Read2Range(1, 0, 2, 0)))
	assert.Equal(t, "",
		iterToStr(buf.Read2Range(2, 0, 3, 0)))

	buf.Validate()
}

func TestReadLineAtValidIndex(t *testing.T) {
	buf := textbuf.New()

	buf.Insert(0, "Lorem\naliqua.")
	buf.Insert(6, "ipsum\nmagna\n")
	buf.Insert(12, "dolor\ndolore\n")
	buf.Insert(18, "sit\net\n")
	buf.Insert(22, "amet,\nlabore\n")
	buf.Insert(28, "consectetur\nut\n")
	buf.Insert(40, "adipiscing\nincididunt\n")
	buf.Insert(51, "elit,\ntempor\n")
	buf.Insert(57, "sed\neiusmod\n")
	buf.Insert(61, "do\n")

	assert.Equal(t, "Lorem\n",
		iterToStr(buf.Read2Range(0, 0, 1, 0)))
	assert.Equal(t, "ipsum\n",
		iterToStr(buf.Read2Range(1, 0, 2, 0)))
	assert.Equal(t, "dolor\n",
		iterToStr(buf.Read2Range(2, 0, 3, 0)))
	assert.Equal(t, "sit\n",
		iterToStr(buf.Read2Range(3, 0, 4, 0)))
	assert.Equal(t, "amet,\n",
		iterToStr(buf.Read2Range(4, 0, 5, 0)))
	assert.Equal(t, "consectetur\n",
		iterToStr(buf.Read2Range(5, 0, 6, 0)))
	assert.Equal(t, "adipiscing\n",
		iterToStr(buf.Read2Range(6, 0, 7, 0)))
	assert.Equal(t, "elit,\n",
		iterToStr(buf.Read2Range(7, 0, 8, 0)))
	assert.Equal(t, "sed\n",
		iterToStr(buf.Read2Range(8, 0, 9, 0)))
	assert.Equal(t, "do\n",
		iterToStr(buf.Read2Range(9, 0, 10, 0)))
	assert.Equal(t, "eiusmod\n",
		iterToStr(buf.Read2Range(10, 0, 11, 0)))
	assert.Equal(t, "tempor\n",
		iterToStr(buf.Read2Range(11, 0, 12, 0)))
	assert.Equal(t, "incididunt\n",
		iterToStr(buf.Read2Range(12, 0, 13, 0)))
	assert.Equal(t, "ut\n",
		iterToStr(buf.Read2Range(13, 0, 14, 0)))
	assert.Equal(t, "labore\n",
		iterToStr(buf.Read2Range(14, 0, 15, 0)))
	assert.Equal(t, "et\n",
		iterToStr(buf.Read2Range(15, 0, 16, 0)))
	assert.Equal(t, "dolore\n",
		iterToStr(buf.Read2Range(16, 0, 17, 0)))
	assert.Equal(t, "magna\n",
		iterToStr(buf.Read2Range(17, 0, 18, 0)))
	assert.Equal(t, "aliqua.",
		iterToStr(buf.Read2Range(18, 0, 19, 0)))

	buf.Validate()
}

func TestReadLineAtIndexGTELineCount(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "amet",
		iterToStr(buf.Read2Range(4, 0, 5, 0)))
	assert.Equal(t, "",
		iterToStr(buf.Read2Range(5, 0, 6, 0)))
	assert.Equal(t, "",
		iterToStr(buf.Read2Range(6, 0, 7, 0)))

	buf.Validate()
}

func TestReadLineAtIndexLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append("Lorem\nipsum\ndolor\nsit\namet")

	assert.Equal(t, "Lorem\n",
		iterToStr(buf.Read2Range(0, 0, 1, 0)))
	assert.Equal(t, "amet",
		iterToStr(buf.Read2Range(buf.LineCount()-1, 0, buf.LineCount(), 0)))
	assert.Equal(t, "sit\n",
		iterToStr(buf.Read2Range(buf.LineCount()-2, 0, buf.LineCount()-1, 0)))

	buf.Validate()
}

func TestInsertAddsLines(t *testing.T) {
	buf := textbuf.New()

	for i := 0; i < 10; i += 1 {
		buf.Insert(buf.Count(), fmt.Sprintf("%d\n", i))

		assert.Equal(t, i+2, buf.LineCount())
		assert.Equal(t, fmt.Sprintf("%d\n", i),
			iterToStr(buf.Read2Range(i, 0, i+1, 0)))
		buf.Validate()
	}

	assert.Equal(t, 11, buf.LineCount())
	assert.Equal(t, "",
		iterToStr(buf.Read2Range(11, 0, 12, 0)))
	buf.Validate()
}
