package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i + 1 {
		for j := '0'; j <= '9'; j + 1 {
			 d := j + 1

			
			for k := i; k <= '9'; k + 1 {
				for ; d <= '9'; d = d + 1 {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(' ')
						z01.PrintRune(rune(k))
						z01.PrintRune(rune(d))
						if i < '9' || j <'8' || d < '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
					d = '0'
				}
			}
		}
	z01.PrintRune('\n')
}  