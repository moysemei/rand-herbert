# Drill 008 — Revisão da semana: o estacionamento

**Hoje:** sem conceito novo. Você vai juntar a semana (operadores, `/` e `%`, `if`, `switch`) num problema só: quanto um estacionamento cobra. E uma função vai usar a outra.

## O que é

Revisão do que você já sabe:

- **`/` responde "quantos inteiros cabem?"; `%` responde "o que sobra?".** (drills 004 e 005)
- **Calcular primeiro, formatar depois.** O `Sprintf` só coloca valores prontos nos slots. (drill 004)
- **Do exemplo para o código.** Faça a conta com números e pergunte de onde veio cada número; o nome da variável sai daí. (drill 005)
- **Escolher entre candidatos.** Liste as respostas possíveis, ache a conta de cada uma e só então escreva o `if`. (drill 006)
- **A primeira condição verdadeira vence,** tanto na escada de `if` quanto no `switch`. O caso mais específico vai primeiro. (drills 006 e 007)
- **Retorno antecipado:** trate o caso especial, `return`, e siga sem `else`. (drill 006)
- **Constante com a unidade no nome:** `minutesPerHour`, não `oneHour`. (drill 004)
- **Uma função pode usar outra,** em vez de repetir a regra. (drill 003)

## Por que é assim

"Toda hora começada conta como hora cheia" é um arredondamento **para cima**, e ele aparece o tempo todo em backend:
- **Paginação:** 101 itens com 10 por página pedem 11 páginas, não 10.
- **Cobrança por hora ou por minuto iniciado.**
- **Quantas máquinas** são necessárias para atender uma carga.

Em Python, o reflexo seria `math.ceil(minutes / 60)`, passando por `float`. Em Go, com `int`, você resolve sem sair dos inteiros: a divisão inteira dá o candidato de baixo, e o `%` diz se sobrou alguma coisa. É o mesmo raciocínio do `RoundToTen`, com outra regra para escolher.

## Como se usa

Exemplo de outro assunto, para lembrar constantes com unidade e uma função usando a outra:

```go
package gym

const (
	monthlyFee     = 120 // reais per month
	discountMonths = 12  // plans this long or longer get a discount
	discountFee    = 100 // reais per month with the discount
)

// FeePerMonth returns the monthly fee for a plan of the given length.
func FeePerMonth(months int) int {
	if months >= discountMonths {
		return discountFee
	}
	return monthlyFee
}

// PlanTotal returns the total price of a plan. It reuses FeePerMonth
// instead of repeating the discount rule.
func PlanTotal(months int) int {
	return months * FeePerMonth(months)
}
```

## Armadilhas

Os "pontos de atenção" que apareceram nos reviews:

- **Os limites contam.** "Até 15 minutos" inclui o 15: `<=`, não `<`. E 60 minutos é **uma** hora cheia, não duas.
- **Confira o intervalo:** a parte dos minutos no texto fica sempre entre 0 e 59.
- **Faça o exemplo à mão antes de codar,** e depois troque cada número pelo nome de onde ele veio.
- **Comentários com verbo na terceira pessoa:** `returns` para quem devolve número ou texto; `reports whether` para quem devolve `bool`.
- **Não repita a regra.** Se o `Price` refizer a conta das horas, o `BilledHours` perde o sentido.

## O problema

Crie o arquivo **`parking.go`** nesta pasta, com **`package parking`**. As regras do estacionamento:

- os primeiros **15 minutos** são grátis (a tolerância);
- passou disso, a **primeira hora** custa **R$ 8** e cada **hora adicional** custa **R$ 5**;
- toda hora começada conta como hora cheia.

Use constantes, com a unidade no nome, para o 60, o 15, o 8 e o 5.

1. **`func FormatDuration(minutes int) string`**: o tempo em horas e minutos.
   - `FormatDuration(125)` → `"2h 5m"`
   - `FormatDuration(59)` → `"0h 59m"`
   - `FormatDuration(60)` → `"1h 0m"`
2. **`func BilledHours(minutes int) int`**: quantas horas são cobradas. Qualquer minuto a mais começa uma nova hora.
   - `BilledHours(0)` → `0`
   - `BilledHours(1)` → `1`
   - `BilledHours(60)` → `1`
   - `BilledHours(61)` → `2`
   - `BilledHours(125)` → `3`
3. **`func Price(minutes int) int`**: o preço em reais.
   - `Price(10)` → `0` e `Price(15)` → `0` (tolerância)
   - `Price(16)` → `8` e `Price(60)` → `8` (só a primeira hora)
   - `Price(61)` → `13` (8 + 5)
   - `Price(125)` → `18` (3 horas cobradas: 8 + 5 + 5)
   - **Regra:** use o `BilledHours` aqui dentro.

O `BilledHours` e o `Price` são as partes de pensar. **Antes de codar**, faça no `// Plano:`:
- **`BilledHours`, com 60 e com 61:** quais são os dois candidatos em cada caso? O que escolhe entre eles?
- **`Price(125)`:** a conta com números (125 → quantas horas → quanto custa cada uma). Depois troque cada número pelo nome de onde ele veio.

Se travar, as dicas estão em `HINTS.md`.

## Checklist

- [ ] `// Plano:` com os candidatos do `BilledHours` (60 e 61) e a conta do `Price(125)` com números e depois com nomes.
- [ ] Comentário do pacote e de cada função, com verbo na terceira pessoa.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  git add .
  git commit -m "drill 008: week 2 review"
  git push
  ```
