# Dicas — Drill 013

> **Regra:** os ajudantes são de memória. Tente primeiro sem abrir nada, nem os drills 011 e 012. Se travar, faça a tabela de rastreio, escreva no `puzzles.go` o que tentou, e abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat.

## isPrime e sumDigits (ajudantes)

<details>
<summary>Nível 1 — direção</summary>

- **`isPrime`:** padrão busca. O que você procura? O que já dá para responder quando acha? E quando o loop acaba sem achar?
- **`sumDigits`:** "pega o último, tira o último", até sobrar 0.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- **`isPrime`:** menores que 2 não são primos. Os candidatos a divisor vão de 2 até `d*d <= n`. Se algum dividir, não é primo. Se o loop acabar, é primo.
- **`sumDigits`:** um `total` em 0. Enquanto `n > 0`, some `n % 10` ao total e faça `n /= 10`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

Neste nível, pode abrir os `REVIEW.md` dos drills 011 e 012: a solução de referência está lá. Depois, feche e reescreva sem olhar.

</details>

## CountPrimes

<details>
<summary>Nível 1 — direção</summary>

Essa função conta ou soma? É o formato do `CountMultiples` do drill 010, com outra pergunta dentro do `if`.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Um contador em 0. Passe por todos os números até o `limit` (incluindo ele). Para cada um que o `isPrime` aprovar, conte 1.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for i := 2; i <= limit; i++ {
	if isPrime(i) {
		count++
	}
}
```

</details>

## IsPalindromeNumber

<details>
<summary>Nível 1 — direção</summary>

Um número é palíndromo quando ele é igual ao próprio número invertido. Você inverteu números no drill 011. O que acontece com o `n` durante a inversão, e com o que você compara no fim?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Guarde uma cópia do `n` **antes** do loop (por exemplo, `original := n`).
- Inverta o número com o "abre espaço e encaixa".
- No fim, compare o número invertido com a **cópia**. A comparação já é o `bool` que a função devolve (drill 004).

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
original := n
rev := 0
for n > 0 {
	rev = rev*10 + n%10
	n /= 10
}
return rev == original
```

</details>

## DigitalRoot

<details>
<summary>Nível 1 — direção</summary>

Quantas vezes você soma os dígitos de 9875? E de 7? Você não sabe antes quantas vezes vai ser: que tipo de `for` serve quando a gente não sabe quantas voltas são? (drill 011)

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Enquanto o `n` tiver mais de um dígito (ou seja, `n >= 10`), troque o `n` pela soma dos dígitos dele, usando o `sumDigits`. Quando sair do loop, o `n` já é a resposta.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for n >= 10 {
	n = sumDigits(n)
}
return n
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `IsPalindromeNumber(121) = false, want true` | Você comparou o invertido com o `n` depois do loop, e ele já vale 0. Compare com a cópia. |
| `declared and not used: original` | Você guardou a cópia, mas comparou com o `n`. Use a cópia na comparação. |
| `CountPrimes(10) = 17, want 4` | Você somou os primos (2 + 3 + 5 + 7) em vez de contar. Use `count++`. |
| `CountPrimes(2) = 0, want 1` | O `limit` ficou de fora: "de 1 até `limit`" é `i <= limit`. |
| `DigitalRoot(9875) = 29, want 2` | Você somou os dígitos uma vez só. Repita enquanto tiver mais de um dígito. |
| `DigitalRoot(10) = 10, want 1` | A condição do loop foi `n > 10`, mas 10 tem dois dígitos. Use `n >= 10`. |
| O `go test` fica parado no `DigitalRoot` | O resultado do `sumDigits(n)` não foi guardado no `n`, então o `n` nunca muda. É `n = sumDigits(n)`. Ctrl+C para sair. |
| `undefined: isPrime` | O nome do ajudante tem que ser igual em todos os lugares, com a mesma letra minúscula. |
