package node

var NIL *Node

func init() {
	NIL = &Node{
		Nil: true,
		Red: false,
	}

	NIL.P = NIL
	NIL.Left = NIL
	NIL.Right = NIL
}

type Node struct {
	Nil          bool
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

func Create(bufIndex int, start int, len int, eols_start int, eols_len int) Node {
	return Node{
		Nil:          false,
		Red:          true,
		P:            NIL,
		Left:         NIL,
		Right:        NIL,
		TotalLen:     len,
		TotalEolsLen: eols_len,
		BufIndex:     bufIndex,
		Start:        start,
		Len:          len,
		EolsStart:    eols_start,
		EolsLen:      eols_len,
	}
}

func (x *Node) Find(index int) (*Node, int) {
	for !x.Nil {
		if index < x.Left.TotalLen {
			x = x.Left
			continue
		}

		index -= x.Left.TotalLen

		if index < x.Len {
			return x, index
		}

		index -= x.Len
		x = x.Right
	}

	return nil, 0
}

func Bubble(x *Node) {
	for !x.Nil {
		x.TotalLen = x.Left.TotalLen + x.Len + x.Right.TotalLen
		x.TotalEolsLen = x.Left.TotalEolsLen + x.EolsLen + x.Right.TotalEolsLen

		x = x.P
	}
}

func Minimum(x *Node) *Node {
	for !x.Left.Nil {
		x = x.Left
	}

	return x
}

func Maximum(x *Node) *Node {
	for !x.Right.Nil {
		x = x.Right
	}

	return x
}

func Successor(x *Node) *Node {
	if !x.Right.Nil {
		return Minimum(x.Right)
	}

	y := x.P

	for !y.Nil && x == y.Right {
		x = y
		y = y.P
	}

	return y
}

func (x Node) Validate() {
	// 1. Every node is either red or black.
	// 2. The root is black.
	if x.P.Nil {
		if x.Red {
			panic("2")
		}
	}

	x.validate()

	// 5. For each node, all simple paths from the node to descendant leaves contain the same number of black nodes.
	leafs := map[Node]bool{}
	collectLeafs(x, leafs)

	heights := []int{}
	for node := range leafs {
		height := 0
		x := node
		for !x.P.Nil {
			if !x.Red {
				height += 1
			}
			x = *x.P
		}
		heights = append(heights, height)
	}

	for _, h := range heights {
		if heights[0] != h {
			panic("5")
		}
	}

}

func (x Node) validate() {
	// 3. Every leaf (NIL) is black.
	if x.Nil {
		if x.Red {
			panic("3")
		}
	} else {
		// 4. If a node is red, then both its children are black.
		if x.Red {
			if x.Left.Red || x.Right.Red {
				panic("4")
			}
		}

		x.Left.validate()
		x.Right.validate()

		// 6. len > 0
		if x.Len <= 0 {
			panic("6")
		}
	}
}

func collectLeafs(x Node, leaf_parents map[Node]bool) {
	if !x.Nil {
		if x.Left.Nil || x.Right.Nil {
			leaf_parents[x] = true
		}

		collectLeafs(*x.Left, leaf_parents)
		collectLeafs(*x.Right, leaf_parents)
	}
}
