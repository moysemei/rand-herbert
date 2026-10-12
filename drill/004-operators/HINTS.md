# Dicas — Drill 004

> **Regra:** antes de abrir uma dica, escreva no `numbers.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**.

## IsEven

<details>
<summary>Nível 1 — direção</summary>

Um número par dividido por 2 deixa qual resto?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Calcule o resto da divisão por 2 e compare com o valor que só os pares têm. Releia a armadilha dos negativos: compare com o resto dos **pares**, não com o dos ímpares.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`n % 2 == 0`. O resultado da comparação já é o `bool` que a função devolve.

</details>

## SecondsToClock

<details>
<summary>Nível 1 — direção</summary>

Pense em como você faria na mão com 3665: quantas horas inteiras cabem? O que sobra? Desse resto, quantos minutos inteiros cabem? O que sobra?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Horas: divisão inteira do total pelos segundos de uma hora.
- Minutos: pegue o **resto** que sobrou das horas e divida pelos segundos de um minuto.
- Segundos: o resto da divisão do total por 60.

Depois monte o texto com `fmt.Sprintf` e três `%d`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`minutes := total % secondsPerHour / secondsPerMinute`. O `%` e o `/` têm a mesma prioridade e são lidos da esquerda para a direita: primeiro o resto, depois a divisão.

</details>

## IsLeapYear

<details>
<summary>Nível 1 — direção</summary>

São três condições com `%`. Escreva cada uma separada, em português, no plano: "divisível por 4", "não divisível por 100", "divisível por 400". Como você liga elas: com "e" ou com "ou"?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Um ano é bissexto em dois casos: **(divisível por 4 E não divisível por 100)** OU **(divisível por 400)**. "Divisível por X" é `year % X == 0`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`(year%4 == 0 && year%100 != 0) || year%400 == 0`

</details>

## Erros comuns do compilador

| Mensagem | O que significa |
|---|---|
| `syntax error: unexpected = at end of statement` | Você usou `=` (atribuição) onde queria `==` (comparação). |
| `non-boolean condition in if statement` | Você escreveu algo como `if n % 2`. Em Go, condição tem que ser `bool`: compare com `==`. |
| `invalid operation: operator % not defined on f (variable of type float64)` | `%` só funciona com inteiros. |
| `IsEven(-3) = true, want false` | Você comparou com o resto dos ímpares (`!= 1`). Com negativo, o resto é `-1`. Compare com `0`. |
| `SecondsToClock(3665) = "1h 61m 5s"` | Os minutos foram calculados do total, e não do que sobrou depois de tirar as horas. |
| `IsLeapYear(1900) = true, want false` | Faltou a exceção dos anos divisíveis por 100. |
