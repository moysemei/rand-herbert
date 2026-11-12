# Drill 023 — Métodos

## Antes de tudo: refazer de memória (~15 min)

Do drill 020: **`func Unique(nums []int) []int`**, cada número uma vez só, na ordem em que apareceu primeiro.
- `Unique([]int{3, 1, 3, 2, 1})` → `[3 1 2]`
- `Unique([]int{5, 5})` → `[5]`
- `Unique([]int{})` → `[]`

Crie o arquivo **`redo/redo.go`**, com **`package redo`**. Não abra o `maps.go` nem o `REVIEW.md` do drill 020. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Pendurar funções num tipo seu (`r.Area()`, `acc.Deposit(50)`) e escolher entre um método que só lê a struct e um que muda a struct original.

## O que é

Um **método** é uma função com um **receiver**: um parâmetro especial, escrito **antes** do nome, que diz a qual tipo ela pertence.

```go
func (r Rectangle) Area() int { ... }
//    ^^^^^^^^^^^ receiver: o método pertence ao Rectangle; dentro dele, o retângulo se chama r
```

**Função × método**, a mesma conta escrita dos dois jeitos:

| | Função (o que você já conhece) | Método |
|---|---|---|
| Declaração | `func Area(r Rectangle) int` | `func (r Rectangle) Area() int` |
| Chamada | `Area(box)` | `box.Area()` |
| Onde mora | solta no pacote | pendurada no tipo |
| Dois tipos com o mesmo nome | não pode (um `Area` por pacote) | pode: `Rectangle.Area` e `Circle.Area` |

**Os dois tipos de receiver.** É aqui que entra o `MarkRead` do drill 022:

| | Receiver por valor `(r Rectangle)` | Receiver por ponteiro `(a *Account)` |
|---|---|---|
| O método recebe | uma **cópia** da struct | o **endereço** da struct original |
| Mudar um campo lá dentro | muda só a cópia, que some no fim | muda o original, e a mudança fica |
| Use quando | o método só lê | o método muda algum campo |
| Ler e mudar campos dentro | `r.Width` | `a.Balance` (igual: o Go segue o endereço sozinho) |
| Chamar | `box.Area()` | `acc.Deposit(50)` (igual: o Go passa o endereço sozinho) |

O `*` antes do tipo quer dizer "ponteiro para". O que é um ponteiro de verdade fica para o tópico 18. Hoje basta a regra: **método que muda a struct leva `*`**.

**As operações, uma por uma:**

| Quero... | Em Go | O que acontece | Em Python |
|---|---|---|---|
| declarar um método que lê | `func (r Rectangle) Perimeter() int { return 2 * (r.Width + r.Height) }` | `r` é uma cópia | `def perimeter(self): return 2 * (self.width + self.height)` |
| declarar um método que muda | `func (c *Counter) Click() { c.Count++ }` | `c` aponta para o original | `def click(self): self.count += 1` |
| chamar | `box.Area()` | o Go passa `box` como receiver | `box.area()` |
| usar o resultado | `total := box.Area() + 10` | um método que devolve valor é uma expressão como qualquer outra | igual |
| chamar outro método dentro de um método | `return r.Area() > 100` | o receiver também tem os métodos do tipo | `self.area()` |

**As assinaturas de hoje:** os quatro métodos estão no "O problema", e a única ferramenta da stdlib é a de sempre: `fmt.Sprintf(format string, a ...any) string` (se precisar).

## Por que é assim

- **Go não tem classes,** mas tem métodos. Qualquer tipo que você declara no pacote pode ganhar métodos, e os dados continuam separados do comportamento: a struct diz o que é, os métodos dizem o que ela faz.
- **O receiver é explícito.** Em Python, o `self` é sempre uma referência, então qualquer método pode mudar o objeto. Em Go, quem escreve o método **escolhe**: sem `*`, o método promete que não muda nada; com `*`, avisa que muda. Quem lê a assinatura já sabe.
- **Um tipo, um tipo de receiver.** O guia oficial do Go recomenda não misturar: se algum método do tipo precisa de `*`, use `*` em todos os métodos dele. Por isso, hoje, o `Rectangle` (que nunca muda) tem só receiver por valor, e o `Account` (que muda) tem só `*`.
- **O nome do receiver é curto,** uma ou duas letras do tipo (`r`, `a`, `c`). Nunca `self` nem `this`.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package clicks

// Counter counts clicks on a button.
type Counter struct {
	Name  string
	Count int
}

// NewCounter returns a counter with no clicks.
// Each value of the literal comes from a parameter (the gap from drill 022).
func NewCounter(name string) Counter {
	return Counter{Name: name}
}

