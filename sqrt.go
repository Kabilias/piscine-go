package piscine

func Sqrt(nb int) int {
	for a := 0; (a * a) <= nb; r++ {
		if a*a == nb {
			return a
		}
	}
	return 0
}
