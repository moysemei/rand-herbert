# Dicas — Drill 021

> **Regra:** antes de abrir uma dica, confira a tabela de operações do README: na maioria das vezes, a operação que falta está lá. Depois confira o seu plano (operação pronta? padrão? passos?) e abra **só o próximo nível**. Se mesmo assim travar, mande uma mensagem no chat.

## Refazer de memória (LongestWord)

Sem dicas aqui. Se travar 10 minutos: anote o que tentou e abra o `REVIEW.md` do drill 018. Depois feche e reescreva sem olhar.

## Price

<details>
<summary>Nível 1 — direção</summary>

Operação pronta. O café de graça (`"water": 0`) é o caso que decide qual dos três jeitos de perguntar "existe?" você precisa. Olhe a lista do README.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Como o preço 0 é um valor de verdade, perguntar `menu[item] != 0` diria que a água não existe. Use o "vírgula ok": ele devolve o valor e se a chave existe, que é exatamente o que a função promete devolver.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
p, ok := menu[item]
return p, ok
```

</details>

## RemoveItem

<details>
<summary>Nível 1 — direção</summary>

São duas operações: saber se existia e apagar. O `delete` não devolve nada. Em que ordem elas precisam acontecer?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- **Primeiro** pergunte se a chave existe, com o "vírgula ok" (o valor não interessa: use `_`).
- **Depois** apague com `delete`.
- Devolva a resposta da primeira pergunta.

Se apagar primeiro, a chave nunca vai existir na hora de perguntar.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
_, ok := stock[item]
delete(stock, item)
return ok
```

</details>

## CharFrequency

<details>
<summary>Nível 1 — direção</summary>

É o `LetterCount` do README, sem o filtro do espaço. Qual padrão?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

**Contar por chave:** um `map[rune]int{}` vazio; para cada rune de `s`, some 1 na chave dela; devolva o map.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
counts := map[rune]int{}
for _, r := range s {
	counts[r]++
}
return counts
```

</details>

## IsAnagram

<details>
<summary>Nível 1 — direção</summary>

O que "mesmas letras nas mesmas quantidades" tem a ver com a função que você acabou de escrever? Faça na mão: qual é o `CharFrequency("listen")`? E o `CharFrequency("silent")`?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Duas palavras são anagramas exatamente quando os maps de frequência das duas são iguais. Então: calcule o `CharFrequency` de cada uma e compare os dois maps. Para comparar maps existe uma operação pronta na tabela do README (o `==` não funciona com maps).

Só verificar se cada letra de `b` existe em `a` não basta: `"aab"` e `"abb"` têm as mesmas letras, mas não as mesmas quantidades.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`return maps.Equal(CharFrequency(a), CharFrequency(b))`, com `import "maps"`.

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `Price(menu, "water") = 0, false, want 0, true` | Você decidiu "existe" olhando se o valor é diferente de 0. Um preço 0 existe. Use o "vírgula ok". |
| `RemoveItem(stock, "apple") = false, want true` | Você perguntou se existia **depois** de apagar. Pergunte antes. |
| `RemoveItem(map[water:0], "water") = false` | Você usou o valor (0) para saber se existia. Use o "vírgula ok". |
| `delete(stock, item) (no value) used as value` | O `delete` não devolve nada: ele não pode estar do lado direito de um `=`. |
| `IsAnagram("aab", "abb") = true, want false` | Você só verificou se as letras existem, não as quantidades. Compare os maps de frequência inteiros. |
| `invalid operation: ... (map can only be compared to nil)` | Maps não se comparam com `==`. Use `maps.Equal` (pacote `maps`). |
| `panic: assignment to entry in nil map` | O map foi criado com `var`. Use `map[rune]int{}`. |
