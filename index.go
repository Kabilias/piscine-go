package piscine

func Index(s string, toFind string) int {
	runes := []rune(s)
	runes_to_find := []rune(toFind)
	length_main := 0
	for i := range runes {
		length_main = i + 1
	}
	length := 0
	for i := range runes_to_find {
		length = i + 1
	}
	if length_main < length {
		return -1
	}

	if length == 0 {
		return 0
	}

	found := false
	for i := range runes {
		if runes[i] == runes_to_find[0] {
			if length == 1 {
				return i
			}
			if length_main-i < length {
				return -1
			}
			for j := 0; j < length-1; j++ {
				if runes_to_find[j] == runes[i+j] {
					found = true
				} else {
					found = false
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}
