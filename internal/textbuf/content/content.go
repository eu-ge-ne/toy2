package content

import (
	"iter"

	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
	"github.com/eu-ge-ne/toy2/internal/textbuf/piece"
)

type Content struct {
	Table []*piece.Piece
}

func (c *Content) Create(text string) *node.Node {
	buffer := piece.Create(text)
	c.Table = append(c.Table, &buffer)
	BufIndex := len(c.Table) - 1

	return node.Create(BufIndex, 0, buffer.Len, 0, len(buffer.Eols))
}

func (c *Content) Split(x *node.Node, index int, gap int) *node.Node {
	buf := c.Table[x.PieceIndex]

	start := x.Start + index + gap
	len := x.Len - index - gap

	c.resize(x, index)
	node.Bubble(x)

	eols_start := buf.FindEolIndex(start, x.EolsStart+x.EolsLen)
	eols_end := buf.FindEolIndex(start+len, eols_start)
	eols_len := eols_end - eols_start

	return node.Create(x.PieceIndex, start, len, eols_start, eols_len)
}

func (c *Content) Read(x *node.Node, offset int, n int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for x != node.NIL && (n > 0) {
			count := min(x.Len-offset, n)

			if !yield(c.Table[x.PieceIndex].Read(x.Start+offset, x.Start+offset+count)) {
				return
			}

			x = node.Successor(x)
			offset = 0
			n -= count
		}
	}
}

func (c *Content) Growable(x *node.Node) bool {
	buf := c.Table[x.PieceIndex]

	return (buf.Len < 100) && (x.Start+x.Len == buf.Len)
}

func (c *Content) Grow(x *node.Node, text string) {
	c.Table[x.PieceIndex].Append(text)

	c.resize(x, x.Len+len(text))
}

func (c *Content) TrimStart(x *node.Node, n int) {
	buf := c.Table[x.PieceIndex]

	x.Start += n
	x.Len -= n

	x.EolsStart = buf.FindEolIndex(x.Start, x.EolsStart)

	eols_end := buf.FindEolIndex(x.Start+x.Len, x.EolsStart)

	x.EolsLen = eols_end - x.EolsStart
}

func (c *Content) TrimEnd(x *node.Node, n int) {
	c.resize(x, x.Len-n)
}

func (c *Content) resize(x *node.Node, len int) {
	buf := c.Table[x.PieceIndex]

	x.Len = len

	eols_end := buf.FindEolIndex(x.Start+x.Len, x.EolsStart)

	x.EolsLen = eols_end - x.EolsStart
}
