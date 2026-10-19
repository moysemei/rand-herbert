# Dicas — Drill 009

> **Regra:** antes de abrir uma dica, escreva no `loops.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**.

## SumTo

<details>
<summary>Nível 1 — direção</summary>

É o acumulador do README. Quais números o `i` precisa visitar: de onde até onde?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Um `total` começando em 0, antes do loop. O `i` vai de 1 até `n`, **incluindo** o `n`. Em cada volta, some o `i` ao `total`. Depois do loop, devolva o `total`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`for i := 1; i <= n; i++ { total += i }`

</details>

## Factorial

<details>
<summary>Nível 1 — direção</summary>

Mesmo formato do `SumTo`, trocando a soma por multiplicação. Qual valor inicial não estraga uma multiplicação?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

O resultado começa em **1** (multiplicar por 1 não muda nada). O `i` vai de 1 (ou de 2, já que multiplicar por 1 não faz diferença) até `n`, e em cada volta o resultado é multiplicado pelo `i`. Com `n = 0`, o loop não roda nenhuma vez, e o 1 inicial já é a resposta certa.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`result := 1`, depois `for i := 2; i <= n; i++ { result *= i }`. O `*=` é o `+=` da multiplicação.

</details>

## SumOfMultiples

<details>
<summary>Nível 1 — direção</summary>

São três coisas que você já sabe juntas: um acumulador, um `if` dentro do loop, e "é múltiplo de" com `%`. E repare: **menores** que `limit`.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- O `i` vai de 1 até `limit - 1`, ou seja, com `i < limit`.
- Dentro do loop, um `if` pergunta "é múltiplo de 3 **ou** de 5?". Só se for, some ao total.
- Uma condição só, com `||`: assim o 15 entra uma vez.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for i := 1; i < limit; i++ {
	if i%3 == 0 || i%5 == 0 {
		total += i
	}
}
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `SumTo(5) = 10, want 15` | O loop parou um antes: `i < n` no lugar de `i <= n`. |
| `Factorial(5) = 0, want 120` | O resultado começou em 0, e qualquer coisa vezes 0 é 0. Comece em 1. |
| `undefined: total` | O `total` foi criado dentro do `for`, e depois do `}` ele não existe mais. Crie antes do loop. |
| `declared and not used: total` | Você usou `total := total + i` dentro do loop, o que cria um `total` novo a cada volta. Use `=` ou `+=`. |
| `SumOfMultiples(10) = 33, want 23` | O 10 entrou na soma, mas é para os **menores** que `limit`: `i < limit`. |
| `SumOfMultiples(16) = 75, want 60` | O 15 foi somado duas vezes (um `if` para o 3 e outro para o 5). Junte as duas perguntas num `if` só, com `\|\|`. |
| `SumOfMultiples(10) = 0, want 23` | Você usou `&&`: só entrariam os múltiplos de 3 **e** de 5 ao mesmo tempo. É `\|\|`. |
| `syntax error: unexpected :=, expected )` | Parênteses em volta das três partes do `for`. Em Go, sem parênteses: `for i := 0; i < n; i++ {`. |
| `syntax error: unexpected ++ at end of statement` | `i++` usado como valor (`x = i++`). Em Go, `i++` fica sozinho na linha. |
