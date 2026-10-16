# Dicas — Drill 008

> **Regra:** antes de abrir uma dica, escreva no `parking.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**.

## FormatDuration

<details>
<summary>Nível 1 — direção</summary>

É o `SecondsToClock` do drill 004, só que com duas partes em vez de três.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Horas: quantas horas inteiras cabem nos minutos.
- Minutos: o que sobra depois de tirar as horas inteiras.

Uma variável para cada, e depois um `Sprintf` com dois `%d`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`minutes / minutesPerHour` e `minutes % minutesPerHour`, com o molde `"%dh %dm"`.

</details>

## BilledHours

<details>
<summary>Nível 1 — direção</summary>

Com 61 minutos, a resposta é 1 ou 2. Com 60, também é 1 ou 2. Quais são os dois candidatos, e o que muda entre 60 e 61 para escolher um ou outro?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- O candidato de baixo são as horas inteiras (divisão inteira por 60).
- O de cima é ele mais 1.
- O que decide é se sobrou algum minuto: se o resto da divisão por 60 for maior que zero, cobre o de cima. Senão, o de baixo.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
fullHours := minutes / minutesPerHour
if minutes%minutesPerHour > 0 {
	return fullHours + 1
}
return fullHours
```

</details>

## Price

<details>
<summary>Nível 1 — direção</summary>

Faça o 125 na mão: quantas horas são cobradas? Quanto custa a primeira? Quantas sobram, e quanto custa cada uma?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Primeiro a tolerância, com retorno antecipado: até 15 minutos, devolva 0.
- Depois pegue as horas cobradas com o `BilledHours`.
- O preço é a primeira hora, mais as horas que sobram (todas menos a primeira) vezes o preço da hora extra.

Com 125: 3 horas cobradas → 8 + (3 - 1) × 5 = 18.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
if minutes <= freeMinutes {
	return 0
}
hours := BilledHours(minutes)
return firstHourPrice + (hours-1)*extraHourPrice
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `BilledHours(60) = 2, want 1` | Você soma 1 sempre. Só sobe quando sobram minutos. |
| `BilledHours(61) = 1, want 2` | Você só pegou as horas inteiras. Falta o caso de subir. |
| `Price(15) = 8, want 0` | "Até 15 minutos" inclui o 15: use `<=`. |
| `Price(16) = 13, want 8` | A primeira hora foi cobrada duas vezes: as horas extras são todas **menos a primeira**. |
| `Price(16) = 5, want 8` | Você cobrou todas as horas pelo preço da hora extra. A primeira tem preço próprio. |
| `Price(16) = 4, want 8` | Faltaram parênteses: `hours-1*extraHourPrice` faz a multiplicação antes, como na matemática. Escreva `(hours-1)*extraHourPrice`. |
| `undefined: billedHours` | O nome tem que ser igual, com a mesma maiúscula: `BilledHours`. |
| `declared and not used: ...` | Você criou uma variável e não usou. Confira se ela entrou na conta ou no `Sprintf`. |
| `"fmt" imported and not used` | Só o `FormatDuration` precisa do `fmt`. Se você ainda não escreveu ele, comente o `import` por enquanto. |
