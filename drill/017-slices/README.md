# Drill 017 — Slices

## Antes de tudo: refazer de memória (~15 min)

Do drill 013: **`func DigitalRoot(n int) int`**, que soma os dígitos de `n`, depois os dígitos do resultado, e assim por diante, até sobrar um dígito só (com `n` >= 0). Escreva também o ajudante **`func sumDigits(n int) int`**.
- `DigitalRoot(9875)` → `2` (9875 → 29 → 11 → 2)
- `DigitalRoot(99)` → `9`
- `DigitalRoot(7)` → `7`

Crie o arquivo **`redo/redo.go`**, com **`package redo`**. Não abra o `puzzles.go` nem o `REVIEW.md` do drill 013. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Guardar vários valores numa lista (um slice), percorrer, adicionar e trocar elementos de lugar.

## O que é

**Um slice é uma lista de valores do mesmo tipo,** que pode crescer:

```go
nums := []int{4, 8, 15}     // a slice with 3 ints
first := nums[0]            // 4: indexes start at 0
size := len(nums)           // 3
last := nums[len(nums)-1]   // 15: the last index is len - 1
nums[1] = 99                // nums is now [4 99 15]
```

**Percorrer** é com o mesmo `for range` das strings, só que agora ele te dá a posição e o **elemento**:

```go
for i, v := range nums {
	// i = 0, 1, 2 (the index)
	// v = 4, 99, 15 (the value at that index)
}
```

**Adicionar no fim** é com `append`, que **devolve** o slice novo. Você sempre guarda o resultado:

```go
var list []int          // an empty slice (nil): len 0, ready to use
list = append(list, 7)  // [7]
list = append(list, 3)  // [7 3]
```

**Trocar dois elementos de lugar** tem um atalho em Go, a atribuição dupla:

```go
nums[0], nums[2] = nums[2], nums[0] // swap the first and the last
```

(Também existem os *arrays*, como `[3]int`, que têm tamanho fixo. Em código Go do dia a dia, quase sempre você usa slices.)

## Por que é assim

- **O slice é a "lista" do Python,** mas com um tipo só: `[]int` guarda só `int`. Em Python, `nums.append(7)` é um método e muda a lista. Em Go, `append` é uma função e **devolve** o slice, porque às vezes precisa de mais espaço e cria um novo por baixo. Por isso é `list = append(list, 7)`.
- **Não existe índice negativo.** O `nums[-1]` do Python não existe; o último elemento é `nums[len(nums)-1]`.
- **Índice fora do limite não é erro de compilação: é erro na hora de rodar** (um *panic*). O programa para com `index out of range`.
- **Slices não se comparam com `==`.** O Go não sabe se você quer comparar os elementos ou outra coisa, então não deixa. Nos testes, eu uso `slices.Equal`.
- **Uma função que recebe um slice e muda os elementos dele muda o slice de quem chamou.** Isso é diferente do `int` do drill 011, que era uma cópia. O porquê disso é o assunto do próximo tópico; hoje você só vai usar.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package scores

// Average returns the average of the scores.
// ok is false when there are no scores.
func Average(scores []int) (float64, bool) {
	if len(scores) == 0 {
		return 0, false // no scores: no average
	}
	total := 0
	for _, s := range scores { // s is one score per round
		total += s
	}
	return float64(total) / float64(len(scores)), true
}

// Doubled returns a new slice with every score times 2.
// The original slice is not changed.
func Doubled(scores []int) []int {
	var result []int
	for _, s := range scores {
		result = append(result, s*2)
	}
	return result
}
```

## Armadilhas

- **`append` sem guardar o resultado.** `append(list, 7)` sozinho não compila: `append(...) (value of type []int) is not used`.
- **Começar o "maior" em 0.** Com `[-3, -1, -7]`, o 0 é maior que todos e vira a resposta errada. Comece com um elemento de verdade, o primeiro (a mesma ideia do `MinMaxDigit` do drill 014).
- **Slice vazio.** Antes de pegar `nums[0]`, confira se existe um elemento 0. Com um slice vazio, `nums[0]` dá panic.
- **Inverter trocando até o fim.** Se você troca as pontas e vai até o fim do slice, ele é invertido duas vezes e volta ao original. Pare no meio.
- **`for i := range nums` dá só o índice.** Para o valor, `for _, v := range nums` (o mesmo cuidado do drill 015).

## O problema

Crie o arquivo **`slices.go`** nesta pasta, com **`package lists`**, e implemente:

1. **`func Sum(nums []int) int`**: a soma de todos os elementos.
   - `Sum([]int{1, 2, 3})` → `6`
   - `Sum([]int{-5, 5})` → `0`
   - `Sum([]int{})` → `0`
2. **`func MaxOf(nums []int) (int, bool)`**: o maior elemento. Com um slice vazio, não existe maior: devolva `0, false`.
   - `MaxOf([]int{3, 9, 2})` → `9, true`
   - `MaxOf([]int{-3, -1, -7})` → `-1, true`
   - `MaxOf([]int{5})` → `5, true`
   - `MaxOf([]int{})` → `0, false`
3. **`func Countdown(n int) []int`**: um slice com `n`, `n-1`, ..., `1`.
   - `Countdown(3)` → `[3 2 1]`
   - `Countdown(1)` → `[1]`
   - `Countdown(0)` → `[]` (vazio)
4. **`func ReverseInPlace(nums []int)`**: inverte o próprio slice, trocando elementos de lugar. Não devolve nada e não cria slice novo.
   - `[1 2 3 4]` vira `[4 3 2 1]`
   - `[1 2 3]` vira `[3 2 1]`
   - `[7]` continua `[7]`

O `MaxOf` e o `ReverseInPlace` são as partes de pensar. O `ReverseInPlace` é a sua primeira ideia do drill 015 ("fazer com que a posição que o `i` me traga, sendo ela a última, vá para o começo"), agora com a ferramenta certa.

**Antes de codar**, no `// Plano:`:
- **Uma linha por função:** soma, procura, monta ou troca? Com que valor começa?
- **A tabela do `ReverseInPlace([1 2 3 4])`,** com as colunas volta, `i`, `j` (o índice da outra ponta) e o slice depois da troca. Em que volta o loop precisa parar?

E no fim do arquivo, as **"Dificuldades"** ("nenhuma", se for o caso).

Se travar, faça a tabela antes de abrir a dica. Depois as dicas, depois o chat.

## Checklist

- [ ] `// Plano:` com uma linha por função e a tabela do `ReverseInPlace`.
- [ ] Comentário do pacote e de cada função.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos quatro testes.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 017: slices"
  git push
  ```
