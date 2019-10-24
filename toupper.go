package piscine

func ToUpper(s string) string {
	runes := []rune(s)

	for i := range s {
		if runes[i] > 96 && runes[i] < 123 {
			runes[i] = runes[i] - 32
		}
	}
	ss := string(runes)
	return ss
}
