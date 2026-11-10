# Drill 022 — Structs

## Antes de tudo: refazer de memória (~15 min)

Do drill 020: **`func WordCount(s string) map[string]int`**, quantas vezes cada palavra aparece, sem diferenciar maiúsculas e minúsculas.
- `WordCount("go is fun go")` → `{"go": 2, "is": 1, "fun": 1}`
- `WordCount("Go GO go")` → `{"go": 3}`
- `WordCount("")` → `{}`

Crie o arquivo **`redo/redo.go`**, com **`package redo`**. Não abra o `maps.go` nem o `REVIEW.md` do drill 020. Se travar por 10 minutos, anote no código o que tentou, e aí pode consultar. Para testar: `cd redo` e `go test -v`.

Num dia ruim, pule direto para o drill.

## Hoje

Juntar vários dados de uma mesma coisa num tipo seu (um livro com título, autor e páginas) e trabalhar com uma lista desses tipos.

## O que é

Uma **struct** é um tipo que você cria, com **campos** de nome e tipo fixos. É uma ficha: todo livro tem as mesmas linhas, cada uma com o seu tipo.

```go
type Book struct {
	Title  string
	Author string
	Pages  int
	Read   bool
}
```

**As operações, uma por uma** (com o `Book` acima):

| Quero... | Em Go | O que acontece | Em Python (dataclass) |
|---|---|---|---|
| criar com valores | `b := Book{Title: "Dune", Author: "Herbert", Pages: 412}` | os campos não citados ficam com o valor zero (`Read` = `false`) | `b = Book("Dune", "Herbert", 412)` |
| criar vazio | `var b Book` | todos os campos no valor zero: `""`, `""`, `0`, `false` | — |
| ler um campo | `b.Pages` | `412` | `b.pages` |
| mudar um campo | `b.Pages = 500` | muda só esse campo | `b.pages = 500` |
| comparar | `b1 == b2` | `true` se **todos** os campos forem iguais | `b1 == b2` |
| imprimir com nomes | `fmt.Sprintf("%+v", b)` | `{Title:Dune Author:Herbert Pages:412 Read:false}` | `print(b)` |
| lista de structs | `books := []Book{b1, b2}` | um slice comum, de `Book` | `books = [b1, b2]` |
| percorrer a lista | `for _, b := range books { b.Pages }` | `b` é cada livro, um por volta | `for b in books: b.pages` |

**Struct não é map**, e esta é a diferença que importa:

| | Struct `Book` | Map `map[string]int` |
|---|---|---|
| Os nomes (campos ou chaves) | fixos, escritos no código | decididos enquanto o programa roda |
| Os tipos dos valores | cada campo tem o seu (`string`, `int`, `bool`) | todos os valores têm o mesmo tipo |
| Acessar | `b.Pages` (ponto) | `m["pages"]` (colchete) |
| Errar o nome | erro de **compilação** | lê o valor zero, sem aviso |
| Use quando | a "forma" do dado é conhecida (um livro, um usuário) | as chaves vêm dos dados (contar palavras) |

**Struct é copiada quando passa para uma função.** Ao contrário do slice e do map, a função recebe uma **cópia**: mudar um campo lá dentro não muda o livro de quem chamou. Para devolver a mudança, a função devolve a struct nova. (Mudar o original direto é com ponteiro, assunto do tópico 18.)

## Por que é assim

- **Go não tem classes.** A struct guarda os dados, e o comportamento fica em funções (hoje) e em métodos (próximo tópico). É mais simples que herança: os dados são só dados.
- **Campos com maiúscula são exportados** (a mesma regra das funções, drill 001). `Title` é visível fora do pacote; um campo `title` não seria. Para o JSON da fase 2, isso vai importar.
- **O compilador confere os nomes.** Num map, `m["pgaes"]` (com erro de digitação) compila e lê 0. Numa struct, `b.Pgaes` não compila. É uma das maiores vantagens de usar struct para dados com forma conhecida.
- **A cópia na passagem é previsível:** uma função que recebe um `Book` não muda o seu livro escondido de você.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje), com o estado da lista volta a volta:

```go
package school

// Student is one student and their grade.
type Student struct {
	Name  string
	Grade int
}

// Passed returns the names of the students with grade 6 or more, in order.
// Pattern: build with a filter, reading a field of each struct.
func Passed(students []Student) []string {
	var names []string
	for _, s := range students { // s is one Student per round
		if s.Grade >= 6 {
			names = append(names, s.Name)
		}
	}
	return names
}
```

**Rastreando** com três alunos: Ana (8), Bia (5) e Caio (6):

| volta | `s` | `s.Grade >= 6`? | `names` depois |
|---|---|---|---|
| — | | | `[]` |
| 1 | `{Ana 8}` | sim | `[Ana]` |
| 2 | `{Bia 5}` | não | `[Ana]` |
| 3 | `{Caio 6}` | sim | `[Ana Caio]` |

## Armadilhas

- **Esquecer o nome do campo na criação.** `Book{"Dune", "Herbert", 412, false}` (sem os nomes) compila, mas exige **todos** os campos, na ordem certa, e quebra quando alguém adiciona um campo. Use sempre `Campo: valor`.
- **Mudar a cópia e esperar que o original mude.** Uma `func Promote(s Student)` que faz `s.Grade++` não muda o aluno de quem chamou: ela mexeu numa cópia, e a cópia some quando a função termina.
- **Mudar o `s` do `range`.** No `for _, s := range students`, o `s` também é uma cópia de cada aluno: `s.Grade++` ali não muda o slice.
- **Struct × map:** campo é com **ponto** (`s.Name`), chave de map é com **colchete** (`m["Ana"]`). E um campo é um valor como qualquer outro: `s.Name` é uma `string` e vale em qualquer lugar que aceite `string`.

## O problema

Crie o arquivo **`books.go`** nesta pasta, com **`package books`**. O tipo `Book` vai no seu arquivo, exatamente como está no "O que é" (os testes usam os mesmos nomes de campo).

1. **`func NewBook(title, author string, pages int) Book`**: um livro novo, ainda não lido.
   - `NewBook("Dune", "Herbert", 412)` → `{Title:Dune Author:Herbert Pages:412 Read:false}`
2. **`func TotalPages(books []Book) int`**: a soma das páginas de todos os livros.
   - Com Dune (412) e Emma (474) → `886`
   - Com a lista vazia → `0`
3. **`func CountByAuthor(books []Book) map[string]int`**: quantos livros de cada autor.
   - Com Dune (Herbert), Emma (Austen) e Dune Messiah (Herbert) → `{"Herbert": 2, "Austen": 1}`
   - Com a lista vazia → `{}`
4. **`func MarkRead(b Book) Book`**: devolve o livro marcado como lido. O livro original não pode mudar.
   - `MarkRead(dune)` → o mesmo livro, com `Read: true`; e o `dune` continua com `Read: false`

O `CountByAuthor` e o `MarkRead` são as partes de pensar.

**Antes de codar, no `// Plano:`, o método para cada função:** operação pronta (ou função sua que já faz parte)? Que padrão? Passos em português. No `CountByAuthor`, a tabela com o **map depois de cada volta**, para os três livros do exemplo.

E no fim do arquivo, as suas **"Dificuldades"**, mesmo que curtas ("nenhuma", se for o caso).

## Checklist

- [ ] `// Plano:` com o método para as quatro funções e a tabela do `CountByAuthor`.
- [ ] Comentário do pacote, do tipo `Book` e de cada função.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos quatro testes.
- [ ] Se fez o "refazer": `cd redo` e `go test -v` mostra `PASS`.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 022: structs"
  git push
  ```
