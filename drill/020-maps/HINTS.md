# Dicas — Drill 020

> **Regra:** antes de abrir uma dica, confira se o seu plano tem o método (operação pronta? padrão? passos em português?) para a função em que travou. Depois abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat.

## Refazer de memória (ReverseInPlace)

Sem dicas aqui. Se travar 10 minutos: anote o que tentou e abra o `REVIEW.md` do drill 017. Depois feche e reescreva sem olhar.

## AddStock

<details>
<summary>Nível 1 — direção</summary>

Operação pronta: é uma linha. O que vale `stock["pear"]` quando a pera ainda não existe?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Como a chave ausente vale 0, somar ao valor atual funciona nos dois casos: item que já existe e item novo. E como o map é passado por referência, mudar o `stock` aqui dentro muda o map de quem chamou.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`stock[item] += qty`

</details>

## WordCount

<details>
<summary>Nível 1 — direção</summary>

**Contar por chave.** E duas operações prontas do drill 016 preparam as palavras. Quais?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Crie o map vazio com `map[string]int{}` (não com `var`, que cria um map `nil`).
- As palavras: `strings.Fields` do texto já em minúsculas (`strings.ToLower`).
- Para cada palavra, `counts[word]++`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
counts := map[string]int{}
for _, w := range strings.Fields(strings.ToLower(s)) {
	counts[w]++
}
return counts
```

</details>

## Unique

<details>
<summary>Nível 1 — direção</summary>

É o `RemoveDuplicates` do drill 019, com uma troca: em vez de perguntar ao `contains` "o resultado já tem esse número?", você pergunta a um map **"já vi esse número?"**. Mas atenção: a ordem do resultado vem de onde?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Dois acumuladores: `seen := map[int]bool{}` (o que já vi) e `var result []int` (a resposta, na ordem).
- Percorra o **slice** de entrada (é ele que tem a ordem certa).
- Se o número ainda não foi visto: marque como visto e faça o `append`. ("Só faço X quando Y", do review 019.)
- Não monte o resultado percorrendo o map: a ordem do map é aleatória.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
seen := map[int]bool{}
var result []int
for _, n := range nums {
	if !seen[n] {
		seen[n] = true
		result = append(result, n)
	}
}
return result
```

</details>

## FirstRepeated

<details>
<summary>Nível 1 — direção</summary>

São dois padrões: **buscar** (a primeira que se repete: achou, acabou) e **já vi?**. Na tabela de `"abcba"`, em que volta aparece a primeira rune que já estava no `seen`?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- `seen := map[rune]bool{}`.
- Percorra as runes de `s` (é a string que tem a ordem).
- Em cada volta, **primeiro pergunte** se a rune já foi vista. Se sim, é a resposta: retorno antecipado com `true`.
- **Depois** marque a rune como vista.
- Se o loop acabar, nenhuma se repetiu: `0, false`.

A ordem dessas duas linhas importa: marcar antes de perguntar faz toda rune parecer repetida.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
seen := map[rune]bool{}
for _, r := range s {
	if seen[r] {
		return r, true
	}
	seen[r] = true
}
return 0, false
```

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `panic: assignment to entry in nil map` | O map foi criado com `var m map[K]V`, que é `nil`. Crie com `map[K]V{}` ou `make(map[K]V)`. |
| `stock = map[apple:2], want map[apple:5]` | Você usou `=` em vez de `+=`: o estoque foi substituído, não somado. |
| `WordCount("Go GO go") = map[GO:1 Go:1 go:1]` | Faltou passar o texto para minúsculas antes de contar. |
| `Unique([3 1 3 2 1]) = [1 2 3]` (ou outra ordem) | Você montou o resultado percorrendo o map, e a ordem do map é aleatória. Percorra o slice de entrada. |
| `FirstRepeated("abcba") = 'a', true, want 'b', true` | Ou você marcou a rune como vista **antes** de perguntar (toda rune vira "repetida" na primeira volta), ou contou tudo e depois percorreu o map (ordem aleatória). Pergunte primeiro, marque depois, percorrendo a string. |
| `invalid operation: ... (map can only be compared to nil)` | Maps não se comparam com `==`. Compare chave por chave, ou use `maps.Equal`. |
