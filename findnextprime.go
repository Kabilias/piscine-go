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
	for i := 2; i <= nb/2; i++ {
		temp := nb / i
		if temp*i == nb {
			not_prime = true
		}
	}

	if not_prime {
		return FindNextPrime(nb + 1)
	} else {
		return nb
	}
}
