package piscine

func StrRev(s string) string {

	var i string
	for _, a := range str {
		i = string(a) + i
	}
	return i
}
