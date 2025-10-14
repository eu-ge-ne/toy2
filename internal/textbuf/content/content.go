package content

import (
	"iter"

	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
	"github.com/eu-ge-ne/toy2/internal/textbuf/piece"
)

type Content struct {
	Table []*piece.Piece
}

func (c *Content) Create(data []byte) *node.Node {
	piece := piece.Create(data)
	c.Table = append(c.Table, &piece)
	pieceIndex := len(c.Table) - 1

	return node.Create(pieceIndex, 0, piece.Len, 0, len(piece.Eols))
}

func (c *Content) Split(x *node.Node, index int, gap int) *node.Node {
	piece := c.Table[x.PieceIndex]

	start := x.Start + index + gap
	len := x.Len - index - gap

	c.resize(x, index)
	node.Bubble(x)

	eols_start := piece.FindEolIndex(start, x.EolsStart+x.EolsLen)
	eols_end := piece.FindEolIndex(start+len, eols_start)
	eols_len := eols_end - eols_start

	return node.Create(x.PieceIndex, start, len, eols_start, eols_len)
}

func (c *Content) Read(x *node.Node, offset int, n int) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
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

func (c *Content) Chunk(x *node.Node, offset int) []byte {
	count := x.Len - offset

	return c.Table[x.PieceIndex].Read(x.Start+offset, x.Start+offset+count)
}

func (c *Content) Growable(x *node.Node) bool {
	piece := c.Table[x.PieceIndex]

	return (piece.Len < 100) && (x.Start+x.Len == piece.Len)
}

func (c *Content) Grow(x *node.Node, data []byte) {
	c.Table[x.PieceIndex].Append(data)

	c.resize(x, x.Len+len(data))
}

func (c *Content) TrimStart(x *node.Node, n int) {
	piece := c.Table[x.PieceIndex]

	x.Start += n
	x.Len -= n

	x.EolsStart = piece.FindEolIndex(x.Start, x.EolsStart)

	eols_end := piece.FindEolIndex(x.Start+x.Len, x.EolsStart)

	x.EolsLen = eols_end - x.EolsStart
}

func (c *Content) TrimEnd(x *node.Node, n int) {
	c.resize(x, x.Len-n)
}

func (c *Content) resize(x *node.Node, len int) {
	piece := c.Table[x.PieceIndex]

	x.Len = len

	eols_end := piece.FindEolIndex(x.Start+x.Len, x.EolsStart)

	x.EolsLen = eols_end - x.EolsStart
}
