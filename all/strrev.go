package piscine

func StrRev(s string) string {

	var i string
	for _, a := range s {
		i = string(a) + i
	}
	return i
}
