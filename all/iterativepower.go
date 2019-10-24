package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	p := 1
	for power > 0 {
		p = p * nb
		power--
	}
	return p
}
