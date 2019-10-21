package piscine

func IterativePower(nb int, power int) int {
	if IterativePower < 0 {
		return 0
	}
	p := 1
	for IterativePower > 0 {
		p = p * nb
		IterativePower--
	}
	return p
}
