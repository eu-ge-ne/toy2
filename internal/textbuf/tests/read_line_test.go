package textbuf_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestReadEmptyLine(t *testing.T) {
	buf := textbuf.New()

	assert.Equal(t, 0, buf.LineCount())
	assert.Equal(t, "", buf.Read2(0, 0, 1, 0))

	buf.Validate()
}

func Test1Line(t *testing.T) {
	buf := textbuf.New()
	buf.Append([]byte("0"))

	assert.Equal(t, 1, buf.LineCount())
	assert.Equal(t, "0", buf.Read2(0, 0, 1, 0))

	buf.Validate()
}

func Test2Lines(t *testing.T) {
	buf := textbuf.New()
	buf.Append([]byte("0\n"))

	assert.Equal(t, 2, buf.LineCount())
	assert.Equal(t, "0\n", buf.Read2(0, 0, 1, 0))
	assert.Equal(t, "", buf.Read2(1, 0, 2, 0))

	buf.Validate()
}

func Test3Lines(t *testing.T) {
	buf := textbuf.New()
	buf.Append([]byte("0\n1\n"))

	assert.Equal(t, 3, buf.LineCount())
	assert.Equal(t, "0\n", buf.Read2(0, 0, 1, 0))
	assert.Equal(t, "1\n", buf.Read2(1, 0, 2, 0))
	assert.Equal(t, "", buf.Read2(2, 0, 3, 0))

	buf.Validate()
}

func TestReadLineAtValidIndex(t *testing.T) {
	buf := textbuf.New()

	buf.InsertString(0, "Lorem\naliqua.")
	buf.InsertString(6, "ipsum\nmagna\n")
	buf.InsertString(12, "dolor\ndolore\n")
	buf.InsertString(18, "sit\net\n")
	buf.InsertString(22, "amet,\nlabore\n")
	buf.InsertString(28, "consectetur\nut\n")
	buf.InsertString(40, "adipiscing\nincididunt\n")
	buf.InsertString(51, "elit,\ntempor\n")
	buf.InsertString(57, "sed\neiusmod\n")
	buf.InsertString(61, "do\n")

	assert.Equal(t, "Lorem\n", buf.Read2(0, 0, 1, 0))
	assert.Equal(t, "ipsum\n", buf.Read2(1, 0, 2, 0))
	assert.Equal(t, "dolor\n", buf.Read2(2, 0, 3, 0))
	assert.Equal(t, "sit\n", buf.Read2(3, 0, 4, 0))
	assert.Equal(t, "amet,\n", buf.Read2(4, 0, 5, 0))
	assert.Equal(t, "consectetur\n", buf.Read2(5, 0, 6, 0))
	assert.Equal(t, "adipiscing\n", buf.Read2(6, 0, 7, 0))
	assert.Equal(t, "elit,\n", buf.Read2(7, 0, 8, 0))
	assert.Equal(t, "sed\n", buf.Read2(8, 0, 9, 0))
	assert.Equal(t, "do\n", buf.Read2(9, 0, 10, 0))
	assert.Equal(t, "eiusmod\n", buf.Read2(10, 0, 11, 0))
	assert.Equal(t, "tempor\n", buf.Read2(11, 0, 12, 0))
	assert.Equal(t, "incididunt\n", buf.Read2(12, 0, 13, 0))
	assert.Equal(t, "ut\n", buf.Read2(13, 0, 14, 0))
	assert.Equal(t, "labore\n", buf.Read2(14, 0, 15, 0))
	assert.Equal(t, "et\n", buf.Read2(15, 0, 16, 0))
	assert.Equal(t, "dolore\n", buf.Read2(16, 0, 17, 0))
	assert.Equal(t, "magna\n", buf.Read2(17, 0, 18, 0))
	assert.Equal(t, "aliqua.", buf.Read2(18, 0, 19, 0))

	buf.Validate()
}

func TestReadLineAtIndexGTELineCount(t *testing.T) {
	buf := textbuf.New()
	buf.Append([]byte("Lorem\nipsum\ndolor\nsit\namet"))

	assert.Equal(t, "amet", buf.Read2(4, 0, 5, 0))
	assert.Equal(t, "", buf.Read2(5, 0, 6, 0))
	assert.Equal(t, "", buf.Read2(6, 0, 7, 0))

	buf.Validate()
}

func TestReadLineAtIndexLT0(t *testing.T) {
	buf := textbuf.New()
	buf.Append([]byte("Lorem\nipsum\ndolor\nsit\namet"))

	assert.Equal(t, "Lorem\n", buf.Read2(0, 0, 1, 0))
	assert.Equal(t, "amet", buf.Read2(buf.LineCount()-1, 0, buf.LineCount(), 0))
	assert.Equal(t, "sit\n", buf.Read2(buf.LineCount()-2, 0, buf.LineCount()-1, 0))

	buf.Validate()
}

func TestInsertAddsLines(t *testing.T) {
	buf := textbuf.New()

	for i := 0; i < 10; i += 1 {
		buf.InsertString(buf.Count(), fmt.Sprintf("%d\n", i))

		assert.Equal(t, i+2, buf.LineCount())
		assert.Equal(t, fmt.Sprintf("%d\n", i), buf.Read2(i, 0, i+1, 0))
		buf.Validate()
	}

	assert.Equal(t, 11, buf.LineCount())
	assert.Equal(t, "", buf.Read2(11, 0, 12, 0))
	buf.Validate()
}
