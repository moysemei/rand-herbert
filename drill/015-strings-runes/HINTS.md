# Dicas — Drill 015

> **Regra:** antes de abrir uma dica, faça a tabela de rastreio e escreva no `runes.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat.

## Refazer de memória (ReverseNumber)

Sem dicas aqui, porque o objetivo é ver o que ficou. Se travar 10 minutos: anote o que tentou e abra o `REVIEW.md` do drill 011. Depois feche e reescreva sem olhar.

## CountVowels

<details>
<summary>Nível 1 — direção</summary>

É um contador com filtro (drill 010), só que agora o loop percorre as runes de uma string. Qual `for` faz isso?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Um contador em 0.
- `for _, r := range s`: para cada rune, pergunte se ela é uma das cinco vogais.
- Um `switch r` com as cinco num `case` só, separadas por vírgula (drill 007), fica mais limpo do que cinco `||`.
- As vogais vão entre **aspas simples**.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for _, r := range s {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		count++
	}
}
```

</details>

## Sizes

<details>
<summary>Nível 1 — direção</summary>

Um dos dois números o Go já te dá de graça. O outro você conta.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Bytes: `len(s)`.
- Runes: conte as voltas de um `for range s`. Você nem precisa usar a rune: só conta 1 por volta.
- Devolva os dois na ordem da assinatura: bytes primeiro.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
count := 0
for range s {
	count++
}
return len(s), count
```

</details>

## Reverse

<details>
<summary>Nível 1 — direção</summary>

Na tabela do `Reverse("ação")`, o `result` tem que passar por `"a"`, `"ça"`, `"ãça"` e `"oãça"`. Onde cada rune nova entra: na frente ou atrás do que já existe?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Comece com uma string vazia.
- Percorra as runes com `for range`, na ordem normal (do começo para o fim).
- Cada rune nova entra **na frente** do resultado: converta para string com `string(r)` e coloque antes do `result`.

É o mesmo "remontar" do `ReverseNumber`: lá, cada dígito novo entra no fim com `*10 + digit`; aqui, cada rune nova entra no começo.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
result := ""
for _, r := range s {
	result = string(r) + result
}
return result
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `cannot convert "a" (untyped string constant) to type rune` | A vogal está entre aspas duplas. Rune é com aspas simples: `'a'`. |
| `invalid operation: r + result (mismatched types rune and string)` | Rune e string não se somam. Converta a rune antes: `string(r) + result`. |
| `conversion from int to string yields a string of one rune, not a string of digits` | Você escreveu `for r := range s`: com uma variável só, ela é a **posição** (um `int`), não a rune. Use `for _, r := range s`. |
| `Reverse("hello") = "hello", want "olleh"` | A rune entrou atrás (`result += string(r)`). Para inverter, ela entra na frente. |
| `Reverse("ação") = "o£Ã§Ãa", want "oãça"` | Você inverteu **bytes** (com `s[i]`), e os caracteres de 2 bytes foram partidos ao meio. Percorra runes com `for range`. |
| `Sizes("ação") = 6, 6, want 6, 4` | Você usou `len(s)` para os dois. `len` conta bytes; runes você conta com `for range`. |