// Click adds one click. Pointer receiver: it changes the original counter.
func (c *Counter) Click() {
	c.Count++
}

// Reset sets the count back to zero. Pointer receiver, like every method of Counter.
func (c *Counter) Reset() {
	c.Count = 0
}
```

E a armadilha principal, para comparar: o mesmo `Click` **sem** o `*`.

```go
// WRONG: value receiver. c is a copy, so the click is lost.
func (c Counter) Click() {
	c.Count++
}
```

**Rastreando** o estado de `likes`, linha a linha:

| linha | `likes.Count` depois | por quê |
|---|---|---|
| `likes := NewCounter("likes")` | 0 | o literal não citou `Count`: valor zero |
| `likes.Click()` | 1 | receiver `*`: mudou o original |
| `likes.Click()` | 2 | |
| `likes.Reset()` | 0 | |
| `likes.Click()` com a versão **errada** | 0 | mudou uma cópia, que sumiu no fim do método |

Repare que a versão errada **compila e roda**. O único sinal é o valor que não muda, e é isso que o teste de hoje confere.

## Armadilhas

- **Esquecer o `*` no método que muda.** Nenhum erro, nenhum aviso: a mudança só some. Se o teste disser que o saldo continuou o mesmo depois de um depósito, olhe o receiver primeiro.
- **Chamar o método como função.** `Area(box)` não existe; é `box.Area()`. E o contrário: um método não aparece solto no pacote.
- **Escrever o receiver no lugar errado.** `func Area(r Rectangle) int` é uma função comum. O receiver vai **entre o `func` e o nome**: `func (r Rectangle) Area() int`.
- **Hábito do Python: `return` num método que não devolve nada.** Se a assinatura não tem tipo de retorno (como o `Click`), o método não devolve nada: ele só muda o receiver.

## O problema

Crie o arquivo **`methods.go`** nesta pasta, com **`package methods`** e estes dois tipos, exatamente assim (os testes usam os mesmos nomes):

```go
// Rectangle is a rectangle with integer sides.
type Rectangle struct {
	Width  int
	Height int
}

// Account is a bank account; Balance is in reais.
type Account struct {
	Owner   string
	Balance int
}
```

**Rectangle** (só lê: receiver por valor):

1. **`func (r Rectangle) Area() int`**: a área.
   - `Rectangle{Width: 3, Height: 4}` → `12`
2. **`func (r Rectangle) Fits(box Rectangle) bool`**: o retângulo `r` cabe dentro do `box`? Ele pode entrar **em pé ou deitado** (girado 90°). Encostar na borda vale.
   - `{Width: 2, Height: 5}` em `{Width: 3, Height: 6}` → `true` (em pé)
   - `{Width: 2, Height: 5}` em `{Width: 6, Height: 3}` → `true` (deitado: o 5 vai na largura, o 2 na altura)
   - `{Width: 4, Height: 4}` em `{Width: 3, Height: 6}` → `false`
   - `{Width: 3, Height: 6}` em `{Width: 3, Height: 6}` → `true` (do mesmo tamanho)

**Account** (muda: receiver por ponteiro):

3. **`func (a *Account) Deposit(amount int)`**: soma `amount` ao saldo. Valor zero ou negativo não muda nada.
   - Saldo 100, `Deposit(50)` → saldo 150
   - Saldo 150, `Deposit(-20)` → saldo continua 150
4. **`func (a *Account) Withdraw(amount int) bool`**: tira `amount` do saldo, **só se** for positivo e não passar do saldo. Devolve se o saque aconteceu. Se não aconteceu, o saldo não muda.
   - Saldo 100, `Withdraw(30)` → `true`, saldo 70
   - Saldo 70, `Withdraw(500)` → `false`, saldo continua 70
   - Saldo 70, `Withdraw(70)` → `true`, saldo 0

O `Area` e o `Deposit` são as vitórias rápidas. O `Fits` e o `Withdraw` são as partes de pensar. No `Fits`, a pergunta é: **quais são os candidatos** (as formas de o retângulo entrar), e qual é a condição de cada um? No `Withdraw`: **que casos impedem o saque**, e o que a função faz em cada um?

## Checklist

- [ ] Comentário do pacote, dos tipos e de cada método.
- [ ] Se travar: `// Plano:` ou tabela **antes** de abrir uma dica ou pedir ajuda.
- [ ] No fim do arquivo, as suas **"Dificuldades"** ("nenhuma", se for o caso).
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos quatro testes.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 023: methods"
  git push
  ```
