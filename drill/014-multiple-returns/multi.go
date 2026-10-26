// Package multi provides exercises with multiple return values.
package multi

// Plano:
// SplitDigits: separar um número em centena, dezena e unidade,
// retornando os três valores nessa ordem.
//
// Divide: verificar se b é zero. Se for, retornar 0 e false.
// Caso contrário, retornar a divisão inteira e true.
//
// MinMaxDigit: percorrer os dígitos de n e atualizar o menor e
// o maior encontrados, retornando min e max nessa ordem.
//
// Tabela MinMaxDigit(4729):
// Volta | digit | min | max
//   1   |   9   |  9  |  9
//   2   |   2   |  2  |  9
//   3   |   7   |  2  |  9
//   4   |   4   |  2  |  9
//
//
// Para n = 0, min e max começam em 0.
// O loop não executa e a função retorna 0, 0.
//
// Revisão: d <= n/d é equivalente a d*d <= n para d positivo.
// A divisão evita o possível overflow da multiplicação d*d.

// SplitDigits returns the hundreds, tens, and ones digits of n.
// It expects n to be between 100 and 999.
func SplitDigits(n int) (hundreds, tens, ones int) {

	hundreds = n / 100
	tens = (n / 10) % 10
	ones = n % 10

	return hundreds, tens, ones
}

// Divide returns the integer division of a by b.
// It returns 0 and false if b is zero.
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// MinMaxDigit returns the smallest and largest digits of n.
// It expects n to be non-negative
func MinMaxDigit(n int) (int, int) {
	min, max := n%10, n%10

	for n > 0 {
		digit := n % 10

		if digit < min {
			min = digit
		}

		if digit > max {
			max = digit
		}
		n /= 10
	}

	return min, max
}

// Dificuldades:
// Tive mais dificuldade na função MinMaxDigit.
// Entendi que precisava percorrer os dígitos e identificar
// o menor e o maior, mas tive dificuldade em transformar
// esse raciocínio em código.
//
// Inicialmente, pensei que precisaria armazenar todos os dígitos.
// Depois entendi que bastava inicializar min e max com um
// dígito real e atualizar essas variáveis durante o loop.
//
// Também precisei corrigir alguns erros de sintaxe,
// como a diferença entre = e == e a posição dos ifs.
