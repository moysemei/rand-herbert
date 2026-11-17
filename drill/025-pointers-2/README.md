# Drill 025 — Ponteiros 2: devolver × escrever, e uma pilha

## Antes de tudo: refazer de memória (~15 min)

Do drill 023: **`func (r Rectangle) Fits(box Rectangle) bool`**: o retângulo `r` cabe dentro do `box`, em pé ou deitado? Encostar na borda vale.
- `{Width: 2, Height: 5}` em `{Width: 3, Height: 6}` → `true` (em pé)
- `{Width: 2, Height: 5}` em `{Width: 6, Height: 3}` → `true` (deitado)
- `{Width: 4, Height: 4}` em `{Width: 3, Height: 6}` → `false`

Crie o arquivo **`redo/redo.go`**, com **`package redo`** e este tipo:

```go
// Rectangle is a rectangle with integer sides.
type Rectangle struct {
	Width  int
	Height int
}
```

Não abra o `methods.go` nem o `REVIEW.md` do drill 023. Se travar no caso deitado, use a tabela do método: **primeiro com os números do exemplo, depois de onde vem cada número**. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Firmar o `&` e o `*` do drill 024: entregar o resultado com `return` ou escrevendo no endereço, e montar uma pilha com métodos de receiver ponteiro.

## O que é

**Os níveis** (a tabela do review 024, para ter à mão):

| Expressão | Tipo | O que é |
|---|---|---|
| `x` | `int` | o número |
| `&x` | `*int` | o endereço de `x` |
| `n`, dentro de `F(&x)` | `*int` | uma cópia do endereço de `x` |
| `*n` | `int` | o número que está no endereço (o próprio `x`) |

**O `&` sobe um nível** (do valor para o endereço). **O `*` desce um nível** (do endereço para o valor).

**Dois jeitos de uma função entregar o resultado:**

| | Devolve | Escreve no endereço |
|---|---|---|
| Recebe | uma cópia do valor | o endereço do valor |
| Entrega com | `return valor` | `*n = valor` |
| Quem chama | `y := F(x)`: o `x` não muda | `F(&x)`: o `x` muda |
| `return` | obrigatório, com valor | nenhum (ou um `return` vazio, para sair antes) |

**Uma pilha** é uma lista em que só se mexe no **topo**: o último que entrou é o primeiro que sai. Como uma pilha de pratos. Você vai montar a sua sobre um slice:

```go
// Stack is a stack of runes: the last one pushed is the first one popped.
type Stack struct {
	items []rune
}
```

**As operações:**

| Quero... | Na sua pilha | O que acontece | Em Python |
|---|---|---|---|
| empilhar | `s.Push('(')` | entra no topo (o fim do slice) | `stack.append("(")` |
| desempilhar | `r, ok := s.Pop()` | sai o do topo; `ok` é `false` se a pilha estiver vazia | `stack.pop()` (com a pilha vazia: `IndexError`) |
| começar vazia | `var s Stack` | `items` é `nil`: tamanho 0, e o `append` funciona nele | `stack = []` |

**O estado da pilha, chamada a chamada:**

| chamada | devolve | `items` depois |
|---|---|---|
| `var s Stack` | | `[]` |
| `s.Push('(')` | | `['(']` |
| `s.Push('[')` | | `['(' '[']` |
| `s.Pop()` | `'[', true` | `['(']` |
| `s.Pop()` | `'(', true` | `[]` |
| `s.Pop()` | `0, false` | `[]` |

## Por que é assim

- **Por que os métodos da pilha levam `*`:** o `Push` e o `Pop` mudam o `items`. Com receiver por valor, eles mexeriam numa cópia da `Stack`, e a pilha de quem chamou continuaria igual (a armadilha do `Click` no drill 023).
- **Por que o `Pop` devolve `(rune, bool)`:** é o mesmo "vírgula ok" do map. Pilha vazia não é erro de programa, é uma resposta: "não tinha nada". O Python levanta exceção; o Go devolve o `false` e deixa você decidir.

**Para que serve num backend de verdade:**

