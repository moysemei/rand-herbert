package redo

func ReverseNumber(n int) int {
	reversed := 0

	for n > 0 {
		digit := n % 10

		reversed = reversed*10 + digit

		n /= 10
	}

	return reversed
}
