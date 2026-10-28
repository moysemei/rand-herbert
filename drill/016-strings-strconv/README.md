# Drill 016 — Os pacotes `strings` e `strconv`

## Antes de tudo: refazer de memória (~15 min)

Do drill 012: **`func FizzBuzz(n int) string`**.
- `FizzBuzz(15)` → `"FizzBuzz"`
- `FizzBuzz(9)` → `"Fizz"`
- `FizzBuzz(10)` → `"Buzz"`
- `FizzBuzz(7)` → `"7"`

Crie o arquivo **`redo/redo.go`**, com **`package redo`**. Não abra o `rules.go` nem o `REVIEW.md` do drill 012. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Usar a biblioteca padrão para trabalhar com texto, e conferir se uma conversão de texto para número deu certo.

## O que é

**O pacote `strings`** tem as funções de texto. Todas recebem uma string e **devolvem uma nova**, porque strings são imutáveis (drill 015).

| Função | Exemplo | Resultado |
|---|---|---|
| `strings.ToLower(s)` | `strings.ToLower("GoLang")` | `"golang"` |
| `strings.TrimSpace(s)` | `strings.TrimSpace("  hi  ")` | `"hi"` |
| `strings.ReplaceAll(s, old, new)` | `strings.ReplaceAll("a-b-c", "-", "")` | `"abc"` |
| `strings.Contains(s, sub)` | `strings.Contains("golang", "lang")` | `true` |
| `strings.ContainsRune(s, r)` | `strings.ContainsRune("aeiou", 'e')` | `true` |
| `strings.Fields(s)` | `strings.Fields(" a  b ")` | as palavras `"a"` e `"b"` |

O `strings.Fields` devolve uma lista de palavras (um *slice*, assunto do tópico 12). Por hoje, basta saber que dá para percorrer com `for _, word := range strings.Fields(s)`.

**O pacote `strconv`** converte entre texto e número. O `strconv.Itoa` (número → texto) você já conhece do drill 002. O caminho contrário é o `strconv.Atoi`, que pode falhar (`"abc"` não é número), e por isso devolve dois valores: o número e um **`error`**.

```go
n, err := strconv.Atoi("42")
if err != nil {
	// the conversion failed: n is 0 and must not be used
}
```

O `error` é o irmão mais velho do `ok` do drill 014. O `ok` só diz "deu certo ou não"; o `err` diz **por que** não deu. `nil` quer dizer "nenhum erro". Por hoje, basta isto: **`err != nil` quer dizer que deu errado.** O resto vem na fase 2.

Para importar dois pacotes, junte os dois entre parênteses:

```go
import (
	"strconv"
	"strings"
)
```

## Por que é assim

- **Em Python, texto tem métodos:** `s.lower()`, `s.strip()`. Em Go, as funções ficam num pacote: `strings.ToLower(s)`, `strings.TrimSpace(s)`. Para ler a documentação no terminal: `go doc strings.TrimSpace`. Na web: pkg.go.dev/strings.
- **Em Python, `int("abc")` lança uma exceção.** Em Go, o `Atoi` devolve o erro como valor, e você confere na linha seguinte. É o padrão `value, err :=` que você vai escrever centenas de vezes no backend.
- **Não ignore o erro.** `n, _ := strconv.Atoi(s)` compila, mas o `Atoi("abc")` devolve 0, e você não tem como diferenciar de um `"0"` de verdade.
- **No backend, tudo chega como texto:** parâmetro de URL, campo de formulário, linha de arquivo. Limpar e validar texto é trabalho de todo dia.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package form

import (
	"strconv"
	"strings"
)

// CleanEmail returns email without spaces at the ends and in lowercase.
func CleanEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// ParseQuantity converts text like " 3 " into a quantity.
// ok is false when the text is not a number or the quantity is not positive.
func ParseQuantity(text string) (int, bool) {
	n, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return 0, false // not a number
	}
	if n <= 0 {
		return 0, false // a number, but not a valid quantity
	}
	return n, true
}
```

## Armadilhas

- **Chamar a função e jogar o resultado fora.** `strings.ToLower(s)` sozinho numa linha compila, mas não muda o `s`. É `s = strings.ToLower(s)`. Nem o compilador nem o `go vet` avisam.
- **Esquecer um import, ou importar sem usar.** `undefined: strings` ou `"strconv" imported and not used`.
- **O `Atoi` não aceita espaço:** `strconv.Atoi(" 5")` dá erro. Limpe antes com `TrimSpace`, ou separe as palavras com `Fields`.
- **Usar o número depois de um erro.** Quando `err != nil`, o número vale 0 e não significa nada. Trate o erro antes de usar o valor.
- **Comparar strings** é com `==` mesmo. Não precisa de loop.

## O problema

Crie o arquivo **`text.go`** nesta pasta, com **`package text`**, e implemente:

1. **`func Normalize(s string) string`**: a string em minúsculas e sem nenhum espaço.
   - `Normalize("A Man A Plan")` → `"amanaplan"`
   - `Normalize("  Go  ")` → `"go"`
   - `Normalize("")` → `""`
2. **`func IsPalindrome(s string) bool`**: o texto é palíndromo, ignorando maiúsculas e espaços?
   - `IsPalindrome("Ana")` → `true`
   - `IsPalindrome("A man a plan a canal Panama")` → `true`
   - `IsPalindrome("socorram me subi no onibus em marrocos")` → `true`
   - `IsPalindrome("Go")` → `false`
   - `IsPalindrome("")` → `true`
   - Use o seu `Normalize`, e escreva um ajudante **`func reverse(s string) string`** (o `Reverse` do drill 015, de memória).
3. **`func SumNumbers(s string) (int, bool)`**: a soma dos números separados por espaço. Se algum pedaço não for número, devolva `0, false`.
   - `SumNumbers("1 2 3")` → `6, true`
   - `SumNumbers(" 10   20 ")` → `30, true`
   - `SumNumbers("-2 5")` → `3, true`
   - `SumNumbers("")` → `0, true` (nenhum número: a soma é 0)
   - `SumNumbers("4 x 5")` → `0, false`

O `IsPalindrome` e o `SumNumbers` são as partes de pensar.

**Antes de codar**, no `// Plano:`:
- **Uma linha por função:** qual função do `strings` ou do `strconv` resolve qual pedaço?
- **À mão:** o `Normalize("A man a plan a canal Panama")` (qual texto sai?) e o `SumNumbers("4 x 5")` (em qual volta para, e o que devolve?).

E no fim do arquivo, as **"Dificuldades"** ("nenhuma", se for o caso).

Se travar, faça a tabela antes de abrir a dica. Depois as dicas, depois o chat.

## Checklist

- [ ] `// Plano:` com uma linha por função e as duas contas à mão.
- [ ] Comentário do pacote e de cada função (o ajudante também).
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 016: strings and strconv"
  git push
  ```
