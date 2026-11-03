# Drill 020 — Maps

## Antes de tudo: refazer de memória (~15 min)

Do drill 017: **`func ReverseInPlace(nums []int)`**, que inverte o próprio slice trocando elementos de lugar, sem devolver nada e sem criar slice novo.
- `[1 2 3 4]` vira `[4 3 2 1]`
- `[1 2 3]` vira `[3 2 1]`
- `[7]` continua `[7]`

Crie o arquivo **`redo/redo.go`**, com **`package redo`**. Não abra o `slices.go` nem o `REVIEW.md` do drill 017. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Guardar valores por **chave** (um map), contar ocorrências e lembrar "o que eu já vi" sem percorrer uma lista.

## O que é

**Um map liga chaves a valores:**

```go
prices := map[string]int{"coffee": 6, "bread": 2} // key type string, value type int
prices["juice"] = 8          // add or replace
p := prices["coffee"]        // 6
size := len(prices)          // 3
delete(prices, "bread")      // remove a key
```

**Ler uma chave que não existe devolve o valor zero** do tipo (0, `""`, `false`), sem erro. Para saber se ela existe de verdade, use o "vírgula ok" do drill 014:

```go
p, ok := prices["tea"] // p = 0, ok = false: "tea" is not in the map
```

**Percorrer** é com `for range`, que dá a chave e o valor. **A ordem é aleatória:** muda de uma execução para outra.

```go
for item, price := range prices {
	// item = key, price = value, in no particular order
}
```

**Criar um map vazio** para ir preenchendo: `counts := map[string]int{}` (ou `make(map[string]int)`).

**Dois padrões novos para o cardápio:**

| Padrão | Pergunta | Esqueleto |
|---|---|---|
| **Contar por chave** | "Quantas vezes aparece cada um?" | `counts := map[K]int{}` · `for _, x := range xs { counts[x]++ }` · `return counts` |
| **Já vi?** | "Esse eu já encontrei antes?" | `seen := map[K]bool{}` · `for _, x := range xs { if seen[x] { ... }; seen[x] = true }` |

O `counts[x]++` funciona mesmo na primeira vez, porque a chave ausente vale 0. E o `seen[x]` de uma chave ausente vale `false`.

## Por que é assim

- **É o `dict` do Python e o `Map`/objeto do JS,** com tipos fixos: `map[string]int` só aceita chave `string` e valor `int`.
- **Chave ausente não é erro.** No Python, `d["x"]` lança `KeyError`. Em Go, devolve o valor zero, e o "vírgula ok" serve para quando você precisa diferenciar "não existe" de "existe e vale 0". Por isso o `counts[x]++` não precisa de `if`.
- **A ordem aleatória é de propósito.** O Go embaralha a ordem para ninguém depender dela sem querer. Se a ordem importa, ela vem de um slice, não do map.
- **Buscar num map é rápido,** não importa o tamanho. Perguntar "o 7 está no slice?" percorre o slice inteiro; perguntar `seen[7]` vai direto. É por isso que maps aparecem em quase toda solução boa de entrevista.
- **Map, como slice, é passado por referência:** uma função que muda o map que recebeu muda o map de quem chamou.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package votes

// Tally returns how many votes each candidate got.
// Pattern: count by key.
func Tally(ballots []string) map[string]int {
	counts := map[string]int{}
	for _, name := range ballots {
		counts[name]++ // a missing key starts at 0
	}
	return counts
}

// HasVoted reports whether voter is in the registry. Uses the comma ok.
func HasVoted(registry map[string]bool, voter string) bool {
	_, ok := registry[voter]
	return ok
}
```

## Armadilhas

- **Map `nil` não aceita escrita.** `var counts map[string]int` cria um map `nil`; ler funciona, mas `counts["a"]++` dá `panic: assignment to entry in nil map`. Crie com `map[string]int{}` ou `make`.
- **Depender da ordem do `range`.** Se o resultado precisa de ordem, percorra o **slice** de entrada e use o map só para consultar.
- **Confundir "valor zero" com "não existe".** `prices["tea"]` vale 0 tanto se o chá não existe quanto se ele custa 0. Quando a diferença importa, use o "vírgula ok".
- **Maps não se comparam com `==`.** Nos testes, eu uso `maps.Equal`.

## O problema

Crie o arquivo **`maps.go`** nesta pasta, com **`package counting`**, e implemente:

1. **`func AddStock(stock map[string]int, item string, qty int)`**: soma `qty` ao estoque do item, no próprio map, criando o item se ele não existir. Não devolve nada.
   - `{"apple": 3}` + `("apple", 2)` → `{"apple": 5}`
   - `{"apple": 3}` + `("pear", 4)` → `{"apple": 3, "pear": 4}`
2. **`func WordCount(s string) map[string]int`**: quantas vezes cada palavra aparece, sem diferenciar maiúsculas e minúsculas. As palavras são separadas por espaços.
   - `WordCount("go is fun go")` → `{"go": 2, "is": 1, "fun": 1}`
   - `WordCount("Go GO go")` → `{"go": 3}`
   - `WordCount("")` → `{}` (map vazio)
3. **`func Unique(nums []int) []int`**: cada número uma vez só, **na ordem em que apareceu primeiro**. É o `RemoveDuplicates` do drill 019, agora com um map no lugar do `contains`.
   - `Unique([]int{3, 1, 3, 2, 1})` → `[3 1 2]`
   - `Unique([]int{5, 5})` → `[5]`
   - `Unique([]int{})` → `[]`
4. **`func FirstRepeated(s string) (rune, bool)`**: lendo da esquerda para a direita, a primeira rune que aparece pela **segunda vez**. Se nenhuma se repete, `0, false`.
   - `FirstRepeated("abcba")` → `'b', true` (o segundo `b` aparece antes do segundo `a`)
   - `FirstRepeated("hello")` → `'l', true`
   - `FirstRepeated("abc")` → `0, false`

O `Unique` e o `FirstRepeated` são as partes de pensar.

**Antes de codar, no `// Plano:`, o método para cada função:** existe operação pronta? Se não, qual padrão (os dois novos estão na tabela acima)? Depois, os passos em português. No `FirstRepeated`, inclua a tabela do `"abcba"`, com as colunas `r`, "já vi?" e `seen` depois da volta.

E no fim do arquivo, as **"Dificuldades"** ("nenhuma", se for o caso).

## Checklist

- [ ] `// Plano:` com o método para as quatro funções e a tabela do `FirstRepeated("abcba")`.
- [ ] Comentário do pacote e de cada função.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos quatro testes.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 020: maps"
  git push
  ```
