// Package puzzles contains number and loop exercises.
package puzzles

// Plano:
// isPrime: procurar divisores de n; parar se encontrar um.
// sumDigits: somar cada dígito, reduzindo n até zero.
// CountPrimes: contar números primos de 2 até limit.
// IsPalindromeNumber: inverter n e comparar com o original.
// DigitalRoot: somar os dígitos até restar apenas um.

// Contas à mão:
// IsPalindromeNumber(10): 10 -> 1; 10 != 1, false.
// DigitalRoot(9875): 9875 -> 29 -> 11 -> 2.

// isPrime reports whether n is a prime number.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for d := 2; d <= n/d; d++ {
		if n%d == 0 {
			return false
		}
	}

	return true
}

// sumDigits returns the sum of the digits of n.
func sumDigits(n int) int {
	total := 0

	for n > 0 {
		total += n % 10
		n /= 10
	}

	return total
}

// CountPrimes returns the number of primes up to limit, inclusive.
func CountPrimes(limit int) int {
	count := 0

	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

// IsPalindromeNumber reports whether n reads the same backwards.
func IsPalindromeNumber(n int) bool {
	original := n
	reversed := 0

	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}

	return original == reversed
}

// DigitalRoot returns the repeated digit sum until one digit remains.
func DigitalRoot(n int) int {
	for n >= 10 {
		n = sumDigits(n)
	}

	return n
}
