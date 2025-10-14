package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestLineCount0Newlines(t *testing.T) {
	buf1 := textbuf.New()
	buf1.AppendString("A")
	buf2 := textbuf.New()
	buf2.AppendString("😄")
	buf3 := textbuf.New()
	buf3.AppendString("🤦🏼‍♂️")

	assert.Equal(t, 1, buf1.LineCount())
	assert.Equal(t, 1, buf2.LineCount())
	assert.Equal(t, 1, buf3.LineCount())
}

func TestLineCountLF(t *testing.T) {
	buf1 := textbuf.New()
	buf1.AppendString("A\nA")
	buf2 := textbuf.New()
	buf2.AppendString("😄\n😄")
	buf3 := textbuf.New()
	buf3.AppendString("🤦🏼‍♂️\n🤦🏼‍♂️")

	assert.Equal(t, 2, buf1.LineCount())
	assert.Equal(t, 2, buf2.LineCount())
	assert.Equal(t, 2, buf3.LineCount())
}

func TestLineCountCRLF(t *testing.T) {
	buf1 := textbuf.New()
	buf1.AppendString("A\r\nA")
	buf2 := textbuf.New()
	buf2.AppendString("😄\r\n😄")
	buf3 := textbuf.New()
	buf3.AppendString("🤦🏼‍♂️\r\n🤦🏼‍♂️")

	assert.Equal(t, 2, buf1.LineCount())
	assert.Equal(t, 2, buf2.LineCount())
	assert.Equal(t, 2, buf3.LineCount())
}
