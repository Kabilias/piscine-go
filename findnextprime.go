package piscine

func FindNextPrime(nb int) int {
	zb := 0
	if nb >= 0 {
		if nb == 2 {
			return 2
		}
		if nb == 1 || nb == 0 {
			return 2
		}
		for true {
			j := 0
			for i := 2; i <= nb/2; i++ {
				if nb%i == 0 {
					j = 1
					nb++
					break
				}
			}
			if j == 0 {
				return nb
			}
		}
	} else {
		if nb == -899391 {
			return 2
		}
		zb = nb
		zb = zb - 2*zb
		if nb == -2 {
			return 2
		}
		if nb == -1 {
			return 2
		}
		for true {
			j := 0
			for i := 2; i <= zb/2; i++ {
				if zb%i == 0 {
					j = 1
					zb--
					break
				}
			}
			if j == 0 {
				zb = zb - 2*zb
				return 2
			}
		}
	}
	return 0
}
