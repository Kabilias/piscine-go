package piscine

func AppendRange(min, max int) []int {
	var ints []int
	if min >= max {
		return ints
	}
	for j := min; j < max; j++ {
		ints = append(ints, j)
	}
	return ints
}
