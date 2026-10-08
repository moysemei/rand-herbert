// Package convert celsius to fahrenheit, minutes to hours and describe a persona name and age.
package convert

/* Plan
-Para a primeira func, criar a variável F := e começar a fórmula.
-Para a segunda func, converter minutes em float e devolver com a divisão por 60.
-Para a terceira func, concatenar todas as informações e colocar as variáveis no final, mas convertendo int para string.
*/
import (
	"fmt"
	"strconv"
)

// CelsiusToFahrenheit converts Celsius tempture to Fahrenheit tempture
func CelsiusToFahrenheit(c float64) float64 {
	temp := c*9/5 + 32
	return temp
}

// MinutesToHours convertes minutes in hours
func MinutesToHours(minutes int) float64 {
	f := float64(minutes)
	return f / 60
}

// Describe describe a persona name and age
func Describe(name string, age int) string {
	return fmt.Sprintf("%s is %v years old", name, strconv.Itoa(age))
}

/*
Dificuldades:
 - Eu tava colocando const temp := c * 9 / 5 + 32, mas não posso fazer isso.
 - na segunda função tentei fazer direto com float64(minutes / 60) mas quando rodava os testes não dava certo
 - na terceira função tive um pouco de dificuldade para converter, confesso que utilizei o review.md de ontem
   para ter a ideia do fmt.Sprintf. o strconv.Itoa() eu pesquisei.
*/
