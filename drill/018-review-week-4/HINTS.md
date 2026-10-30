# Dicas — Drill 018

> **Regra:** antes de abrir uma dica, confira se o seu plano tem os passos 1 a 3 (exemplo, padrão, passos em português) para a função em que travou. Na maioria das vezes, o que falta está ali. Depois abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat.

## Refazer de memória (Reverse)

Sem dicas aqui. Se travar 10 minutos: anote o que tentou e abra o `REVIEW.md` do drill 015. Depois feche e reescreva sem olhar.

## CountChar

<details>
<summary>Nível 1 — direção</summary>

A pergunta é "quantas vezes?". Qual padrão do cardápio responde "quantos são"?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

**Contar.** Um contador em 0; percorra as runes de `s`; quando a rune da vez for igual ao `target`, conte 1.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for _, r := range s {
	if r == target {
		count++
	}
}
```

</details>

## Contains

<details>
<summary>Nível 1 — direção</summary>

A pergunta é "existe algum?". Qual padrão responde isso? Onde fica o "achei" e onde fica o "não achei"?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

**Buscar.** Percorra o slice; se o elemento da vez for o `target`, devolva `true` na hora. Se o loop acabar sem achar, devolva `false`, **depois** do loop.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
for _, v := range nums {
	if v == target {
		return true
	}
}
return false
```

</details>

## LongestWord

<details>
<summary>Nível 1 — direção</summary>

"Qual é a mais longa?" é o padrão **melhor até agora**. São duas decisões: de onde vêm as palavras (drill 016) e qual é o ponto de partida.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- As palavras saem de `strings.Fields(s)`.
- O ponto de partida pode ser `""`: qualquer palavra de verdade é mais longa que ele, e se não houver palavras, `""` já é a resposta.
- Uma palavra só substitui a melhor se for **mais longa** (`>`). Com `>=`, um empate trocaria pela última, e o enunciado pede a primeira.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
best := ""
for _, w := range strings.Fields(s) {
	if len(w) > len(best) {
		best = w
	}
}
return best
```

</details>

## ParseAll

<details>
<summary>Nível 1 — direção</summary>

São dois padrões juntos. A função monta uma lista nova (qual padrão?) e pode desistir no primeiro texto inválido (qual padrão?). É o `SumNumbers` do drill 016 com um `append` no lugar da soma.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- **Montar:** comece com `var result []int`.
- Para cada `word`, converta com `strconv.Atoi(word)`.
- **Buscar:** se `err != nil`, devolva `nil, false` na hora.
- Senão, faça o `append` do número.
- Depois do loop, devolva o resultado e `true`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
var result []int
for _, w := range words {
	n, err := strconv.Atoi(w)
	if err != nil {
		return nil, false
	}
	result = append(result, n)
}
return result, true
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `invalid operation: r == "target" (mismatched types rune and untyped string)` | Você comparou com um texto. O `target` já é uma rune: compare `r == target`, sem aspas. |
| `Contains([4 8 15], 8) = false, want true` | Tem um `return false` dentro do loop: ele desiste no primeiro elemento diferente. O "não achei" vem depois do loop. |
| `LongestWord("the quick brown fox") = "brown"` | Você usou `>=`: no empate, a última palavra venceu. Com `>`, a primeira fica. |
| `panic: runtime error: index out of range [0] with length 0` | Você pegou a primeira palavra de um texto sem palavras. Comece com `""`, ou confira se há palavras antes. |
| `ParseAll(["1" "x" "3"]) = [1 3], true` | Você pulou o erro (`continue`) em vez de desistir. O enunciado pede `nil, false` no primeiro texto inválido. |
| `append(result, n) (value of type []int) is not used` | É `result = append(result, n)`. |
