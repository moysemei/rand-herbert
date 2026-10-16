# Review — Drill 007

## Veredito

Entregue, os três testes passam. O `DaysInMonth`, a parte mais longa, saiu sem dica nenhuma. No `Grade`, você chegou ao bug sozinho antes de abrir a primeira dica.

## Processo

- **"Consegui a terceira sem hints!"** A frase é sua e é o fato da semana. Ali tinha `switch`, um `if` dentro do `case` e um ajudante não exportado reaproveitando a regra do bissexto do drill 004.
- **As "Dificuldades" do `Grade` viraram um diário de depuração:** o que você tentou, o que aconteceu, e por quê. "Se coloco `score <= 100`, todos vão entrar." Você explicou o próprio bug antes de qualquer dica. É assim que se depura: hipótese, teste, explicação.
- **As dicas foram abertas uma de cada vez,** e só quando a tentativa anterior não bastou.
- **Os comentários voltaram,** no pacote e em todas as funções, começando pelo nome.
- **A sua curiosidade no plano tem resposta:** sim, o pacote `time` tem os dias da semana prontos (`time.Saturday`, e `time.Saturday.String()` devolve `"Saturday"`). Ele aparece na fase 2, na parte de biblioteca padrão.
- **O plano ficou curto.** No `Grade`, o checklist pedia a ordem dos `case` testada na mão com 95 e 101. Com o 101 no papel, o caso do `"invalid"` teria aparecido antes de você precisar da dica 2.

## Erros e melhorias

**[sintaxe] `90 >= score <= 100` é coisa do Python.** No Python, `90 <= score <= 100` funciona: ele encadeia as comparações. No Go, não:
1. Primeiro ele calcula `90 >= score`, que dá um `bool`.
2. Depois tenta comparar esse `bool` com `100`.

Por isso a mensagem é `mismatched types untyped bool and untyped int`. *Como pensar:* em Go, cada comparação é uma operação separada, e você liga as duas com `&&`: `score >= 90 && score <= 100`.

**[lógica] A vírgula no `case` é "ou", não "e".** Na segunda tentativa, `case score >= 90, score <= 100:` quer dizer "maior ou igual a 90 **ou** menor ou igual a 100", e todo número de 0 a 100 entra. Você percebeu isso sozinho. A solução final foi melhor do que juntar as duas condições com `&&`: com a faixa mais alta primeiro, cada `case` só precisa do limite de baixo, porque os números maiores já foram pegos acima. **A ordem dos `case` faz metade do trabalho.**

**[idiomático] Vírgula num `switch` sem expressão.** Seu `case score < 0, score > 100:` funciona, e o motivo é interessante: `switch {` é um atalho para `switch true {`, então cada item da vírgula é comparado com `true`, e basta um bater. É Go válido. Ainda assim, nesse tipo de `switch` o mais comum é escrever a condição com `||` (`case score < 0 || score > 100:`), porque se lê direto como "menor que 0 ou maior que 100". Guarde a vírgula para o `switch` com expressão, como no `case "saturday", "sunday":`.

**[idiomático] Um valor comparado com várias opções: `switch month`.** O seu `DaysInMonth` usa `switch {` com `case month == 4, month == 6, ...`. Funciona, mas é o caso clássico do `switch month { case 4, 6, 9, 11: }`, como o "Quando usar cada um" do README dizia: menos repetição e a intenção fica clara. E vale olhar o seu `default`:
- **Hoje** ele devolve 31, e por isso você precisou de um `case month <= 0, month > 12` só para os inválidos.
- **Com os sete meses de 31 dias listados num `case`,** o `default` vira naturalmente "inválido", e o caso extra some.

Os dois funcionam. O segundo tem uma vantagem: um mês que você esqueceu de listar cai no 0, que dá para notar, e não no 31, que passaria despercebido.

**[idiomático] Verbos nos comentários.** O nome no começo está certo em todos. Sobram os verbos e o inglês:
- **"reports" é para função que devolve `bool`** ("reports whether"). Para `string` ou `int`, use "returns": `// DayType returns "weekend", "weekday" or "invalid" ...`.
- **`isLeapYear is used to tell if a year is leap or not`** fica `// isLeapYear reports whether year is a leap year.`
- **Inglês:** "specif" → "specific"; "based **in** score" → "based **on** the score".

## Solução de referência

```go
// Package calendar answers questions about days, months and grades.
package calendar

// DayType returns "weekend", "weekday" or "invalid" for a lowercase day name.
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

// Grade returns the letter grade ("A" to "F") for a score from 0 to 100,
// or "invalid" for any other score.
func Grade(score int) string {
	switch {
	case score < 0 || score > 100:
		return "invalid"
	case score >= 90: // scores above 100 were already caught
		return "A"
	case score >= 80: // only reached when score < 90
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

// DaysInMonth returns how many days month has in year,
// or 0 if month is not between 1 and 12.
func DaysInMonth(month, year int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 0
	}
}

// isLeapYear reports whether year is a leap year.
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
```

Por que assim:

- **Cada `switch` escolhido pelo tipo de pergunta:** um valor contra opções fixas (`day`, `month`) usa expressão; faixas (`score`) usam o `switch` sem expressão.
- **Os comentários dentro do `Grade`** dizem o que a ordem garante. É a mesma ideia do "From here on, n is positive" do review 006.
- **O `default` é sempre o "resto".** Em `DayType` e `DaysInMonth`, o resto é inválido; em `Grade`, o inválido já foi pego no primeiro `case`, e o resto é `"F"`.

## Para levar

Em Go, cada comparação é uma operação separada: o `90 <= score <= 100` do Python vira `score >= 90 && score <= 100`. E num `switch` (ou numa escada de `if`), a ordem faz metade do trabalho: com a faixa mais alta primeiro, cada caso só precisa do limite de baixo.
