package piscine

func DivMod(a *int, b *int) {
	c := *a
	*a = *a / (*b)

	*b = c % (*b)

}
