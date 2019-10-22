package piscine

func FirstRune(s string) rune {
	runes := []rune(s)
	length := 0
	for i := range s {
		length = i + 1
	}
	if length == 0 {
		return -1
	}
	return runes[0]
}