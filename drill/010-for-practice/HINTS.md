# Dicas — Drill 010

> **Regra:** antes de abrir uma dica, escreva no `practice.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat antes de procurar a resposta pronta.

## SumEven

<details>
<summary>Nível 1 — direção</summary>

É o formato do `SumOfMultiples` do drill 009, com outra pergunta dentro do `if`. Que pergunta diz se um número é par? (drill 004)

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Um `total` começando em 0, antes do loop.
- O `i` vai de 1 até `n`, **incluindo** o `n`.
- Dentro do loop, um `if` deixa passar só os pares, e eles são somados ao `total`.

(Outro jeito, sem `if`: começar o `i` em 2 e andar de 2 em 2, com `i += 2`.)

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for i := 1; i <= n; i++ {
	if i%2 == 0 {
		total += i
	}
}
```

</details>

## CountMultiples

<details>
<summary>Nível 1 — direção</summary>

Aqui não interessa **o valor** dos números que passam no filtro, só **quantos** passaram. Na sua tabela de `CountMultiples(10, 3)`, a coluna `count` vai 0, 0, 1, 1, 1, 2... O que acontece com ela quando um número passa?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Um contador começando em 0.
- O `i` vai de 1 até `limit`, incluindo o `limit`.
- Se `i` for múltiplo de `k`, o contador aumenta **1**, não o valor de `i`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
if i%k == 0 {
	count++
}
```

</details>

## Power

<details>
<summary>Nível 1 — direção</summary>

2⁴ = 2 × 2 × 2 × 2. Quantas vezes o 2 aparece? De qual parâmetro vem esse número? E de qual parâmetro vem o 2?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- O resultado começa em 1.
- O loop roda `exp` vezes.
- Em cada volta, o resultado é multiplicado pela `base`.

O `i` só conta as voltas e **não entra na conta**, como no `Repeat` do drill 009.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
result := 1
for i := 0; i < exp; i++ {
	result *= base
}
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `SumEven(6) = 6, want 12` | O loop parou antes do `n`: use `i <= n`. |
| `CountMultiples(10, 3) = 18, want 3` | Você somou os múltiplos (3 + 6 + 9) em vez de contar quantos são. Use `count++`. |
| `CountMultiples(10, 5) = 1, want 2` | O `limit` ficou de fora: "de 1 até `limit`" é `i <= limit`. |
| `CountMultiples(10, 3) = 2, want 3` | A ordem do `%` está invertida: é `i % k` ("o `i` dividido por `k` deixa resto?"), não `k % i`. |
| `Power(2, 10) = 0, want 1024` | O resultado começou em 0. Para multiplicar, comece em 1. |
| `Power(2, 10) = 3628800, want 1024` | Você multiplicou pelo `i` (isso é o fatorial). Em cada volta, multiplique pela `base`. |
| `Power(2, 10) = 2048, want 1024` | Uma volta a mais. Começando em 0, o loop de `exp` voltas é `i < exp`. |
| `Power(3, 4) = 256, want 81` | Você multiplicou pelo `exp`. Ele só diz quantas voltas; quem entra na conta é a `base`. |
