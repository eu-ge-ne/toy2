package syntax

type span struct {
	start    int
	end      int
	captures []int
	color    CharFgColor
}

func (s span) match(idx int) int {
	if idx < s.start {
		return -1
	}

	if idx < s.end {
		return 0
	}

	return 1
}
