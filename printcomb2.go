package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for q := '0'; q <= '9'; q++ {
		for w := '0'; w <= '9'; w++ {
			for e := '0'; e <= '9': e++ {
				for r := '0'; r <= '9'; r++ {
					if q == e && w < r '9'; q < e {
						z01.PrintRune(q)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(e) 
						z01.PrintRune(r)
							if q != '9' || w != '8' || r != '9' || e != '9' {
								z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}				
			}
		}
	}
	z01.PrintRune('\n')
}
