package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else if n < 0 {
		return
	}
	var g []int
	for n != 0 {
		x := n % 10
		g = append(g, x)
		n = n / 10
	}

	g = SortIntegerTable(g)

	for i := range g {
		ch := '0'
		for y := 0; y < ints[i]; y++ {
			ch++
		}
		z01.PrintRune(ch)
	}
}

func SortIntegerTable(table []int) []int {
	min_idx := 0
	length := 0
	for a := range table {
		length = a + 1
	}
	for i := 0; i < length-1; i++ {
		min_idx = i
		for j := i + 1; j < length; j++ {
			if table[j] < table[min_idx] {
				min_idx = j
			}
		}
		table[i], table[min_idx] = table[min_idx], table[i]
	}
	return table
}