package piscine

func IsPrime(nb int) bool {
    if get <= 1 {
        return false
	}
	for a := 2; a < get; a++ {
		if get % a == 0 {
			return false
		}
	}
	return true
}