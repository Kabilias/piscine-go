package piscine

func TrimAtoi(s string) int {
	runes := []rune(s)
	var runes_int []rune
	for i := range runes {

		if runes[i] == '-' || runes[i] == '+' {
			length := 0
			for j := range runes_int {
				length = j + 1
			}

			if length == 0 {
				runes_int = append(runes_int, runes[i])
			}
		}
		if runes[i] > 47 && runes[i] < 58 {
			runes_int = append(runes_int, runes[i])
		}
	}

	return Atoi(runes_int)
}

func Atoi(runes []rune) int {

	num := 0
	length := 0
	for i := range runes {
		length = i + 1
	}
	if length == 0 {
		return 0
	}
	ten := 1
	for j := 0; j < length-1; j++ {
		if runes[j] == '+' || runes[j] == '-' {
			continue
		}
		ten *= 10
	}
	for i := range runes {
		if (runes[0] == '+' || runes[0] == '-') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		temp := 0
		for k := '0'; k < runes[i]; k++ {
			temp++
		}
		num += temp * ten
		ten /= 10
	}
	if runes[0] == '-' {
		num *= -1
	}
	return num
}
