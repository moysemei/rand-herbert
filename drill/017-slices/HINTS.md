# Dicas — Drill 017

> **Regra:** antes de abrir uma dica, faça a tabela de rastreio e escreva no `slices.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat.

## Refazer de memória (DigitalRoot)

Sem dicas aqui. Se travar 10 minutos: anote o que tentou e abra o `REVIEW.md` do drill 013. Depois feche e reescreva sem olhar.

## Sum

<details>
<summary>Nível 1 — direção</summary>

É o acumulador de sempre, só que agora o `for range` percorre um slice. O que ele te dá a cada volta?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Um `total` em 0. `for _, v := range nums`: some o `v` (o elemento da vez) ao total.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for _, v := range nums {
	total += v
}
```

</details>

## MaxOf

<details>
<summary>Nível 1 — direção</summary>

É o `MinMaxDigit` do drill 014, com um acumulador só. Qual valor inicial funciona para `[-3 -1 -7]`? E o que fazer antes, se o slice estiver vazio?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Retorno antecipado: slice vazio (`len(nums) == 0`) → `0, false`.
- O "maior até agora" começa com um elemento de verdade: `nums[0]`.
- Percorra o slice; se o elemento da vez for maior que o maior até agora, ele vira o maior.
- No fim, devolva o maior e `true`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
if len(nums) == 0 {
	return 0, false
}
largest := nums[0]
for _, v := range nums {
	if v > largest {
		largest = v
	}
}
return largest, true
```

</details>

## Countdown

<details>
<summary>Nível 1 — direção</summary>

Você vai montar um slice do zero, um elemento por volta. Qual função adiciona um elemento no fim? E em que ordem o `i` precisa andar?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Comece com um slice vazio: `var result []int`.
- Um `for` com o `i` indo de `n` até 1, descendo (`i--`, como no fatorial contando para baixo do drill 009).
- Em cada volta, `result = append(result, i)`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
var result []int
for i := n; i >= 1; i-- {
	result = append(result, i)
}
return result
```

</details>

## ReverseInPlace

<details>
<summary>Nível 1 — direção</summary>

Na tabela de `[1 2 3 4]`: quem troca de lugar com quem? O primeiro com o último, o segundo com o penúltimo... Quantas trocas são, e o que acontece se você continuar trocando depois do meio?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Para cada `i` da primeira metade (`i < len(nums)/2`), a outra ponta é `j := len(nums) - 1 - i`.
- Troque `nums[i]` e `nums[j]` com a atribuição dupla do README.
- Com 4 elementos são 2 trocas; com 3, só 1 (o do meio fica onde está).

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for i := 0; i < len(nums)/2; i++ {
	j := len(nums) - 1 - i
	nums[i], nums[j] = nums[j], nums[i]
}
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `append(...) (value of type []int) is not used` | O `append` devolve o slice novo, e você não guardou. É `result = append(result, i)`. |
| `invalid operation: ... (slice can only be compared to nil)` | Slices não se comparam com `==`. Compare elemento por elemento, ou use `slices.Equal`. |
| `panic: runtime error: index out of range [0] with length 0` | Você acessou `nums[0]` num slice vazio. Confira `len(nums) == 0` antes. |
| `MaxOf([-3 -1 -7]) = 0, true, want -1, true` | O maior começou em 0. Comece com `nums[0]`. |
| `Countdown(3) = [3 2], want [3 2 1]` | O loop parou antes do 1: `i >= 1`, não `i > 1`. |
| `[1 2 3 4] = [1 2 3 4], want [4 3 2 1]` | Você trocou até o fim, e o slice foi invertido duas vezes. Pare no meio. (Ou: você criou um slice novo dentro da função, e quem chamou não vê. Troque os elementos do próprio `nums`.) |
| `[1 2 3 4] = [4 3 3 4], want [4 3 2 1]` | Você trocou em duas linhas separadas, e a segunda copiou o valor que a primeira já tinha mudado. Use a atribuição dupla: `nums[i], nums[j] = nums[j], nums[i]`. |
