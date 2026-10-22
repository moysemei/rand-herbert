# Dicas — Drill 012

> **Regra:** antes de abrir uma dica, faça a tabela de rastreio e escreva no `rules.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat.

## FizzBuzz

<details>
<summary>Nível 1 — direção</summary>

São quatro casos, e o 15 cai em dois deles ao mesmo tempo. Qual caso precisa vir primeiro? (O `switch` sem expressão do drill 007 serve bem aqui.)

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Primeiro "múltiplo de 3 **e** de 5", depois "múltiplo de 3", depois "múltiplo de 5".
- O que sobra vira texto com `strconv.Itoa(n)`, e não com `string(n)`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
switch {
case n%3 == 0 && n%5 == 0:
	return "FizzBuzz"
// ... the 3 case, then the 5 case
default:
	return strconv.Itoa(n)
}
```

</details>

## IsPrime

<details>
<summary>Nível 1 — direção</summary>

É o padrão busca do README: você procura **um divisor**. Quais números são candidatos a divisor de 21? Se achar um, o que já dá para responder? E se o loop acabar sem achar nenhum?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Antes do loop: números menores que 2 não são primos (retorno antecipado).
- Os candidatos a divisor vão de 2 até `n - 1` (ou, mais rápido, enquanto `d*d <= n`).
- Dentro do loop: se `d` divide `n`, já dá para devolver `false`. Não devolva nada quando ele **não** divide: ainda faltam candidatos.
- Depois do loop: nenhum candidato dividiu, então devolva `true`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
if n < 2 {
	return false
}
for d := 2; d < n; d++ {
	if n%d == 0 {
		return false
	}
}
return true
```

</details>

## NextPrime

<details>
<summary>Nível 1 — direção</summary>

Você não sabe quantas voltas vai precisar: é o `for` sem condição (ou com condição) do drill 011. Qual é o primeiro candidato: `n` ou `n + 1`?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Comece do número seguinte ao `n` e vá subindo de 1 em 1. Em cada candidato, pergunte ao seu `IsPrime`. No primeiro que for primo, devolva ele.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for {
	n++
	if IsPrime(n) {
		return n
	}
}
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `FizzBuzz(15) = "Fizz", want "FizzBuzz"` | O caso "3 e 5" veio depois do caso "3". O mais específico vai primeiro. |
| `conversion from int to string yields a string of one rune, not a string of digits` | Você usou `string(n)`. Número vira texto com `strconv.Itoa(n)`. |
| `IsPrime(9) = true, want false` | Tem um `return true` dentro do loop: ele desiste no primeiro candidato que não divide. O "não achei" vai depois do loop. |
| `IsPrime(1) = true, want false` | Faltou o caso pequeno: menores que 2 não são primos. |
| `IsPrime(2) = false, want true` | O loop testou o 1 ou o próprio `n` como divisor, e todo número é divisível por esses dois. Os candidatos começam em 2 e não incluem o `n`. |
| `NextPrime(13) = 13, want 17` | A busca começou no próprio `n`. É o primo **maior** que `n`: comece em `n + 1`. |
| `missing return` | A função pode chegar ao fim sem `return`. Num `for {}` com `break`, falta o `return` depois do loop (ou use `return` direto dentro dele). |
| O `go test` fica parado | Loop infinito. Confira se o candidato muda a cada volta. Ctrl+C para sair. |
