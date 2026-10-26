# Dicas — Drill 014

> **Regra:** antes de abrir uma dica, faça a tabela de rastreio e escreva no `multi.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat.

## Refazer de memória (Power)

Sem dicas aqui, porque o objetivo é ver o que ficou. Se travar 10 minutos: anote o que tentou e abra o `REVIEW.md` do drill 010. Depois feche e reescreva sem olhar.

## SplitDigits

<details>
<summary>Nível 1 — direção</summary>

As contas você já fez no `DescribeDigits` do drill 005. O que muda hoje é só como os valores saem da função.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Os nomes `hundreds`, `tens` e `ones` já existem dentro da função (retorno nomeado). Atribua cada um com `=` (não `:=`) e devolva os três na ordem da assinatura.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
hundreds = n / 100
tens = n % 100 / 10
ones = n % 10
return hundreds, tens, ones
```

</details>

## Divide

<details>
<summary>Nível 1 — direção</summary>

Qual entrada não tem resposta? Trate ela primeiro, como o `ParseGrade` do README trata a letra inválida.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Retorno antecipado: se `b` for 0, devolva o valor zero e `false`. Senão, devolva a divisão e `true`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
if b == 0 {
	return 0, false
}
return a / b, true
```

</details>

## MinMaxDigit

<details>
<summary>Nível 1 — direção</summary>

São dois acumuladores andando juntos no loop de dígitos do drill 011: um guarda o menor visto até agora, o outro o maior. A pergunta difícil é com que valor eles começam. Teste o seu chute com `n = 0` e com `n = 7`.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Comece `min` e `max` com um dígito de verdade: o último (`n % 10`). Assim o 0 e os números de um dígito já saem certos, mesmo que o loop não rode.
- No loop "pega o último, tira o último": se o dígito for menor que o `min`, ele vira o `min`; se for maior que o `max`, vira o `max`. São dois `if` separados, não um `if/else`.
- Devolva na ordem da assinatura: primeiro o menor.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
lo, hi := n%10, n%10
for n > 0 {
	d := n % 10
	if d < lo {
		lo = d
	}
	if d > hi {
		hi = d
	}
	n /= 10
}
return lo, hi
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `assignment mismatch: 1 variable but ... returns 3 values` | Você recebeu menos valores do que a função devolve. Receba todos ou use `_`. |
| `not enough return values` | Faltou um valor no `return`: a função promete `(int, bool)` e você devolveu só o `int`. |
| `no new variables on left side of :=` | O nome já existe (retorno nomeado). Use `=`. |
| `Divide(5, 0) = 0, true, want 0, false` | O caso sem resposta devolveu `true`. O `ok` precisa dizer que não deu. |
| `MinMaxDigit(0) = 9, 0, want 0, 0` | O `min` começou em 9 e o loop não rodou para o 0. Comece com um dígito de verdade. |
| `MinMaxDigit(4729) = 0, 9, want 2, 9` | O `min` começou em 0, que é menor que todos os dígitos de 4729. Comece com um dígito de verdade. |
| `MinMaxDigit(4729) = 9, 2, want 2, 9` | Ordem trocada no `return`: a assinatura promete o menor primeiro. |
| `MinMaxDigit(4729) = 4, 9, want 2, 9` | Você usou `if/else`: todo dígito que não é o maior virou o `min`. São duas perguntas independentes, dois `if`. |
