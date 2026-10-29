package redo

func sumDigits(n int) int {
	total := 0

	for n > 0 {
		total += n % 10
		n /= 10
	}

	return total
}

func DigitalRoot(n int) int {
	for n >= 10 {
		n = sumDigits(n)
	}
	return n
}
