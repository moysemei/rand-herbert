# Review — Drill 020

## Veredito

Entregue, os quatro testes e o refazer passam, mas maps ainda não firmaram, e parte disso é responsabilidade minha. Você tem razão: o README apresentou uma estrutura nova em poucas linhas, sem mostrar cada operação com calma nem comparar com o slice. O drill 021 volta em maps, com a explicação que faltou.

## Processo

- **Você registrou as tentativas erradas e o motivo de cada uma.** É por elas que dá para ver exatamente onde estava a confusão (está tudo abaixo).
- **O seu plano do `AddStock` era uma estratégia válida.** "Se o item já existe, soma; se não, cria com a quantidade." Em Python, seria exatamente assim (`if item in stock: ... else: ...`). Faltava saber uma coisa específica do Go, explicada abaixo, que transforma os dois ramos numa linha só.
- **Você nomeou a sua maior confusão:** "confundo slice com map, tanto a funcionalidade, mas ainda mais a sintaxe". Ela tem um motivo concreto, também explicado abaixo.
- **O `FirstRepeated` saiu com as duas decisões certas:** perguntar ao `seen` antes de marcar, e retornar na primeira repetição.
- **A avaliação do ChatGPT que você colou foi útil.** Concordo com o diagnóstico dela, principalmente com a ideia de separar quatro coisas diferentes: algoritmo errado, estratégia válida porém mais longa, operação que você ainda não conhecia, e erro de sintaxe. Este review segue essa separação.

## O que faltou no README

**Slice e map usam o mesmo colchete, mas o que vai dentro é diferente.** Essa é a origem da confusão:

| | Slice `[]int` | Map `map[string]int` |
|---|---|---|
| O que vai no `[ ]` | uma **posição**: 0, 1, 2... | uma **chave**: `"apple"`, `"pear"`... |
| Analogia | uma fila numerada | uma agenda de contatos: nome → telefone |
| Ordem | tem, e é a das posições | não tem (o `range` sai embaralhado) |
| Adicionar | `s = append(s, v)` | `m[k] = v` |
| Remover | fatiar e montar outro | `delete(m, k)` |
| Ler algo que não existe | panic (`index out of range`) | valor zero, sem erro |
| `for a, b := range` | `a` = posição, `b` = valor | `a` = chave, `b` = valor |

Na sua tentativa do `Unique`, o `nums[num]` usava o **valor** como se fosse **posição**, e o `delete(nums, num)` usava uma operação de map num slice. Era a tabela acima, misturada.

**Por que o `AddStock` é uma linha:** ler uma chave que não existe num `map[string]int` devolve 0. Então, para a pera nova, `stock["pear"] += 4` faz `0 + 4`, e para a maçã que já existe faz `3 + 2`. Os dois ramos do seu plano viram a mesma operação. Na sua tentativa, `if stock[item] == item` comparava o **valor** (um `int`, a quantidade) com a **chave** (uma `string`, o nome), e por isso não compilava.

## Erros e melhorias

**[ferramenta] Saber o tipo que cada função devolve.** No `WordCount`, você fez `words := strings.ToLower(s)` (`words` virou `string`) e depois `words = strings.Fields(s)`, que devolve `[]string` e não cabe numa `string`. A estratégia estava certa; faltava saber o **tipo de saída** de cada ferramenta:
- `strings.ToLower(s string) string`: texto entra, texto sai.
- `strings.Fields(s string) []string`: texto entra, lista de palavras sai.

Para descobrir sem decorar: `go doc strings.Fields` no terminal, ou `K` em cima do nome da função no nvim, que mostra a assinatura quando o LSP está ligado. Na sua cola `notes/toolbox.md`, vale anotar a assinatura de cada ferramenta.

**[lógica] Remover do original contra montar um novo.** No `Unique`, a primeira ideia foi apagar os repetidos do próprio slice. Mesmo com a sintaxe certa, apagar de uma lista enquanto você percorre essa mesma lista é uma fonte clássica de bug (os índices mudam no meio do caminho). A regra do drill 019 continua valendo: **filtrar é montar um novo**, e o map só serve de memória ("já vi?").

**[idiomático] Comentários.** O formato está certo, falta o inglês:
- `AddStock add a quantity to an key value` → `AddStock adds qty to the stock of item.`
- `WordCount count how many words exists in s` diz outra coisa: a função não conta as palavras, conta **quantas vezes cada uma aparece** → `WordCount returns how many times each word appears in s.`
- `an slice` → **a** slice.

## Sobre a ajuda

Pela avaliação que você colou, o trecho central do `Unique` (`!seen[num]` e `seen[num] = true`) veio pronto. Sem culpa nenhuma: o README não tinha te dado as peças para construí-lo sozinho. Só repito uma coisa: quando travar, me mande mensagem aqui. Eu respondo com pergunta ou dica, e o seu raciocínio fica seu.

## Solução de referência

```go
// Package counting counts and remembers values with maps.
package counting

import "strings"

// AddStock adds qty to the stock of item, creating the item if it is new.
// A missing key reads as 0, so "exists" and "new" are the same line.
func AddStock(stock map[string]int, item string, qty int) {
	stock[item] += qty
}

// WordCount returns how many times each word appears in s, ignoring case.
// Pattern: count by key.
func WordCount(s string) map[string]int {
	counts := map[string]int{} // not var: a nil map panics on write
	for _, word := range strings.Fields(strings.ToLower(s)) {
		counts[word]++
	}
	return counts
}

// Unique returns nums with each value once, in order of first appearance.
// Patterns: build + seen. The order comes from the slice, never from the map.
func Unique(nums []int) []int {
	seen := map[int]bool{}
	var result []int
	for _, n := range nums {
		if !seen[n] {
			seen[n] = true
			result = append(result, n)
		}
	}
	return result
}

// FirstRepeated returns the first rune of s that appears for the second time.
// ok is false when no rune repeats.
// Patterns: search + seen. Ask first, then mark.
func FirstRepeated(s string) (rune, bool) {
	seen := map[rune]bool{}
	for _, r := range s {
		if seen[r] {
			return r, true
		}
		seen[r] = true
	}
	return 0, false
}
```

Por que assim:

- **Cada comentário diz a decisão que não é óbvia:** a chave ausente valendo 0, o `var` que daria panic, a ordem vindo do slice, e perguntar antes de marcar.
- **O seu `Unique` e o de referência são iguais na lógica.** A única diferença é a ordem das duas linhas dentro do `if`, e as duas ordens funcionam.

## Refazer de memória (ReverseInPlace)

Ficou certo, mas com consulta depois dos 10 minutos (como a regra permite): o `len(nums)/2` e o `len(nums)-1-i` não voltaram sozinhos. Ele vai aparecer de novo num refazer mais para frente. É assim que a repetição espaçada funciona: o que falha volta.

## Para levar

Slice: o `[ ]` recebe uma posição, e a ordem existe. Map: o `[ ]` recebe uma chave, não existe ordem, e a chave ausente vale zero. Quando o seu plano for válido mas parecer complicado, pergunte se a linguagem tem um atalho para aquilo; no map, o atalho é o valor zero.
