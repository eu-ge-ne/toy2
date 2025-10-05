package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/segbuf/textbuf"
)

func TestLineCount0Newlines(t *testing.T) {
	buf1 := textbuf.New()
	buf1.Append("A")
	buf2 := textbuf.New()
	buf2.Append("😄")
	buf3 := textbuf.New()
	buf3.Append("🤦🏼‍♂️")

	assert.Equal(t, 1, buf1.LineCount())
	assert.Equal(t, 1, buf2.LineCount())
	assert.Equal(t, 1, buf3.LineCount())
}

func TestLineCountLF(t *testing.T) {
	buf1 := textbuf.New()
	buf1.Append("A\nA")
	buf2 := textbuf.New()
	buf2.Append("😄\n😄")
	buf3 := textbuf.New()
	buf3.Append("🤦🏼‍♂️\n🤦🏼‍♂️")

	assert.Equal(t, 2, buf1.LineCount())
	assert.Equal(t, 2, buf2.LineCount())
	assert.Equal(t, 2, buf3.LineCount())
}

func TestLineCountCRLF(t *testing.T) {
	buf1 := textbuf.New()
	buf1.Append("A\r\nA")
	buf2 := textbuf.New()
	buf2.Append("😄\r\n😄")
	buf3 := textbuf.New()
	buf3.Append("🤦🏼‍♂️\r\n🤦🏼‍♂️")

	assert.Equal(t, 2, buf1.LineCount())
	assert.Equal(t, 2, buf2.LineCount())
	assert.Equal(t, 2, buf3.LineCount())
}
