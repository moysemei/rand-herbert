package redo

func Power(base, exp int) int {
	product := 1
	for i := 1; i <= exp; i++ {
		product *= base
	}
	return product
}
