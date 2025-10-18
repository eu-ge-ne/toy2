package syntax

import (
	treeSitter "github.com/tree-sitter/go-tree-sitter"
)

type span struct {
	start    treeSitter.Point
	end      treeSitter.Point
	text     string
	captures []int
	color    CharFgColor
}

func (s span) match(ln, col int) int {
	if ln < int(s.start.Row) {
		return -1
	}

	if ln == int(s.start.Row) {
		if col < int(s.start.Column) {
			return -1
		}

		if col < int(s.end.Column) {
			return 0
		}
	}

	return 1
}
