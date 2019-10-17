package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for _, z := range str {
		fmt.PrintRune(rune(z))
	}
}
