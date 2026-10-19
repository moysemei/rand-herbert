// Package parking returns the price of a parking lot bill.
package parking

/*
 Plano:
  - Para o BilledHours: 60 minutos representam exatamente 1 hora, 61 minutos representam
    2 horas cobradas. Portanto, precisa-se arredondar a quantidades de horas para cima.
  - Para Price(125): 125 minutos ultrapassam a tolerância de 15 minutos. BilledHours(125)
    retorna 3 horas. A primeira hora custa 8 reais e as outras 2 horas custam 5 horas cada.
*/

import "fmt"

const (
	minutesPerHour           = 60
	freeMinutes              = 15
	firstHourPriceReais      = 8
	AdditionalHourPriceReais = 5
)

// FormatDuration shows minutes in hours and minutes
func FormatDuration(minutes int) string {
	hours := minutes / minutesPerHour
	remainingMinutes := minutes % minutesPerHour

	return fmt.Sprintf("%dh %dm", hours, remainingMinutes)
}

// BilledHours calculates minutes in hours
func BilledHours(minutes int) int {
	hours := (minutes + minutesPerHour - 1) / minutesPerHour
	return hours
}

// Price calculates the price based in hours
func Price(minutes int) int {
	if minutes <= freeMinutes {
		return 0
	}

	hours := BilledHours(minutes)
	price := firstHourPriceReais + (AdditionalHourPriceReais * (hours - 1))
	return price
}
