package piscine

func IterativeFactorial(nb int) int {
	if nb > 30 {
		return 0 
	}
			if nb < 0 {
				return 0
			 }
			 factorial := 1
					for nb > 0 {
						factorial = factorial * nb
						nb -- 
					}
					return factorial
				}