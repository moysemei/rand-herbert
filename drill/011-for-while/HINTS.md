# Dicas — Drill 011

> **Regra:** antes de abrir uma dica, faça a tabela de rastreio e escreva no `digitloop.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat.

## SumDigits

<details>
<summary>Nível 1 — direção</summary>

Olhe a tabela do 3856 no README: a coluna "o último" passa por todos os dígitos, um por volta. O que você faz com cada um deles?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Um `total` começando em 0.
- Enquanto `n` for maior que 0: some o último dígito ao `total`, e depois tire o último dígito do `n`.
- Quando o `n` chega a 0, o loop para sozinho. Devolva o `total`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for n > 0 {
	total += n % 10
	n /= 10
}
```

</details>

## CountDigits

<details>
<summary>Nível 1 — direção</summary>

É o mesmo loop do `SumDigits`, mas aqui não interessa o valor dos dígitos, só **quantas voltas** o loop dá. E faça o `CountDigits(0)` na cabeça: quantas voltas o `for n > 0` dá com `n = 0`?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Um contador começando em 0.
- Enquanto `n` for maior que 0: tire o último dígito e conte 1.
- **Antes do loop**, trate o zero com retorno antecipado: ele tem 1 dígito, mas o loop não rodaria nenhuma vez.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
if n == 0 {
	return 1
}
count := 0
for n > 0 {
	n /= 10
	count++
}
```

</details>

## ReverseNumber

<details>
<summary>Nível 1 — direção</summary>

Na tabela de 472, o `rev` tem que passar por 2, depois 27, depois 274. Como você sai de 2 para 27 usando o 7? E de 27 para 274 usando o 4? É a mesma conta nas duas vezes.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

A cada volta:
1. Pegue o último dígito do `n`.
2. "Abra espaço" à direita do `rev`, multiplicando por 10 (27 vira 270).
3. Some o dígito nesse espaço (270 + 4 = 274).
4. Tire o último dígito do `n`.

O `rev` começa em 0.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for n > 0 {
	rev = rev*10 + n%10
	n /= 10
}
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| O `go test` fica parado, ou `panic: test timed out` | Loop infinito: nada dentro do loop encolhe o `n`. Aperte Ctrl+C e confira se tem `n /= 10`. |
| `CountDigits(0) = 0, want 1` | Com `n = 0`, o loop não roda. Trate o zero antes do loop. |
| `SumDigits(472) = 11, want 13` | Você tirou o último dígito **antes** de somar ele, e o 2 ficou de fora. Primeiro pegue, depois tire. |
| `ReverseNumber(472) = 13, want 274` | Você somou os dígitos sem abrir espaço. Multiplique o `rev` por 10 antes de somar. |
| `ReverseNumber(472) = 2740, want 274` | Multiplicou por 10 **depois** de somar: sobra um zero no fim. É `rev*10 + digit`, nessa ordem. |
| `ReverseNumber(472) = 130, want 274` | O ×10 foi no dígito, não no `rev`. Quem abre espaço é o `rev`. |
| `declared and not used: ...` | Você criou uma variável e não usou. Confira se ela entrou na conta. |
