package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestLineCount0Newlines(t *testing.T) {
	buf1 := textbuf.New("A")
	buf2 := textbuf.New("😄")
	buf3 := textbuf.New("🤦🏼‍♂️")

	assert.Equal(t, 1, buf1.LineCount())
	assert.Equal(t, 1, buf2.LineCount())
	assert.Equal(t, 1, buf3.LineCount())
}

func TestLineCountLF(t *testing.T) {
	buf1 := textbuf.New("A\nA")
	buf2 := textbuf.New("😄\n😄")
	buf3 := textbuf.New("🤦🏼‍♂️\n🤦🏼‍♂️")

	assert.Equal(t, 2, buf1.LineCount())
	assert.Equal(t, 2, buf2.LineCount())
	assert.Equal(t, 2, buf3.LineCount())
}

func TestLineCountCRLF(t *testing.T) {
	buf1 := textbuf.New("A\r\nA")
	buf2 := textbuf.New("😄\r\n😄")
	buf3 := textbuf.New("🤦🏼‍♂️\r\n🤦🏼‍♂️")

	assert.Equal(t, 2, buf1.LineCount())
	assert.Equal(t, 2, buf2.LineCount())
	assert.Equal(t, 2, buf3.LineCount())
}
