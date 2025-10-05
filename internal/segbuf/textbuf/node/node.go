package node

var NIL *Node

func init() {
	NIL = &Node{
		Red: false,
	}

	NIL.P = NIL
	NIL.Left = NIL
	NIL.Right = NIL
}

type Node struct {
	Red          bool
	P            *Node
	Left         *Node
	Right        *Node
	TotalLen     int
	TotalEolsLen int
	BufIndex     int
	Start        int
	Len          int
	EolsStart    int
	EolsLen      int
}

func Create(bufIndex int, start int, len int, eolsStart int, eolsLen int) *Node {
	return &Node{
		Red:          true,
		P:            NIL,
		Left:         NIL,
		Right:        NIL,
		TotalLen:     len,
		TotalEolsLen: eolsLen,
		BufIndex:     bufIndex,
		Start:        start,
		Len:          len,
		EolsStart:    eolsStart,
		EolsLen:      eolsLen,
	}
}

func (nd *Node) Clone(p *Node) *Node {
	if nd == NIL {
		return NIL
	}

	r := &Node{
		Red:          nd.Red,
		P:            p,
		TotalLen:     nd.TotalLen,
		TotalEolsLen: nd.TotalEolsLen,
		BufIndex:     nd.BufIndex,
		Start:        nd.Start,
		Len:          nd.Len,
		EolsStart:    nd.EolsStart,
		EolsLen:      nd.EolsLen,
	}

	r.Left = nd.Left.Clone(r)
	r.Right = nd.Right.Clone(r)

	return r
}

func (nd *Node) Find(index int) (*Node, int) {
	for nd != NIL {
		if index < nd.Left.TotalLen {
			nd = nd.Left
			continue
		}

		index -= nd.Left.TotalLen

		if index < nd.Len {
			return nd, index
		}

		index -= nd.Len
		nd = nd.Right
	}

	return nil, 0
}

func Bubble(nd *Node) {
	for nd != NIL {
		nd.TotalLen = nd.Left.TotalLen + nd.Len + nd.Right.TotalLen
		nd.TotalEolsLen = nd.Left.TotalEolsLen + nd.EolsLen + nd.Right.TotalEolsLen

		nd = nd.P
	}
}

func Minimum(nd *Node) *Node {
	for nd.Left != NIL {
		nd = nd.Left
	}

	return nd
}

func Maximum(nd *Node) *Node {
	for nd.Right != NIL {
		nd = nd.Right
	}

	return nd
}

func Successor(nd *Node) *Node {
	if nd.Right != NIL {
		return Minimum(nd.Right)
	}

	y := nd.P

	for y != NIL && nd == y.Right {
		nd = y
		y = y.P
	}

	return y
}

func (nd *Node) Validate() {
	// 1. Every node is either red or black.
	// 2. The root is black.
	if nd.P == NIL {
		if nd.Red {
			panic("2")
		}
	}

	nd.validate()

	// 5. For each node, all simple paths from the node to descendant leaves contain the same number of black nodes.
	parents := map[*Node]bool{}
	collectLeafs(nd, parents)

	heights := []int{}
	for node := range parents {
		height := 0
		x := node
		for x.P != NIL {
			if !x.Red {
				height += 1
			}
			x = x.P
		}
		heights = append(heights, height)
	}

	for _, h := range heights {
		if heights[0] != h {
			panic("5")
		}
	}

}

func (nd *Node) validate() {
	// 3. Every leaf (NIL) is black.
	if nd == NIL {
		if nd.Red {
			panic("3")
		}
	} else {
		// 4. If a node is red, then both its children are black.
		if nd.Red {
			if nd.Left.Red || nd.Right.Red {
				panic("4")
			}
		}

		nd.Left.validate()
		nd.Right.validate()

		// 6. len > 0
		if nd.Len <= 0 {
			panic("6")
		}
	}
}

func collectLeafs(nd *Node, parents map[*Node]bool) {
	if nd != NIL {
		if nd.Left == NIL || nd.Right == NIL {
			parents[nd] = true
		}

		collectLeafs(nd.Left, parents)
		collectLeafs(nd.Right, parents)
	}
}
