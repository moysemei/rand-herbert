// Package calendar reports day of the week, grade and how many days are in specif month and year
package calendar

/*
 Plan:
  - primeira func, tranquilo, só fiquei pensando se teria algum jeito de listar os nomes
    dos dias da semana mais facilmente, mas talvez envolva o pacote time, ou algo com date()
  - segunda func, coloquei em dificuldades o que tentei
  - terceira func,daysinmoth(2, 1900) pega o caso de fevereiro usando isleapyear
*/

// DayType reports if a day of the week is weekend or weekday
func DayType(day string) string {
	switch day {
	case "saturday", "sunday":
		return "weekend"
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		return "weekday"
	default:
		return "invalid"
	}
}

// Grade reports the grade based in score
func Grade(score int) string {
	switch {
	case score < 0, score > 100:
		return "invalid"
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

// isLeapYear is used to tell if a year is leap or not
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// DaysInMonth reports how many days are in a specific month in specific year
func DaysInMonth(month, year int) int {
	switch {
	case month == 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	case month == 4, month == 6, month == 9, month == 11:
		return 30
	case month <= 0, month > 12:
		return 0
	default:
		return 31
	}
}

/*
 Dificuldades:
 - pensar na lógica de Grade: pensando em como englobar as casas dos números, como 90 a 100 etc.
 - primeira tentativa: case 90 >= score <= 100: return "A" - deu erro.
 - segunda tentativa: case score >= 90, score <= 100: return "A" - deu boa até agora.
 - não deu certo, invertendo ou não o switch-case, eles caem no case errado, o que faz sentido
   pois se coloco score <= 100 todos vão entrar.
   - abri o primeiro hint pra func Grade: com isso cheguei na lógica de limitar por baixo,
     mas o 101 me pegou, tive que abrir a segunda hint.
 - terceira func complicou bastante, não consegui montar nada ainda sem as hints.
 - consegui a terceira sem hints!
*/