- **Pilha:** o "voltar" do navegador e o Ctrl+Z são pilhas. O *stack trace* que aparece quando um programa Go dá `panic` é a pilha de chamadas de funções. Validar se um JSON tem as chaves e os colchetes fechados é uma pilha, e isso é o problema extra de hoje, um clássico de entrevista.
- **Devolver × escrever:** em Go, o normal é **devolver**: é mais fácil de ler e de testar. Escrever no endereço fica para quando a função precisa mudar algo que já existe: um método que muda a struct, ou o `json.Unmarshal(data, &config)` da fase 2.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje): a mesma conta, entregue dos dois jeitos.

```go
package numbers

// Square returns n squared. The caller's variable does not change.
func Square(n int) int {
	return n * n
}

// SquareInPlace replaces the value at n with its square.
func SquareInPlace(n *int) {
	*n = *n * *n // the first * of each *n goes down to the int; the middle * multiplies
}
```

**Rastreando**, linha a linha:

| linha | `x` | `y` |
|---|---|---|
| `x := 4` | 4 | |
| `y := Square(x)` | 4 (não mudou) | 16 |
| `SquareInPlace(&x)` | **16** | 16 |

## Armadilhas

- **`&` no lugar de `*` dentro da função.** Com `n *int`, o número é `*n`. O `&n` é o endereço do endereço (`**int`).
- **`return` com valor numa função sem tipo de retorno:** `too many return values`. A função que escreve no endereço não devolve nada.
- **Go não tem índice negativo** (hábito do Python). O último item de um slice é `s[len(s)-1]`, não `s[-1]` (que dá `panic`).
- **O topo é o fim do slice.** O `append` põe no fim, então é do fim que o `Pop` tira. Tirar do começo seria uma fila (o primeiro que entrou sai primeiro), e não uma pilha.
- **Ler o último item de um slice vazio** dá `panic: index out of range [-1]`. Pergunte "está vazia?" **antes** de ler.

## O problema

Crie o arquivo **`pointers2.go`** nesta pasta, com **`package pointers2`** e o tipo `Stack` exatamente como está no "O que é".

1. **`func Abs(n int) int`**: devolve o valor absoluto de `n` (o número sem o sinal).
   - `Abs(-7)` → `7`; `Abs(3)` → `3`; `Abs(0)` → `0`
2. **`func MakeAbs(n *int)`**: troca o valor que está no endereço `n` pelo valor absoluto dele.
   - `x := -7`, `MakeAbs(&x)` → `x` vale `7`
   - `y := 3`, `MakeAbs(&y)` → `y` continua `3`
3. **`func (s *Stack) Push(r rune)`**: põe `r` no topo.
4. **`func (s *Stack) Pop() (rune, bool)`**: tira o topo e devolve ele com `true`. Com a pilha vazia, devolve `0, false`.
   - A sequência da tabela "O estado da pilha".

O `Abs` e o `Push` são as vitórias rápidas. O `MakeAbs` e o `Pop` são as partes de pensar. No `MakeAbs`: **já tenho uma função que faz parte disso?** No `Pop`: o que eu preciso perguntar antes de ler o topo, e como a pilha fica sem ele depois?

**Extra (se sobrar tempo):** **`func IsBalanced(s string) bool`**: os `()`, `[]` e `{}` de `s` estão bem fechados? Use a sua `Stack`. Outros caracteres são ignorados.
- `IsBalanced("([]{})")` → `true`
- `IsBalanced("a(b)c")` → `true`
- `IsBalanced("([)]")` → `false` (o `)` fecha antes do `]`)
- `IsBalanced("((")` → `false` (sobrou aberto)
- `IsBalanced("())")` → `false` (fechou sem ter aberto)
- `IsBalanced("")` → `true`

O `go test -v` cobra só os quatro primeiros. Para testar o extra: `go test -v -tags extra`.

## Checklist

- [ ] Comentário do pacote, do tipo e de cada função.
- [ ] Se travar: `// Plano:` ou tabela **antes** de abrir uma dica ou pedir ajuda.
- [ ] No fim do arquivo, as suas **"Dificuldades"**, com as suas palavras ("nenhuma", se for o caso).
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos quatro testes.
- [ ] Se fez o extra: `go test -v -tags extra` mostra `PASS` em todos.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 025: pointers 2"
  git push
  ```
