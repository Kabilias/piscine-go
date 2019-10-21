package piscine

func IsPrime(nb int) bool {
	if value <= 1 {
		return false
	}
	for a := 2; a < value; a++ {
		if value%a == 0 {
			return false
		}
	}
	return true
}
