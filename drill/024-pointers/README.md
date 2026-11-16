# Drill 024 — Ponteiros

## Antes de tudo: refazer de memória (~15 min)

Do drill 021: **`func IsAnagram(a, b string) bool`**: `a` e `b` têm exatamente as mesmas letras, nas mesmas quantidades, em qualquer ordem?
- `IsAnagram("listen", "silent")` → `true`
- `IsAnagram("aab", "abb")` → `false` (mesmas letras, quantidades diferentes)
- `IsAnagram("", "")` → `true`

Crie o arquivo **`redo/redo.go`**, com **`package redo`**. Pode escrever uma função ajudante no mesmo arquivo. Não abra o `maps2.go` nem o `REVIEW.md` do drill 021. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Entender o `*` que você usou no `Deposit`: guardar o endereço de uma variável, mudar o valor por esse endereço, e devolver "não achei" com `nil`.

## O que é

Toda variável mora num lugar da memória. Um **ponteiro** é uma variável que guarda esse lugar (o **endereço**), e não o valor. É como o endereço de uma casa: com ele, você chega na casa de verdade e pode mudar o que tem lá dentro.

```go
score := 42     // score guarda 42
p := &score     // p guarda o endereço de score; o tipo de p é *int
*p = 50         // vai até o endereço e troca o valor: agora score é 50
```

**O `*` tem dois sentidos**, e essa é a confusão número um:

| Onde aparece | Exemplo | Lê-se |
|---|---|---|
| num **tipo** | `var p *int`, `func Reset(n *int)` | "ponteiro para int" |
| numa **expressão** | `*p = 7`, `x := *p` | "o valor que está no endereço `p`" |

**As operações, uma por uma:**

| Quero... | Em Go | O que acontece | Em Python |
|---|---|---|---|
| pegar o endereço de uma variável | `p := &x` | `p` aponta para `x` | não existe |
| declarar um ponteiro vazio | `var p *int` | `p` é `nil`: não aponta para nada | `p = None` |
| ler o valor apontado | `v := *p` | `v` recebe o valor de `x` | não existe |
| mudar o valor apontado | `*p = 7` | muda o próprio `x` | não existe para `int` |
| campo de struct por ponteiro | `pl.Score` (com `pl` do tipo `*Player`) | o Go segue o endereço sozinho, como no `Deposit` | `pl.score` |
| endereço de um item do slice | `&players[i]` | aponta para o item **dentro** do slice | `players[i]` (já é referência) |
| devolver "não achei" | `return nil` | ponteiro vazio | `return None` |
| conferir se achou | `if p == nil` | | `if p is None` |

**Ponteiro × o que você já conhece**, passando cada coisa para uma função:

| | `int`, `string`, struct | ponteiro (`*int`, `*Account`) | slice, map |
|---|---|---|---|
| A função recebe | uma cópia do valor | uma cópia do **endereço** (o mesmo lugar) | uma cópia que enxerga os mesmos itens |
| A função consegue mudar o de quem chamou? | não | sim, com `*p = ...` ou `p.Campo = ...` | sim, com `s[i] = ...` ou `m[k] = ...` |

**Assinaturas de hoje:** só as suas, no "O problema". Nenhum pacote novo.

## Por que é assim

- **Em Go, tudo é passado por cópia.** Um ponteiro também: a função recebe uma cópia do endereço. Mas uma cópia do endereço de uma casa leva à mesma casa. É assim que uma função muda o que é de quem chamou, e é só assim.
- **Em Python, você nunca escolhe.** Todo nome é uma referência, então uma função muda uma lista que recebe, mas não consegue mudar um `int` de quem chamou. Em Go, você escolhe: sem `*`, a função não mexe em nada seu; com `*`, a assinatura avisa que pode mexer.
- **`nil` é "não aponta para nada".** Ler ou mudar o valor de um ponteiro `nil` derruba o programa (`panic`). Por isso, quem recebe um ponteiro que pode ser `nil` confere antes de usar.

**Para que serve num backend de verdade:**

