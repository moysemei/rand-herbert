# Drill 007 — switch

**Hoje:** trocar uma escada de `if` por um `switch`, com e sem expressão.

## O que é

O `switch` compara um valor com vários casos e executa **só o primeiro** que bater:

```go
switch value {
case 1:
	// value == 1
case 2, 3: // several values in one case, separated by commas
	// value == 2 or value == 3
default:
	// none of the cases matched
}
```

Também existe o **`switch` sem expressão**: cada `case` é uma condição `bool`, testada de cima para baixo. É um jeito mais limpo de escrever uma escada de `if / else if`:

```go
switch {
case condition1:
	// first true condition wins
case condition2:
	// ...
default:
	// nothing was true
}
```

## Por que é assim

- **Não precisa de `break`.** Em JS (e em C, Java...), se você esquece o `break`, a execução "cai" no `case` de baixo. É um bug clássico. No Go é o contrário: cada `case` para sozinho. Se um dia você quiser mesmo continuar no próximo, existe a palavra `fallthrough`, mas quase ninguém usa.
- **Vários valores no mesmo `case`** com vírgula: `case "saturday", "sunday":`. Nada de repetir o `case` ou de usar `||`.
- **O Python só ganhou algo parecido na versão 3.10** (`match`). Antes, era tudo `if/elif`. Em Go, o `switch` sem expressão é o jeito idiomático de escrever uma escada de condições.
- **Quando usar cada um:** compare **um valor** com várias opções → `switch value`. Teste **condições diferentes** (faixas, combinações) → `switch` sem expressão. Duas ou três condições simples → `if` mesmo.

## Como se usa

Exemplos de outro assunto (não são o problema de hoje):

```go
package traffic

// Action says what a driver should do at a traffic light.
func Action(color string) string {
	switch color {
	case "green":
		return "go"
	case "yellow", "amber":
		return "slow down"
	case "red":
		return "stop"
	default:
		return "unknown color"
	}
}

// Weather describes the day. In a switch without an expression,
// each case is a condition, and the first true one wins.
func Weather(celsius int, raining bool) string {
	switch {
	case raining && celsius < 10:
		return "cold rain" // must come before the plain "raining" case
	case raining:
		return "rain"
	case celsius > 30:
		return "hot"
	default:
		return "nice"
	}
}
```

**Funções com nome em minúscula** (por exemplo, `toKelvin`) só são visíveis dentro do próprio pacote. Servem para ajudantes internos: quem usa o pacote enxerga só as funções com maiúscula. É a regra do drill 001 (maiúscula = exportado) vista pelo outro lado.

## Armadilhas

- **`missing return` mesmo com todos os `case` retornando.** Sem `default`, o Go não sabe se algum `case` vai bater. Tenha um `default` com `return`, ou um `return` depois do `switch`.
- **`case 4 || 6` não funciona.** Num `switch month`, cada `case` é um **valor** para comparar com `month`, não uma condição. Para vários valores, use vírgula: `case 4, 6:`.
- **Condição num `switch` com expressão.** `switch score { case score >= 90: ... }` não compila: você está comparando um `int` com um `bool`. Para faixas, use o `switch` sem expressão.
- **A ordem importa no `switch` sem expressão,** igual à escada de `if`: o caso mais específico vai primeiro.
- **Texto é comparado letra por letra:** `"Monday"` não é `"monday"`.

## O problema

Crie o arquivo **`calendar.go`** nesta pasta, com **`package calendar`**, e implemente:

1. **`func DayType(day string) string`**: os dias chegam em inglês, em minúsculas.
   - `"saturday"` ou `"sunday"` → `"weekend"`
   - `"monday"` a `"friday"` → `"weekday"`
   - qualquer outro texto → `"invalid"`
   - Use `switch` **com** expressão.
2. **`func Grade(score int) string`**: transforma uma nota de 0 a 100 em conceito.
   - 90 a 100 → `"A"`, 80 a 89 → `"B"`, 70 a 79 → `"C"`, 60 a 69 → `"D"`, 0 a 59 → `"F"`
   - fora de 0 a 100 → `"invalid"`
   - Exemplos: `Grade(95)` → `"A"`, `Grade(89)` → `"B"`, `Grade(59)` → `"F"`, `Grade(101)` → `"invalid"`, `Grade(-1)` → `"invalid"`
   - Use `switch` **sem** expressão.
3. **`func DaysInMonth(month, year int) int`**: quantos dias tem o mês (1 a 12) naquele ano.
   - abril, junho, setembro e novembro (4, 6, 9, 11) → `30`
   - fevereiro (2) → `29` em ano bissexto, `28` nos outros
   - os demais meses → `31`
   - mês fora de 1 a 12 → `0`
   - Exemplos: `DaysInMonth(1, 2023)` → `31`, `DaysInMonth(4, 2023)` → `30`, `DaysInMonth(2, 2023)` → `28`, `DaysInMonth(2, 2024)` → `29`, `DaysInMonth(2, 1900)` → `28`, `DaysInMonth(13, 2023)` → `0`
   - Crie um ajudante **`func isLeapYear(year int) bool`**, com minúscula, e use ele no caso de fevereiro. A regra do bissexto é a do drill 004: pode abrir o seu `numbers.go` para lembrar.

O `Grade` e o `DaysInMonth` são as partes de pensar. **Antes de codar**, faça no `// Plano:`:
- **`Grade`:** em que ordem ficam os `case`? Onde entra o `"invalid"`? Teste na mão com 95 e com 101.
- **`DaysInMonth(2, 1900)`:** qual `case` pega? O que decide entre 28 e 29? E o 13, cai onde?

Se travar, as dicas estão em `HINTS.md`.

## Checklist

- [ ] `// Plano:` com a ordem do `Grade` e o caminho do `DaysInMonth(2, 1900)`.
- [ ] Comentário do pacote e de cada função (o ajudante também), começando pelo nome e com verbo na terceira pessoa (`returns`, `reports whether`).
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  git add .
  git commit -m "drill 007: switch"
  git push
  ```
