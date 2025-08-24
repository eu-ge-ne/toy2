package std

func Sum(arr []int) int {
	r := 0

	for _, v := range arr {
		r += v
	}

	return r
}
