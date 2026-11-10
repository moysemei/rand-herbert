# Review — Drill 021

## Veredito

Entregue no mesmo dia em que o drill foi reduzido, e com as quatro funções, extras incluídos. O refazer do `LongestWord` saiu certo de memória. Pelo registro honesto da mentoria, houve ajuda em pontos das quatro funções, então maps ainda não estão automáticos. Mas agora você monta e corrige código com map quando entende a operação, o que no 020 ainda não acontecia.

## Sobre o 020 e a pausa

Você ficou quatro dias parado porque saiu do 020 se achando burro. Vale olhar os fatos, porque eles não confirmam isso:

- **Parte da dificuldade do 020 foi do material,** como eu já registrei no review dele.
- **O repo mostra um padrão que se repete:** o que trava na primeira vez sai de memória alguns dias depois.
  - `Power`: dica 2 no drill 010, de memória no refazer do 014.
  - `ReverseNumber`: travou no 011, de memória no 015.
  - `SumNumbers`: o `Atoi` com `err` sumiu no 018 e voltou inteiro no 019.
  - `LongestWord`: buscar e melhor-até-agora confundidos no 018; hoje, certo de memória.

  Isso não é burrice. É aprendizado com dificuldade, que é o tipo que fixa. A estrutura nova sempre vai doer na primeira vez.

E fica um combinado mais útil do que parar: **quando um drill te deixar assim, me mande "drill menor" aqui no chat.** Eu corto na hora, como fiz com este. Um drill pequeno entregue vale mais do que quatro dias sem nenhum.

## Processo

- **No `RemoveItem`, você colocou as duas operações na ordem certa:** perguntar se existe e só depois apagar. É a armadilha principal dessa função.
- **No `IsAnagram`, você achou o `maps.Equal` como ferramenta,** o passo "existe uma operação pronta?" do método funcionando.
- **Os extras foram feitos mesmo sem serem obrigatórios.**
- **Desta vez não teve `// Plano:` nem as suas "Dificuldades",** só a avaliação do ChatGPT. Ela foi útil e honesta, e eu li tudo. Mas as suas palavras, mesmo curtas, me mostram como **você** pensou, e é isso que eu calibro. Os dois juntos são o ideal.
- **Faltaram também os comentários** do pacote e das funções.

## Erros e melhorias

**[ferramenta] `ok` não é `err`.** No `Price`, você tentou comparar o `ok` com `nil`, como faz com o erro do `Atoi`. Os dois usam dois valores de retorno, mas o segundo valor é de tipo diferente:

| | `v, ok := m[k]` | `n, err := strconv.Atoi(s)` |
|---|---|---|
| Tipo do segundo valor | `bool` | `error` |
| "Deu errado" se escreve | `!ok` | `err != nil` |
| O que ele diz | só sim ou não | sim ou não, **e por quê** |

A sua versão final está certa. Uma simplificação: quando a chave não existe, o `value` já vem 0, então o `if !ok { return 0, false }` só repete o que o próprio map já devolveu. Basta `return value, ok`.

**[lógica] Aproveitar o resultado que você já tem.** No `RemoveItem`, a estratégia estava certa, mas a primeira versão terminava com `return true`, ignorando o `ok` calculado duas linhas antes. É a lição do drill 004 em outra forma: **se você já calculou a resposta, devolva a variável que a guarda.** Antes de escrever um `return` com valor fixo, pergunte: "eu já tenho isso numa variável?".

**[ferramenta] Contar contra contar por chave.** No `CharFrequency`, a ideia de um `count` com `count++` conta **quantas runes há no total** (um número só). O pedido era **uma contagem por rune** (um número para cada chave):

| | `count++` | `counts[r]++` |
|---|---|---|
| Quantos números guarda | um | um por chave |
| `"banana"` termina como | `6` | `{a:3 b:1 n:2}` |
| Padrão do cardápio | contar | contar **por chave** |

O `counts[r]++` vem do `WordCount` do drill 020. Por isso o refazer do próximo drill é justamente o `WordCount`.

**[lógica] A representação certa, e reaproveitar o que você acabou de escrever.** No `IsAnagram`, a primeira ideia foi um `map[string]bool`. Um map de presença diz **quais** letras existem, mas perde **quantas**: `"aab"` e `"abb"` teriam o mesmo map. E a função que guarda as quantidades você tinha escrito dez linhas acima. Duas perguntas para acrescentar ao passo 2 do método:
1. **Que representação deixa essa pergunta fácil?** "Mesmas letras nas mesmas quantidades" pede contagens, não presença.
2. **Eu já escrevi uma função que faz parte disso?** Antes de criar algo, olhe o próprio arquivo.

## Solução de referência

```go
// Package maps2 practices map operations one by one.
package maps2

import "maps"

// Price returns the price of item and whether it is on the menu.
// The comma ok already gives both answers: return them as they are.
func Price(menu map[string]int, item string) (int, bool) {
	price, ok := menu[item]
	return price, ok
}

// RemoveItem deletes item from stock and reports whether it existed.
// Ask before deleting: after delete, the key never exists.
func RemoveItem(stock map[string]int, item string) bool {
	_, existed := stock[item]
	delete(stock, item)
	return existed
}

// CharFrequency returns how many times each rune appears in s.
// Pattern: count by key (one counter per rune, not one for all).
func CharFrequency(s string) map[rune]int {
	counts := map[rune]int{}
	for _, r := range s {
		counts[r]++
	}
	return counts
}

// IsAnagram reports whether a and b have the same runes with the same counts.
// Change the representation: two words are anagrams when their frequency maps are equal.
func IsAnagram(a, b string) bool {
	return maps.Equal(CharFrequency(a), CharFrequency(b))
}
```

Por que assim:

- **`existed` no lugar de `ok`** no `RemoveItem`: o nome diz o que a variável significa naquele contexto. `ok` é genérico; `existed` responde à pergunta da função.
- **O `IsAnagram` numa linha:** com a representação certa e a função já escrita, o problema que parecia difícil vira uma comparação. É exatamente esse tipo de "virada" que entrevistador procura.

## Refazer de memória (LongestWord)

Certo de memória, com o ponto de partida `""` e o `>` que mantém a primeira no empate. Os nomes ainda estão compridos (`cutMyPhrase`, `biggestWord`); `words` e `longest` dizem o mesmo.

## Para levar

`ok` é `bool` (`!ok`); `err` é `error` (`err != nil`). E antes de escrever algo novo, faça duas perguntas: que representação deixa isso fácil, e eu já tenho uma função que faz parte disso?
