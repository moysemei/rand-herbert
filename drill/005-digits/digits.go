// Package digits reports tools to tell us about numbers
package digits

import "fmt"

/*
 Plano:
 - Não entendi muito bem o fato de desmotar 472 e montar 274 a mão, então farei o que eu entendi:
   4*100 + 7*10 + 2*1 = 472, 2*100 + 7*10 + 4*1 = 274.
 - Com relação a primeira func, não entendi como posso fazer sem usar slices. Atualização, acho que entendi essa primeira parte
   da sobra da divisão
 - Segunda func, Irei criar três constantes pra depois colocar no sprintf
*/

// LastDigit shows the last number of n.
func LastDigit(n int) int {
	return n % 10
}

// DescribeDigits for n between 100 and 999 how many hundreds, tens and ones n has.
func DescribeDigits(n int) string {
	hundreds := n / 100
	tens := (n % 100) / 10
	ones := n % 10
	return fmt.Sprintf("hundreds: %d, tens: %d, ones: %d", hundreds, tens, ones)
}

// ReverseThreeDigits for n between 100 and 999 return the number with reversed digits.
func ReverseThreeDigits(n int) int {
	hundreds := n / 100
	tens := (n % 100) / 10
	ones := n % 10
	return ones*100 + tens*10 + hundreds
}

/*
 Dificuldades antes de ler os hints:
 - a parte da dezena, constante tens, não estou conseguindo pensar na lógica de extrair o número. fui até o hint nível 2 e então consegui
   para a func DescribeDigits.
 - Para a última func, tive que usar a última hint
 - E usei bastante a colinha do README.md desse drill 005, não sei se eu poderia, mas fiz, caso não possa me diga na próxima.
*/
