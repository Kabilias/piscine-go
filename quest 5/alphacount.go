package piscine

func AlphaCount(str string) int {

	g := 0
	for range str {
		g++
	}
	d := 0
	for a := 0; a < g; a++ {
		k := rune(str[a])
		if (k >= 65) && (k <= 90) || (k >= 97 && k <= 122) {
			d++
		}
	}
	return d
}
