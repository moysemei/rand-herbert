# Dicas — Drill 007

> **Regra:** antes de abrir uma dica, escreva no `calendar.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**.

## DayType

<details>
<summary>Nível 1 — direção</summary>

Você está comparando **um valor** (`day`) com várias opções. Qual dos dois tipos de `switch` serve para isso?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

`switch day` com três caminhos: um `case` com os dois dias do fim de semana separados por vírgula, um `case` com os cinco dias úteis, e o `default` para o resto.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
switch day {
case "saturday", "sunday":
	return "weekend"
// one case with the five weekdays
default:
	return "invalid"
}
```

</details>

## Grade

<details>
<summary>Nível 1 — direção</summary>

Teste na mão com 95: se o primeiro `case` fosse `score >= 80`, o que sairia? E com 101: qual `case` pega primeiro?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Primeiro, o `"invalid"`: menor que 0 **ou** maior que 100.
- Depois, as faixas da nota mais alta para a mais baixa, só com `>=`: 90, 80, 70, 60. Como a de cima já pegou os maiores, cada `case` só precisa do limite de baixo.
- O que sobra é `"F"`, e vai no `default`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```
switch {
case score < 0 || score > 100: "invalid"
case score >= 90:              "A"
case score >= 80:              "B"
... 70 -> "C", 60 -> "D"
default:                       "F"
}
```

</details>

## DaysInMonth

<details>
<summary>Nível 1 — direção</summary>

São quatro grupos de meses: os de 30, o fevereiro, os de 31 e os inválidos. Qual tipo de `switch` compara `month` com valores?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- `switch month` com um `case` para os meses de 30 dias (vírgulas), um para o 2 e um para os de 31 dias.
- O `default` devolve 0: é ele que pega o 13, o 0 e qualquer outro.
- Dentro do `case 2`, um `if` com o `isLeapYear(year)` escolhe entre 29 e 28.
- O `isLeapYear` é uma função separada, com a expressão do `IsLeapYear` do drill 004.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
switch month {
case 4, 6, 9, 11:
	return 30
case 2:
	if isLeapYear(year) {
		return 29
	}
	return 28
// one case with the seven months of 31 days
default:
	return 0
}
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `missing return` | Seu `switch` não tem `default` (ou falta um `return` depois dele). Sem `default`, o Go não tem como saber se algum `case` vai bater. |
| `invalid operation: operator \|\| not defined on 4 (untyped int constant)` | Você escreveu `case 4 \|\| 6`. Para vários valores, use vírgula: `case 4, 6:`. |
| `duplicate case "sunday" (constant of type string) in expression switch` | O mesmo valor aparece em dois `case`. |
| `cannot convert score >= 90 (untyped bool value) to type int` | Você usou condição num `switch score`. Para faixas, use `switch` sem expressão: `switch {`. |
| `undefined: IsLeapYear` | O ajudante se chama `isLeapYear`, com minúscula. Maiúscula e minúscula são nomes diferentes. |
| `syntax error: cannot use assignment (year % 4) = 0 as value` | Um `=` no lugar de `==` dentro da condição. |
| `Grade(95) = "B", want "A"` | As faixas estão de baixo para cima. A primeira condição verdadeira vence: comece pela nota mais alta. |
| `Grade(101) = "A", want "invalid"` | O teste do `"invalid"` veio depois do `score >= 90`. Ele tem que ser o primeiro. |
| `DaysInMonth(2, 1900) = 29, want 28` | O `isLeapYear` está só com `% 4`. Faltam as exceções de 100 e 400. |
