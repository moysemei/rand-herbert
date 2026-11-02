# Review — Drill 018

## Veredito

Entregue, com o refazer e os quatro testes passando, e com o método aplicado de verdade: o plano tem exemplo, padrão e passos para as quatro funções. O `CountChar` e o `Contains` saíram do plano para o código sem estrutura pronta. O `LongestWord` e o `ParseAll` pediram ajuda, e o review explica os dois pontos em que você travou.

## Processo

- **O plano é exatamente o método do review 017,** nas quatro funções: exemplo com valores de verdade, o nome do padrão e os passos em português. No `ParseAll`, você identificou que eram **dois** padrões juntos. Esse é o passo que, há uma semana, vinha do ChatGPT.
- **"Consegui implementar a função sem receber o código"** no `Contains`, incluindo a dúvida certa sobre onde vai o `return false`, resolvida pelo motivo certo: "preciso verificar todos os números".
- **No `LongestWord`, você achou um bug interessante:** comparar `word > biggestWord` compila, mas compara as strings em **ordem alfabética**, não pelo tamanho. Você percebeu e trocou para `len(word) > len(biggestWord)`.
- **O `>` (e não `>=`)** mantém a primeira palavra no empate, como o enunciado pedia.
- **O "Geral" das "Dificuldades" é uma autoavaliação honesta e certa:** você melhorou em reconhecer padrões, mas ainda precisa de prática quando o problema combina operações. Concordo com cada linha.
- **O refazer do `Reverse` saiu de memória, com comentário.** A função que precisou de ajuda no drill 015 agora sai sozinha.

## Erros e melhorias

**[lógica] Buscar ou melhor até agora?** No `LongestWord`, você pensou primeiro em buscar. A pergunta que separa os dois padrões é: **a resposta depende de comparar com todos os outros?**
- **Buscar:** basta achar **um** que serve, e dá para parar cedo. "Existe algum 8?" Achou, acabou.
- **Melhor até agora:** só dá para saber qual é o maior depois de ver **todos**, e não dá para parar cedo. "Qual é a palavra mais longa?" Uma palavra de 5 letras pode perder para a última.

Pergunta prática: "posso dar a resposta antes de terminar o loop?" Se sim, é buscar. Se não, é melhor até agora.

**[lógica] Esquecer ferramentas que você já usou.** O `strings.Fields` e o `strconv.Atoi` com `err` são do drill 016, e os dois sumiram. Isso é normal: é exatamente para isso que existe o refazer de memória. Por isso o refazer do drill de hoje é o `SumNumbers`, que usa os dois. Uma sugestão opcional: um arquivo seu, `notes/toolbox.md`, com uma linha por ferramenta nova (`strings.Fields(s) -> splits on spaces`). Consultar a própria cola é permitido. O que você escreve com as suas palavras você lembra melhor.

**[idiomático] Nomes longos demais.** `phraseToWords`, `newSliceWithOnlyIntegers` e `biggestWord` descrevem bem, mas em Go a convenção é: **escopo pequeno, nome curto**. Dentro de uma função de 10 linhas, `words`, `nums` e `longest` dizem a mesma coisa sem pesar a leitura. Nomes longos ficam para coisas que vivem longe de onde foram criadas, como funções exportadas e campos de struct. Repare que você foi do `f` (drill 011) até o `newSliceWithOnlyIntegers`. O ponto certo fica no meio.

**[idiomático] Comentário do pacote.** "shows how to run through slices, counting, if contains and which one is best" fica difícil de ler. Uma frase com o que o pacote faz: `// Package patterns solves small problems by matching them to known patterns.`

## Solução de referência

```go
// Package patterns solves small problems by matching them to known patterns.
package patterns

import (
	"strconv"
	"strings"
)

// CountChar returns how many times target appears in s.
// Pattern: count.
func CountChar(s string, target rune) int {
	count := 0
	for _, r := range s {
		if r == target {
			count++
		}
	}
	return count
}

// Contains reports whether target is in nums.
// Pattern: search. "Found" can leave from inside the loop; "not found" only after it.
func Contains(nums []int, target int) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}
	return false
}

// LongestWord returns the longest word in s; on a tie, the first one.
// It returns "" when s has no words.
// Pattern: best so far. "" is a safe start: no word is shorter than it.
func LongestWord(s string) string {
	longest := ""
	for _, word := range strings.Fields(s) {
		if len(word) > len(longest) { // > keeps the first one on a tie
			longest = word
		}
	}
	return longest
}

// ParseAll converts every word to an int.
// ok is false, and the result is nil, if any word is not a number.
// Patterns: build + search (the first error decides the answer).
func ParseAll(words []string) ([]int, bool) {
	var nums []int
	for _, word := range words {
		n, err := strconv.Atoi(word)
		if err != nil {
			return nil, false
		}
		nums = append(nums, n)
	}
	return nums, true
}
```

Por que assim:

- **O `strings.Fields(s)` direto no `range`,** sem variável intermediária: ele só é usado ali.
- **O ponto de partida `""` no `LongestWord`** é seguro pelo motivo que está no comentário. Compare com o `MaxOf` do drill 017, em que começar com 0 podia vencer errado (`[-3 -1 -7]`). A pergunta é sempre a mesma: **esse ponto de partida pode vencer errado?**

## Refazer de memória (Reverse)

Certo de memória, com comentário. A função que precisou de ajuda no drill 015 está consolidada.

## Para levar

Buscar para quando acha; melhor até agora precisa ver todos. E o método está funcionando: nas duas funções em que o padrão era claro, você saiu do plano para o código sozinho.
