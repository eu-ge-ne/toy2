package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBuffer(t *testing.T) {
	buf := Create("Lorem ipsum")

	assert.Equal(t, " ", string(buf.Read(5, 6)))
}

func Test0Newlines(t *testing.T) {
	buf := Create("Lorem ipsum")

	assert.Equal(t, 0, len(buf.Eols))
	assert.Equal(t, []Eol(nil), buf.Eols)
}

func TestLF(t *testing.T) {
	buf := Create("Lorem \nipsum \n")

	assert.Equal(t, 2, len(buf.Eols))
	assert.Equal(t, []Eol{{6, 7}, {13, 14}}, buf.Eols)
}

func TestCRLF(t *testing.T) {
	buf := Create("Lorem \r\nipsum \r\n")

	assert.Equal(t, 2, len(buf.Eols))
	assert.Equal(t, []Eol{{6, 8}, {14, 16}}, buf.Eols)
}

func TestLFCRLF(t *testing.T) {
	buf := Create("Lorem \nipsum \r\n")

	assert.Equal(t, 2, len(buf.Eols))
	assert.Equal(t, []Eol{{6, 7}, {13, 15}}, buf.Eols)
}

func TestFindEolIndex(t *testing.T) {
	buf := Create("AA\r\nBB\nCC")

	assert.Equal(t, 2, len(buf.Eols))

	assert.Equal(t, 0, buf.FindEolIndex(0, 0))
	assert.Equal(t, 0, buf.FindEolIndex(1, 0))

	assert.Equal(t, 0, buf.FindEolIndex(2, 0))
	assert.Panics(t, func() { buf.FindEolIndex(3, 0) })

	assert.Equal(t, 1, buf.FindEolIndex(4, 0))
	assert.Equal(t, 1, buf.FindEolIndex(5, 0))

	assert.Equal(t, 1, buf.FindEolIndex(6, 0))

	assert.Equal(t, 2, buf.FindEolIndex(7, 0))
	assert.Equal(t, 2, buf.FindEolIndex(8, 0))

	buf = Create("1\n2\n3\n4\n5")
	//                  01 23 45 67 8
	//                   0  1  2  3

	assert.Equal(t, 4, len(buf.Eols))

	assert.Equal(t, 0, buf.FindEolIndex(0, 0))
	assert.Equal(t, 0, buf.FindEolIndex(1, 0))
	assert.Equal(t, 1, buf.FindEolIndex(2, 0))
	assert.Equal(t, 1, buf.FindEolIndex(3, 0))
	assert.Equal(t, 2, buf.FindEolIndex(4, 0))
	assert.Equal(t, 2, buf.FindEolIndex(5, 0))
	assert.Equal(t, 3, buf.FindEolIndex(6, 0))
	assert.Equal(t, 3, buf.FindEolIndex(7, 0))
	assert.Equal(t, 4, buf.FindEolIndex(8, 0))
	assert.Equal(t, 4, buf.FindEolIndex(8, 0))
}
