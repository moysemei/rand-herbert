# Drill 021 — Maps, operação por operação

## Antes de tudo: refazer de memória (~15 min)

Do drill 018: **`func LongestWord(s string) string`**, a palavra mais longa (separadas por espaço). No empate, a primeira. Sem palavras, `""`.
- `LongestWord("the quick brown fox")` → `"quick"`
- `LongestWord("go is fun")` → `"fun"`
- `LongestWord("")` → `""`

Crie o arquivo **`redo/redo.go`**, com **`package redo`**. Não abra o `patterns.go` nem o `REVIEW.md` do drill 018. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Voltar aos maps com calma: cada operação explicada uma por uma, o estado do map acompanhado volta a volta e, no fim, um problema clássico de entrevista resolvido trocando a representação.

## O que é

Um map é uma **agenda de contatos**: você procura pelo nome (a **chave**) e encontra o telefone (o **valor**). Ao contrário do slice, não existe "o primeiro" nem "o terceiro"; existe só "o de tal chave".

**As operações, uma por uma** (com `stock := map[string]int{"apple": 3}`):

| Quero... | Em Go | O que acontece | Em Python |
|---|---|---|---|
| criar vazio | `m := map[string]int{}` | map vazio, pronto para escrever | `m = {}` |
| criar com dados | `m := map[string]int{"a": 1}` | | `m = {"a": 1}` |
| ler | `v := stock["apple"]` | `v` = 3 | `stock["apple"]` |
| ler o que não existe | `v := stock["pear"]` | `v` = 0 (valor zero), **sem erro** | `KeyError` |
| saber se existe | `v, ok := stock["pear"]` | `v` = 0, `ok` = `false` | `"pear" in stock` |
| escrever / trocar | `stock["pear"] = 4` | cria a chave, ou troca o valor | `stock["pear"] = 4` |
| somar ao que tem | `stock["apple"] += 2` | 3 + 2 = 5; numa chave nova, 0 + 2 | `stock["apple"] = stock.get("apple", 0) + 2` |
| apagar | `delete(stock, "apple")` | some a chave; apagar uma que não existe não dá erro | `del stock["apple"]` |
| tamanho | `len(stock)` | quantas chaves | `len(stock)` |
| percorrer | `for k, v := range stock` | cada chave com seu valor, **em ordem aleatória** | `for k, v in stock.items()` |
| comparar dois | `maps.Equal(a, b)` (pacote `maps`) | `true` se têm as mesmas chaves e valores | `a == b` |

**Três jeitos de perguntar "existe?"**, e quando usar cada um:

1. **`v, ok := m[k]`:** quando 0 (ou `""`, ou `false`) pode ser um valor de verdade. Exemplo: um item que custa 0 é diferente de um item que não existe.
2. **`if seen[k]`** num `map[K]bool`: quando você só guarda `true`. A chave ausente vale `false`, então a pergunta funciona direto.
3. **Nem perguntar:** quando o valor zero já serve, como no `counts[k]++`.

**As assinaturas que você vai usar hoje** (o que entra e o que sai):

| Função | Assinatura |
|---|---|
| `strings.Fields` | `func Fields(s string) []string` |
| `strings.ToLower` | `func ToLower(s string) string` |
| `maps.Equal` | `func Equal(m1, m2 map[K]V) bool` |
| `delete` | `delete(m map[K]V, key K)`, sem retorno |

## Por que é assim

- **O valor zero no lugar do erro** é uma escolha do Go: ler uma chave ausente não quebra o programa, e o "vírgula ok" fica disponível para quando a diferença importa. Por isso o `counts[k]++` não precisa de `if`.
- **Trocar a representação resolve problemas.** Muita pergunta difícil sobre texto ou lista fica fácil quando você transforma os dados num map de contagem. "Essas duas palavras têm as mesmas letras?" vira "esses dois maps são iguais?". Esse é o truque do problema 4 de hoje, e um dos mais usados em entrevista.

## Como se usa

O método completo num problema que não é o de hoje, com o **estado do map** volta a volta:

**Problema:** `LetterCount(s string) map[rune]int`, quantas vezes aparece cada letra, ignorando espaços.

