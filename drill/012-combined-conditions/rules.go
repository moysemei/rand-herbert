package rules

import "strconv"

// Plano:
// FizzBuzz: testar o próprio n; retornar o texto adequado.
// IsPrime: testar divisores a partir de 2 até a raiz de n;
// retornar true se nenhum divisor for encontrado.
// NextPrime: testar candidatos maiores que n;
// parar quando IsPrime retornar true.

func FizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	if n%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}

func IsPrime(n int) bool {
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

func NextPrime(n int) int {
	candidate := n + 1

	for {
		if IsPrime(candidate) {
			return candidate
		}

		candidate++
	}
}
