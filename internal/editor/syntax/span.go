package syntax

type span struct {
	start    int
	end      int
	captures []int
	color    CharFgColor
}
