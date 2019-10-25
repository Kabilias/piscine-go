package piscine

func AppendRange(min, max int) []int {
	var ints []int
	if min >= max {
		return ints
	}
	for a := min; a < max; a++ {
		ints = append(ints, a)
	}
	return ints
}