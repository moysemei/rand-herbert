# Dicas — Drill 016

> **Regra:** antes de abrir uma dica, faça a tabela de rastreio e escreva no `text.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat. Para ver o que uma função faz: `go doc strings.ReplaceAll` no terminal.

## Refazer de memória (FizzBuzz)

Sem dicas aqui. Se travar 10 minutos: anote o que tentou e abra o `REVIEW.md` do drill 012. Depois feche e reescreva sem olhar.

## Normalize

<details>
<summary>Nível 1 — direção</summary>

São duas transformações: minúsculas, e tirar os espaços. Procure na tabela do README a função de cada uma.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Minúsculas: `strings.ToLower`.
- Tirar **todos** os espaços (não só os das pontas): `strings.ReplaceAll`, trocando `" "` por `""`.
- Uma função pode receber o resultado da outra direto, sem variável no meio.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`return strings.ReplaceAll(strings.ToLower(s), " ", "")`

</details>

## IsPalindrome

<details>
<summary>Nível 1 — direção</summary>

É o `IsPalindromeNumber` do drill 013, só que com texto: compare o texto com ele mesmo invertido. Mas qual texto: o original ou o normalizado?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Normalize o texto uma vez e guarde numa variável.
- Escreva o ajudante `reverse` (o `Reverse` do drill 015: cada rune entra na frente).
- Compare o normalizado com o normalizado invertido. A comparação já é o `bool` da resposta.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
clean := Normalize(s)
return clean == reverse(clean)
```

</details>

## SumNumbers

<details>
<summary>Nível 1 — direção</summary>

São três coisas que você já sabe juntas: separar as palavras (`strings.Fields`), converter cada uma (`strconv.Atoi`) e somar (acumulador). E o padrão busca do drill 012: no primeiro pedaço inválido, a resposta já está decidida.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Um `total` em 0.
- `for _, word := range strings.Fields(s)`: converta cada `word` com `strconv.Atoi`.
- Se `err != nil`, devolva `0, false` na hora (retorno antecipado dentro do loop).
- Senão, some ao total.
- Depois do loop, todas as palavras eram números: devolva o total e `true`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
total := 0
for _, word := range strings.Fields(s) {
	n, err := strconv.Atoi(word)
	if err != nil {
		return 0, false
	}
	total += n
}
return total, true
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `undefined: strconv` (ou `strings`) | Faltou o import. Com dois pacotes, use `import ( ... )` com os dois. |
| `"strconv" imported and not used` | Você importou e ainda não usou. Normal enquanto o `SumNumbers` não está pronto. |
| `Normalize("A Man A Plan") = "AManAPlan"` | O resultado do `strings.ToLower` foi jogado fora. Ele devolve uma string nova: guarde ou use direto. |
| `Normalize("A Man A Plan") = "a man a plan"` | O `TrimSpace` só tira os espaços das pontas. Para tirar todos, `strings.ReplaceAll`. |
| `IsPalindrome("Ana") = false` | Você comparou sem normalizar: `"Ana"` invertido é `"anA"`. |
| `SumNumbers("4 x 5") = 9, true` | O erro do `Atoi` foi ignorado. Confira `err != nil` e devolva `0, false`. |
| `SumNumbers("1 2 3") = 1, true` | O `return total, true` ficou dentro do loop, e a função parou na primeira palavra. Ele vai depois do loop. |
| `SumNumbers(" 10   20 ") = 0, false` | Você separou com `strings.Split(s, " ")`, que cria pedaços vazios entre espaços repetidos. O `strings.Fields` ignora espaços sobrando. |
