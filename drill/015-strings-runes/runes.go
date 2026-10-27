// Package runes provides exercises for working with strings and runes.
package runes

// Plano:
// Primeira func, montei para percorrer o s, que vai ser o string utilizado
// e ai já pensei em usar o if e comparar com cada vogal.
// Segunda func, preciso dar um retorno multiplo... utilizando len(s) para contar a quantidade de bytes, e runes vou precisar
// de um contador.
// Terceira func, sabe Deus...
// Tabela Reverse("ação"):
// Volta | r   | result
// 1     | 'a' | "a"
// 2     | 'ç' | "ça"
// 3     | 'ã' | "ãça"
// 4     | 'o' | "oãça"

// CountVowels returns the number of lowercase unaccented vowels in s.
func CountVowels(s string) int {
	count := 0
	for _, r := range s {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			count++
		}
	}
	return count
}

// Sizes returns the number of bytes and runes in s.
func Sizes(s string) (bytes, runes int) {
	bytes = len(s)

	for range s {
		runes++
	}

	return bytes, runes
}

// Reverse returns a string with the runes of s in reverse order.
func Reverse(s string) string {
	result := ""

	for _, r := range s {
		result = string(r) + result
	}
	return result
}

// Dificuldades:
// Para primeira func, foi algo mais sintaxe/idiomático, pois eu primeiro usei s
// s == a, e, i, o, u {} mas na verdade era para ter usado o r, mas ai então tive que
// acrescentar as aspas simples para comparar e então o || (ou), mas ainda não deu certo,
// pois eu não estava colocando r em cada letra, somente tipo r == 'a' ou 'e' ou 'i'.
// em ou quero dizer || mas tenho que ficar copiando e colando pois não sei escrever esses
// pipes em meu teclado.
// Segunda func, executei o plano com for _, r := range s { count ++ } mas ai deu r declarado
// mas não usado
// Terceira func, dificuldade em concatenar minhas ideias, e em executar, primeira ideia foi usar
// for i, r = range s {} e ai fazer com que a posição que o i me traga sendo ela a última vá para
// o começo, mas não tenho ideia como fazer isso.
// Precisei de ajuda para compreender como a concatenação permite inverter uma string.
