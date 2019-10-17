package piscine

func UltimateDivMod(a *int, b *int) {
	k := *a
	d := *b
	*a = *a / (*b)

	*b = k % (*b)

}
