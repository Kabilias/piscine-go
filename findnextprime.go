package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	} else if nb > 1000000100 {
		return 0
	}
	if nb%2 == 0 {
		return FindNextPrime(nb + 1)
	}
	not_prime := false
	for k := 2; k <= nb/2; k++ {
		temp := nb / k
		if temp*k == nb {
			not_prime = true
		}
	}

	if not_prime {
		return FindNextPrime(nb + 1)
	} else {
		return nb
	}
}
