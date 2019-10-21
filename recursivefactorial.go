package piscine

func RecursiveFactorial(nb int) int {
	i := 0
	if nb == 0 {
		return 1
	} else if nb < 0 || nb >= 13 {
		return 0
	} else {
		i = nb * RF(nb-1)
	}
	return a
}
