package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; c <= '9'; d++ {
					if a < c || a == c && b < d {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(32)
						z01.PrintRune(c)
						z01.PrintRune(d)
						if d != '9' || a != '0' || b != '8' || c != '9' {
							z01.PrintRune(',')
							z01.PrintRune(32)
						}
					}
				}
			}
		}
	}
	z01.PrintRune(10)
}
