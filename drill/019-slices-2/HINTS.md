# Dicas — Drill 019

> **Regra:** antes de abrir uma dica, confira se o seu plano tem exemplo, padrão e passos em português para a função em que travou. Se faltar, comece por aí. Depois abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat.

## Refazer de memória (SumNumbers)

Sem dicas aqui. Se travar 10 minutos: anote o que tentou e abra o `REVIEW.md` do drill 016. Depois feche e reescreva sem olhar.

## Evens

<details>
<summary>Nível 1 — direção</summary>

"Qual é a nova lista?" com um filtro. É o `Without` do README, com outra condição.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

**Montar com filtro:** comece com `var result []int`; para cada `v`, se ele for par (drill 004), faça o `append`. Não mexa no `nums`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
var result []int
for _, v := range nums {
	if v%2 == 0 {
		result = append(result, v)
	}
}
return result
```

</details>

## FirstN

<details>
<summary>Nível 1 — direção</summary>

São três decisões: qual fatia pega os primeiros `n`; o que fazer quando `n` é maior que o slice; e como garantir que o resultado seja uma cópia. O `Skip` do README faz as três coisas, só que pela outra ponta.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Os primeiros `n` são `nums[:n]`.
- Antes de fatiar, se `n` for maior que `len(nums)`, troque o `n` por `len(nums)` (senão: panic).
- Devolva uma cópia da fatia com `slices.Clone` (`import "slices"`).

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
if n > len(nums) {
	n = len(nums)
}
return slices.Clone(nums[:n])
```

</details>

## RemoveDuplicates

<details>
<summary>Nível 1 — direção</summary>

São dois padrões juntos: você **monta** uma lista nova e, para cada número, **busca** uma resposta. Que pergunta de sim ou não decide se o número entra? E ela é feita sobre qual slice?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Escreva o `contains` (o `Contains` do drill 018).
- **Montar:** `var result []int`.
- Para cada `v` de `nums`: se o `result` **ainda não** contém o `v`, faça o `append`.
- A pergunta é sobre o `result` (o que você já guardou), não sobre o `nums`: o `nums` sempre contém o `v`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
var result []int
for _, v := range nums {
	if !contains(result, v) {
		result = append(result, v)
	}
}
return result
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `after changing FirstN's result, the original changed too` | Você devolveu `nums[:n]`, que é uma janela para os mesmos elementos. Devolva uma cópia: `slices.Clone(nums[:n])`. |
| `panic: runtime error: slice bounds out of range [:10] with capacity 4` | Você fatiou além do fim. Limite o `n` ao `len(nums)` antes. |
| `FirstN([1 2 3 4], 2) = [1], want [1 2]` | O `b` de `s[a:b]` já é exclusivo: os primeiros 2 são `nums[:2]`, não `nums[:1]`. |
| `RemoveDuplicates([1 2 1 3 2]) = [], want [1 2 3]` | Você perguntou se o `nums` contém o `v` (sempre contém), ou esqueceu o `!`. A pergunta é "o resultado **ainda não** tem esse número?". |
| `Evens changed its input` | O resultado foi montado em cima do próprio `nums` (por exemplo, com `nums[:0]`). Comece de um slice novo: `var result []int`. |
| `cannot use b (variable of type []int) as int value in argument to append` | Para juntar dois slices: `append(a, b...)`. |
