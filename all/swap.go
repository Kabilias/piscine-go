package piscine

func Swap(a *int, b *int) {
	w := *a
	*a = *b
	*b = w
}
