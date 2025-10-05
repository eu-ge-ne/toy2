package tree

import (
	"github.com/eu-ge-ne/toy2/internal/textbuf/node"
)

type Tree struct {
	Root *node.Node
}

func (t *Tree) InsertLeft(p *node.Node, z *node.Node) {
	p.Left = z
	z.P = p

	node.Bubble(z)

	t.insertFixup(z)
}

func (t *Tree) InsertRight(p *node.Node, z *node.Node) {
	p.Right = z
	z.P = p

	node.Bubble(z)

	t.insertFixup(z)
}

func (t *Tree) InsertBefore(p *node.Node, z *node.Node) {
	if p.Left == node.NIL {
		t.InsertLeft(p, z)
	} else {
		t.InsertRight(node.Maximum(p.Left), z)
	}
}

func (t *Tree) InsertAfter(p *node.Node, z *node.Node) {
	if p.Right == node.NIL {
		t.InsertRight(p, z)
	} else {
		t.InsertLeft(node.Minimum(p.Right), z)
	}
}

func (t *Tree) insertFixup(z *node.Node) {
	for z.P.Red {
		if z.P == z.P.P.Left {
			y := z.P.P.Right
			if y.Red {
				z.P.Red = false
				y.Red = false
				z.P.P.Red = true
				z = z.P.P
			} else {
				if z == z.P.Right {
					z = z.P
					t.leftRotate(z)
				}
				z.P.Red = false
				z.P.P.Red = true
				t.rightRotate(z.P.P)
			}
		} else {
			y := z.P.P.Left
			if y.Red {
				z.P.Red = false
				y.Red = false
				z.P.P.Red = true
				z = z.P.P
			} else {
				if z == z.P.Left {
					z = z.P
					t.rightRotate(z)
				}
				z.P.Red = false
				z.P.P.Red = true
				t.leftRotate(z.P.P)
			}
		}
	}

	t.Root.Red = false
}

func (t *Tree) Delete(z *node.Node) {
	y := z
	y_original_color := y.Red
	var x *node.Node

	if z.Left == node.NIL {
		x = z.Right

		t.transplant(z, z.Right)
		node.Bubble(z.Right.P)
	} else if z.Right == node.NIL {
		x = z.Left

		t.transplant(z, z.Left)
		node.Bubble(z.Left.P)
	} else {
		y = node.Minimum(z.Right)

		y_original_color = y.Red
		x = y.Right

		if y != z.Right {
			t.transplant(y, y.Right)
			node.Bubble(y.Right.P)

			y.Right = z.Right
			y.Right.P = y
		} else {
			x.P = y
		}

		t.transplant(z, y)

		y.Left = z.Left
		y.Left.P = y
		y.Red = z.Red

		node.Bubble(y)
	}

	if !y_original_color {
		t.deleteFixup(x)
	}
}

func (t *Tree) deleteFixup(x *node.Node) {
	for x != t.Root && !x.Red {
		if x == x.P.Left {
			w := x.P.Right

			if w.Red {
				w.Red = false
				x.P.Red = true
				t.leftRotate(x.P)
				w = x.P.Right
			}

			if !w.Left.Red && !w.Right.Red {
				w.Red = true
				x = x.P
			} else {
				if !w.Right.Red {
					w.Left.Red = false
					w.Red = true
					t.rightRotate(w)
					w = x.P.Right
				}

				w.Red = x.P.Red
				x.P.Red = false
				w.Right.Red = false
				t.leftRotate(x.P)
				x = t.Root
			}
		} else {
			w := x.P.Left

			if w.Red {
				w.Red = false
				x.P.Red = true
				t.rightRotate(x.P)
				w = x.P.Left
			}

			if !w.Right.Red && !w.Left.Red {
				w.Red = true
				x = x.P
			} else {
				if !w.Left.Red {
					w.Right.Red = false
					w.Red = true
					t.leftRotate(w)
					w = x.P.Left
				}

				w.Red = x.P.Red
				x.P.Red = false
				w.Left.Red = false
				t.rightRotate(x.P)
				x = t.Root
			}
		}
	}

	x.Red = false
}

func (t *Tree) leftRotate(x *node.Node) {
	y := x.Right

	x.Right = y.Left
	if y.Left != node.NIL {
		y.Left.P = x
	}

	y.P = x.P

	if x.P == node.NIL {
		t.Root = y
	} else if x == x.P.Left {
		x.P.Left = y
	} else {
		x.P.Right = y
	}

	y.Left = x
	x.P = y

	node.Bubble(x)
}

func (t *Tree) rightRotate(y *node.Node) {
	x := y.Left

	y.Left = x.Right
	if x.Right != node.NIL {
		x.Right.P = y
	}

	x.P = y.P

	if y.P == node.NIL {
		t.Root = x
	} else if y == y.P.Left {
		y.P.Left = x
	} else {
		y.P.Right = x
	}

	x.Right = y
	y.P = x

	node.Bubble(y)
}

func (t *Tree) transplant(u *node.Node, v *node.Node) {
	if u.P == node.NIL {
		t.Root = v
	} else if u == u.P.Left {
		u.P.Left = v
	} else {
		u.P.Right = v
	}

	v.P = u.P
}
