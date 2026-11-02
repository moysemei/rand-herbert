# Drill 019 — Slices 2: fatiar, copiar e filtrar

## Antes de tudo: refazer de memória (~15 min)

Do drill 016: **`func SumNumbers(s string) (int, bool)`**, a soma dos números separados por espaço. Se algum pedaço não for número, `0, false`.
- `SumNumbers("1 2 3")` → `6, true`
- `SumNumbers(" 10   20 ")` → `30, true`
- `SumNumbers("4 x 5")` → `0, false`

Crie o arquivo **`redo/redo.go`**, com **`package redo`**. Não abra o `text.go` nem o `REVIEW.md` do drill 016. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Pegar um pedaço de um slice, entender quando dois slices dividem os mesmos elementos, e montar slices novos sem estragar o original.

## O que é

**Fatiar.** `s[a:b]` é o pedaço que começa no índice `a` e vai **até `b`, sem incluir o `b`**:

```go
nums := []int{10, 20, 30, 40, 50}
nums[1:3] // [20 30]: indexes 1 and 2
nums[:2]  // [10 20]: from the start
nums[3:]  // [40 50]: to the end
```

**Um slice é uma janela, não uma caixa.** Por baixo de todo slice existe um bloco de memória (um *array*), e o slice só guarda onde a janela começa e quantos elementos ela mostra. Fatiar cria **outra janela para o mesmo bloco**:

```go
nums := []int{10, 20, 30, 40, 50}
part := nums[1:3] // [20 30], looking at the same memory as nums
part[0] = 99      // changes nums too: nums is now [10 99 30 40 50]
```

É por isso que o `ReverseInPlace` do drill 017 mudava o slice de quem chamou: a função recebeu uma janela para o mesmo bloco.

**Para ter uma cópia independente,** copie os elementos para um slice novo:

```go
backup := slices.Clone(part) // package "slices": a new block, same values
backup[0] = 1                // part and nums don't change
```

(Também dá para fazer à mão, com `make` e `copy`, mas o `slices.Clone` diz a intenção numa palavra.)

**Juntar um slice no fim de outro:** `append` com `...` depois do segundo slice.

```go
a := []int{1, 2}
b := []int{3, 4}
a = append(a, b...) // [1 2 3 4]: "..." spreads b into separate values
```

## Por que é assim

- **No Python, `lista[1:3]` cria uma lista nova.** No JS, `arr.slice(1, 3)` também. No Go, `nums[1:3]` **não** copia nada: é uma janela para os mesmos elementos. É uma escolha de desempenho: passar um slice de um milhão de elementos para uma função custa o mesmo que passar um de três, porque nada é copiado.
- **O preço é ter cuidado.** Se uma função devolve um pedaço do slice que recebeu, quem chamou e quem recebeu passam a mexer nos mesmos elementos, e esse é um bug clássico. Quando o resultado precisa ser independente, copie.
- **Fatiar além do tamanho dá panic,** como o índice: `nums[:10]` num slice de 5 dá `slice bounds out of range`.
- **Uma função que filtra devolve um slice novo** e não mexe no original. Quem chama decide o que fazer com os dois.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package playlist

import "slices"

// Skip returns the songs after the first n, as an independent copy.
// If n is larger than the playlist, it returns an empty slice.
func Skip(songs []string, n int) []string {
	if n > len(songs) {
		n = len(songs) // never slice past the end
	}
	return slices.Clone(songs[n:]) // the caller can change it safely
}

// Without returns a new playlist without the given song.
// Pattern: build with a filter. The original is not changed.
func Without(songs []string, banned string) []string {
	var result []string
	for _, s := range songs {
		if s != banned {
			result = append(result, s)
		}
	}
	return result
}
```

## Armadilhas

- **O `b` de `s[a:b]` não entra.** `nums[:2]` são dois elementos: os índices 0 e 1.
- **Devolver um pedaço do original.** `return nums[:n]` funciona até alguém mudar o resultado, e aí o original muda junto. Se o resultado precisa ser independente, copie.
- **Fatiar além do fim:** `panic: slice bounds out of range`. Limite o índice antes.
- **`append(a, b)` com dois slices não compila:** `cannot use b (variable of type []int) as int value`. É `append(a, b...)`.
- **Procurar no lugar errado.** No `RemoveDuplicates`, a pergunta "eu já vi esse número?" é sobre o resultado que você está montando, não sobre a entrada (que sempre contém o número).

## O problema

Crie o arquivo **`slices2.go`** nesta pasta, com **`package slices2`**. Nenhuma função pode mudar o slice que recebe.

1. **`func Evens(nums []int) []int`**: só os números pares, na mesma ordem.
   - `Evens([]int{1, 2, 3, 4})` → `[2 4]`
   - `Evens([]int{1, 3})` → `[]`
   - `Evens([]int{})` → `[]`
2. **`func FirstN(nums []int, n int) []int`**: os primeiros `n` elementos (com `n` >= 0), numa **cópia independente**: mudar o resultado não pode mudar o original. Se `n` for maior que o slice, devolva todos.
   - `FirstN([]int{1, 2, 3, 4}, 2)` → `[1 2]`
   - `FirstN([]int{1, 2, 3, 4}, 10)` → `[1 2 3 4]`
   - `FirstN([]int{1, 2, 3, 4}, 0)` → `[]`
3. **`func RemoveDuplicates(nums []int) []int`**: cada número uma vez só, na ordem em que apareceu primeiro.
   - `RemoveDuplicates([]int{1, 2, 1, 3, 2})` → `[1 2 3]`
   - `RemoveDuplicates([]int{5, 5, 5})` → `[5]`
   - `RemoveDuplicates([]int{4, 3, 2})` → `[4 3 2]`
   - `RemoveDuplicates([]int{})` → `[]`
   - Escreva um ajudante **`func contains(nums []int, target int) bool`**: é o `Contains` do drill 018, de memória.

O `FirstN` e o `RemoveDuplicates` são as partes de pensar.

**Antes de codar, no `// Plano:`, o método para cada função:** exemplo, padrão (ou padrões) e passos em português. O cardápio está no README do drill 018. No `RemoveDuplicates`, inclua a tabela do `[1 2 1 3 2]`, com as colunas `v`, "já está no resultado?" e `result`.

E no fim do arquivo, as **"Dificuldades"** ("nenhuma", se for o caso): em qual passo do método você travou?

## Checklist

- [ ] `// Plano:` com exemplo, padrão e passos para as três funções, e a tabela do `RemoveDuplicates`.
- [ ] Comentário do pacote e de cada função (o ajudante também).
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 019: slices 2"
  git push
  ```
