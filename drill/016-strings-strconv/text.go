// Package text provides exercises for string manipulation
// and conversion between strings and integers.
package text

import (
	"strconv"
	"strings"
)

// Plano:
// Para a primeira func: strings.ToLower() e // strings.TrimSpace(), na verdade strings.ReplaceAll
// Para segunda func: também ToLower e TrimSpace
// Para a terceira func: strconv.Atoi()
// À mão: Normalize("A man a plan a canal Panama") -> amanaplanacanalpanama
// SumNumbers("4 x 5"): esse não entendi como faz.

// Normalize returns a string without space and in lower case.
func Normalize(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func reverse(s string) string {
	result := ""

	for _, r := range s {
		result = string(r) + result
	}

	return result
}

// IsPalindrome compare two strings and return if they are palindrome
func IsPalindrome(s string) bool {
	s = Normalize(s)
	newString := reverse(s)
	return newString == s
}

// SumNumbers return the sum of numbers wrote as strings and if the sum is possible
func SumNumbers(s string) (int, bool) {
	sum := 0

	for _, word := range strings.Fields(s) {
		numb, err := strconv.Atoi(word)

		if err != nil {
			return 0, false
		}

		sum += numb
	}
	return sum, true
}

// Dificuldades:
// Para a segunda func, dificuldade com a lógica de comparar as duas strings e fazer o return.
// Comecei criando uma variável newString := Normalize(s) e depois usando reverse, mas não tá fazendo sentido.
// Tá eu tava criando a variável já no Normalize, mas eu precisava normalizar o s primeiro e dai sim comparar.
// Na terceira func, eu fiz algo como:
/*
sum := 0

	for range strings.Fields(s) {
		numb, err := strconv.Atoi(s)
		sum += numb
		if err != nil {
			return 0, false
		}
	}
	return sum, true
*/

// o que me pegou foi, como posso verificar o erro depois, já que terei a variável word e numb, e ai
// pensando nisso eu tirei o word, pois pensei que não saberia o que fazer com duas variáveis.
// E no fim o que me destruiu foi não ter colocado em strconv.Atoi(word) mas sim s.