1. **Operação pronta?** Não existe uma que conte letras.
2. **Padrão:** "quantas vezes aparece cada uma?" → **contar por chave**.
3. **Passos:** "começo com um map vazio"; "para cada rune, se não for espaço, somo 1 na chave dela"; "devolvo o map".
4. **Go:**

```go
func LetterCount(s string) map[rune]int {
	counts := map[rune]int{}
	for _, r := range s {
		if r != ' ' {
			counts[r]++
		}
	}
	return counts
}
```

**Rastreando `LetterCount("ab a")`**, olhando o map depois de cada volta:

| volta | `r` | o que acontece | `counts` depois |
|---|---|---|---|
| — | | | `{}` |
| 1 | `'a'` | `counts['a']` era 0, vira 1 | `{a:1}` |
| 2 | `'b'` | `counts['b']` era 0, vira 1 | `{a:1 b:1}` |
| 3 | `' '` | espaço: nada | `{a:1 b:1}` |
| 4 | `'a'` | `counts['a']` era 1, vira 2 | `{a:2 b:1}` |

## Armadilhas

- **Valor no lugar da chave.** O que vai dentro do `[ ]` é a chave. `stock[3]` num `map[string]int` não compila: a chave é texto.
- **Operação de slice num map, ou o contrário.** `append` é de slice; `delete` é de map.
- **Map criado com `var`.** `var m map[string]int` é `nil`: escrever nele dá panic. Use `map[string]int{}`.
- **Guardar o resultado de uma ferramenta no tipo errado.** Confira a assinatura (tabela acima, `go doc`, ou `K` no nvim).
- **`delete` não devolve nada.** Para saber se a chave existia, pergunte **antes** de apagar.

## O problema

Crie o arquivo **`maps2.go`** nesta pasta, com **`package maps2`**, e implemente:

1. **`func Price(menu map[string]int, item string) (int, bool)`**: o preço do item, e se ele existe no cardápio.
   - Com `menu = {"coffee": 6, "water": 0}`:
   - `Price(menu, "coffee")` → `6, true`
   - `Price(menu, "water")` → `0, true` (existe, e é de graça)
   - `Price(menu, "cake")` → `0, false` (não existe)
2. **`func RemoveItem(stock map[string]int, item string) bool`**: apaga o item do próprio map. Devolve `true` se ele existia e `false` se não existia.
   - `{"apple": 3, "pear": 1}` com `"apple"` → devolve `true`, e o map fica `{"pear": 1}`
   - `{"pear": 1}` com `"kiwi"` → devolve `false`, e o map continua `{"pear": 1}`
3. **`func CharFrequency(s string) map[rune]int`**: quantas vezes aparece cada rune (espaços contam também).
   - `CharFrequency("banana")` → `{'a': 3, 'b': 1, 'n': 2}`
   - `CharFrequency("ação")` → `{'a': 1, 'ç': 1, 'ã': 1, 'o': 1}`
   - `CharFrequency("")` → `{}`
4. **`func IsAnagram(a, b string) bool`**: `a` e `b` têm exatamente as mesmas letras, nas mesmas quantidades (só que em outra ordem)?
   - `IsAnagram("listen", "silent")` → `true`
   - `IsAnagram("aab", "abb")` → `false` (as letras são as mesmas, as quantidades não)
   - `IsAnagram("rat", "car")` → `false`
   - `IsAnagram("", "")` → `true`

O `RemoveItem` e o `IsAnagram` são as partes de pensar. No `IsAnagram`, a pergunta-chave é: **que representação das duas palavras deixa essa pergunta fácil?**

**Antes de codar, no `// Plano:`, o método para cada função** (operação pronta? padrão? passos em português?). No `CharFrequency("banana")`, a tabela com o **map depois de cada volta**, como no exemplo acima.

E no fim do arquivo, as **"Dificuldades"** ("nenhuma", se for o caso).

## Checklist

- [ ] `// Plano:` com o método para as quatro funções e a tabela do `CharFrequency("banana")`.
- [ ] Comentário do pacote e de cada função.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos quatro testes.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 021: maps 2"
  git push
  ```
