package numbers

import "fmt"

/*
 Plano:
  - Para primeira func, irei usar n % 2 == 0;
  - Para a segunda func, preciso criar uma constante para segundo em hora, e 60 segundos de um minuto.
  - Para a terceira func, preciso fazer algo tipo year % 4 == 0 && year % 100 == 0 && year % 400 == 0, Não sei se dá pra colocar
    essas três condições.
*/

const (
	oneHour   = 3600
	oneMinute = 60
)

func IsEven(n int) bool {
	return n%2 == 0
}

func SecondsToClock(total int) string {
	minutes := total % oneHour / oneMinute
	hours := total / oneHour
	seconds := total % oneHour
	return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
}

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

/*
 Dificuldades:
  - definir o sprintf com as variáveis, a ordem e não entrou muito bem na minha cabeça da onde vem a hour, minutos e segundos para enfilerar
    no sprintf.
*/
