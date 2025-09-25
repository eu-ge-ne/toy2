package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNIL(t *testing.T) {
	expected := &Node{
		Red:          false,
		P:            NIL,
		Left:         NIL,
		Right:        NIL,
		TotalLen:     0,
		TotalEolsLen: 0,
		BufIndex:     0,
		Start:        0,
		Len:          0,
		EolsStart:    0,
		EolsLen:      0,
	}

	assert.Equal(t, expected, NIL)
}

func TestCreate(t *testing.T) {
	node := Create(0, 0, 0, 0, 0)

	assert.Equal(t, true, node.Red)
}

func TestClone(t *testing.T) {
	parent := Create(0, 0, 10, 0, 1)
	child := Create(1, 0, 20, 0, 2)

	parent.Right = child
	child.P = parent

	clone := parent.Clone(NIL)

	assert.Equal(t, clone, parent)
	assert.Equal(t, clone.Right, child)
}

func TestBubble(t *testing.T) {
	parent := Create(0, 0, 0, 0, 0)
	child := Create(0, 0, 0, 0, 0)

	parent.Right = child
	child.P = parent

	Bubble(child)

	assert.Equal(t, 0, parent.TotalLen)

	child.Len = 10

	Bubble(child)

	assert.Equal(t, 10, parent.TotalLen)
}

func TestMinimum(t *testing.T) {
	parent := Create(0, 0, 0, 0, 0)
	child := Create(0, 0, 0, 0, 0)

	parent.Left = child
	child.P = parent

	assert.Equal(t, child, Minimum(parent))
}