- **Métodos que mudam a struct:** o `*Account` do `Deposit` é um ponteiro. Hoje você vê o que ele é por dentro.
- **Funções que preenchem uma struct sua:** na fase 2, `json.Unmarshal(data, &config)` lê um JSON e preenche o `config`. Ela precisa do endereço (`&`), senão preencheria uma cópia.
- **Buscas que podem não achar:** `func FindUser(id int) *User` devolve o usuário, ou `nil` se ele não existe. É o mesmo papel do `ok` do map, num valor só. Em código real, quase sempre junto com um `error`: `(*User, error)`.
- **Structs grandes:** passar o endereço evita copiar a struct inteira a cada chamada.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package game

// Clamp lowers *n to limit when it is bigger. It changes the caller's variable.
func Clamp(n *int, limit int) {
	if *n > limit { // *n reads the value at the address
		*n = limit // *n = ... writes at the address
	}
}
```

**Rastreando** o estado, linha a linha:

| linha | `score` | `p` | `*p` |
|---|---|---|---|
| `score := 42` | 42 | — | — |
| `p := &score` | 42 | endereço de `score` | 42 |
| `*p = 120` | **120** | o mesmo | 120 |
| `Clamp(&score, 100)` | **100** | o mesmo | 100 |
| `Clamp(&score, 500)` | 100 | o mesmo | 100 (120 não é maior que 500: nada muda) |

Repare: nenhuma linha depois da segunda escreve `score =`, e mesmo assim ele muda. Quem muda é quem tem o endereço.

## Armadilhas

- **Esquecer o `&` na chamada.** `Clamp(score, 100)` não compila: a função pede um `*int` e você passou um `int`. O erro diz `cannot use score (variable of type int) as *int value`.
- **Esquecer o `*` na expressão.** `n = 0` dentro da função tenta trocar o **endereço** guardado em `n`, não o valor. Para mudar o valor, é `*n = 0`.
- **Trocar os ponteiros em vez dos valores.** Dentro da função, `a, b = b, a` troca só as cópias locais dos endereços; quem chamou não vê nada. Para trocar o que está **nos** endereços, o `*` aparece dos dois lados.
- **A variável do `range` é uma cópia** (drill 022). No `for _, pl := range players`, mudar `pl.Score` não muda o slice, e `&pl` aponta para a cópia, não para o item. Para chegar no item de verdade, use o índice: `players[i]` ou `&players[i]`.
- **Usar um ponteiro `nil`.** `p.Score` com `p` valendo `nil` dá `panic: runtime error: invalid memory address or nil pointer dereference`.

## O problema

Crie o arquivo **`pointers.go`** nesta pasta, com **`package pointers`** e este tipo, exatamente assim:

```go
// Account is a bank account; Balance is in reais.
type Account struct {
	Owner   string
	Balance int
}
```

1. **`func Double(n *int)`**: dobra o valor que está no endereço `n`.
   - `x := 5`, `Double(&x)` → `x` vale `10`
2. **`func Swap(a, b *int)`**: troca os valores que estão nos dois endereços.
   - `x, y := 1, 2`, `Swap(&x, &y)` → `x` vale `2` e `y` vale `1`
3. **`func AddBonus(accounts []Account, bonus int)`**: soma `bonus` ao saldo de **todas** as contas. O slice de quem chamou tem que mudar.
   - `[{Ana 100} {Bia 50}]` com bônus `10` → o slice fica `[{Ana 110} {Bia 60}]`
4. **`func FindByOwner(accounts []Account, owner string) *Account`**: o endereço da conta daquele dono **dentro do slice** (mudar o saldo por esse ponteiro muda o slice), ou `nil` se não existir.
   - `FindByOwner(accs, "Bia")` → aponta para `accs[1]`; depois de `p.Balance += 5`, o `accs[1].Balance` também mudou
   - `FindByOwner(accs, "Caio")` → `nil`

O `Double` e o `Swap` são as vitórias rápidas. O `AddBonus` e o `FindByOwner` são as partes de pensar: os dois dependem de chegar no item **de verdade** do slice, e não numa cópia dele. No `FindByOwner`, qual padrão do cardápio é esse, e o que a função devolve quando o loop termina sem achar?

## Checklist

- [ ] Comentário do pacote, do tipo e de cada função.
- [ ] Se travar: `// Plano:` ou tabela **antes** de abrir uma dica ou pedir ajuda.
- [ ] No fim do arquivo, as suas **"Dificuldades"**, com as suas palavras ("nenhuma", se for o caso).
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos quatro testes.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 024: pointers"
  git push
  ```
