// Package digitloop for-while study purpose
package digitloop

// Plano:
// Primeira func, utilizar n%10 para remover os números, após pegar cada um somá-los.
// Segunda func, novamente, enquanto fui digitar as ideias vieram e eu não fiz o plano, mas basicamente era dividir e adicionar +1 ao contador.
// Terceira func, eu tinha ideia de extrair o número, mas montar novamente com a fórmula eu não tinha um plano.

// SumDigits explit the n in section of each number and sum them.
func SumDigits(n int) int {
	f := 0
	for n > 0 {
		f += n % 10
		n /= 10
	}
	return f
}

// CountDigits reports how many digits n have
func CountDigits(n int) int {
	count := 0
	if n == 0 {
		count++
	}
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

// ReverseNumber returns the reverse number of n
func ReverseNumber(n int) int {
	rev := 0
	for n > 0 {
		digit := n % 10
		rev = rev*10 + digit
		n /= 10
	}
	return rev
}

// Dificuldades:
// Primeira func, em mudar o valor do n, se não caia em loop infinito, mas consegui resolver após um tempo.
// Segunda func, sinto que ficou "ajambrado" mas deu certo.
// Terceira func, dificuldade foi em adicionar a fórmula, tipo, montar o número novamente depois de pegar o último numero com n%10.
