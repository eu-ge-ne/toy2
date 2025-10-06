package std

func Clamp(x, min, max int) int {
	switch {
	case x < min:
		return min
	case x > max:
		return max
	default:
		return x
	}
}

func Sum(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}
