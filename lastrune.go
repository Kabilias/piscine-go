package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	last := 0
	for l := range s {
		last = l
	}
	return runes[last]
}
