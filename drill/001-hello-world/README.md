# Drill 001 — Hello, World

**Hoje:** escrever seu primeiro arquivo Go do zero e fazer os testes passarem.

## O que é

Todo arquivo Go tem a mesma estrutura, de cima para baixo:

1. **`package`** — a primeira linha. Diz a qual pacote o arquivo pertence. O pacote é a unidade de organização do Go: todos os arquivos `.go` de uma mesma pasta pertencem ao mesmo pacote e enxergam tudo uns dos outros.
2. **`import`** — os outros pacotes que este arquivo usa (da biblioteca padrão ou de terceiros).
3. **Declarações** — funções, tipos, variáveis e constantes.

Existem dois tipos de pacote:

- **`package main`** vira um programa executável. Precisa ter uma `func main()`, que é por onde o programa começa. É o que você roda com `go run`.
- **Qualquer outro nome** (`package hello`) é uma biblioteca: código feito para ser usado por outro código. Não tem `main` e não roda sozinho. Você verifica se ela funciona com `go test`.

Hoje você escreve uma biblioteca, e quem usa ela são os testes.

A forma de uma função:

```go
func Name(param Type) ReturnType {
	return value
}
```

## Por que é assim

- **O tipo vem depois do nome** (`name string`, e não `string name` como em Java). Lê-se da esquerda para a direita: "name, que é uma string". E diferente de Python, todo parâmetro e todo retorno têm tipo declarado. Se os tipos não batem, o compilador recusa o programa: o erro aparece antes de rodar, não em produção.
- **Letra maiúscula = público.** Go não tem `public` nem `private`. Se o nome começa com maiúscula (`Hello`), ele é *exportado* e outros pacotes podem usar. Se começa com minúscula (`hello`), só é visível dentro do próprio pacote. Essa regra vale para tudo: funções, tipos, campos de struct.
- **Import não usado é erro de compilação**, não aviso. Variável declarada e não usada também. Go é rígido com código morto para manter projetos limpos e a compilação rápida.
- **O teste mora do lado do código**, num arquivo que termina em `_test.go`, no mesmo pacote. O `go test` acha e roda sozinho. Sem framework e sem configuração: o pacote `testing` já vem com a linguagem.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package area

import "math"

// Square returns the area of a square.
func Square(side float64) float64 {
	return side * side
}

// Circle returns the area of a circle.
func Circle(radius float64) float64 {
	return math.Pi * radius * radius
}

// Label returns a text like "Area of the circle".
func Label(shape string) string {
	return "Area of the " + shape
}
```

Repare que:

- A ordem é `package`, depois `import`, depois as funções.
- Para usar algo de outro pacote, você escreve `pacote.Nome` (`math.Pi`). Só funciona porque `Pi` começa com maiúscula.
- O `+` junta (concatena) strings.
- `float64` é número com casas decimais. Tipos são o assunto do drill 002, não se preocupe com eles agora.
- O comentário acima de uma função exportada começa com o nome dela. É a convenção de documentação do Go.

## Armadilhas

- **Aspas.** String em Go é sempre com aspas duplas: `"texto"`. Aspas simples (`'a'`) são um único caractere (uma *rune*), não uma string. Vindo de Python, esse é o erro número 1.
- **O nome do pacote não é o nome da pasta.** A pasta se chama `001-hello-world`, mas isso não é um nome de pacote válido (começa com número e tem hífen). Por isso o pacote é `hello`. O arquivo de teste já declara `package hello`, e o seu arquivo tem que declarar o mesmo.
- **A chave `{` fica na mesma linha do `func`.** Colocar na linha de baixo é erro de sintaxe em Go.
- **Vermelho no começo é o ponto de partida**, não fracasso. Se você rodar `go test` antes de escrever qualquer coisa, vai ver `undefined: Hello`. É o compilador dizendo exatamente o que está faltando.

## O problema

Crie o arquivo **`hello.go`** nesta pasta, com **`package hello`**, e implemente três funções:

1. **`Hello() string`**
   - Retorna `"Hello, World!"`.
2. **`Greet(name string) string`**
   - `Greet("Gopher")` → `"Hello, Gopher!"`
   - `Greet("Ada")` → `"Hello, Ada!"`
3. **`Shout(name string) string`**
   - Igual a `Greet`, mas tudo em maiúsculas.
   - `Shout("Gopher")` → `"HELLO, GOPHER!"`
   - **Desafio:** não escreva o texto de novo dentro de `Shout`. Reaproveite `Greet`. A biblioteca padrão do Go tem um pacote inteiro só para trabalhar com texto, e achar a função certa faz parte do exercício. A lista de pacotes está em https://pkg.go.dev/std.

Se travar, as dicas estão em `HINTS.md`.

## Checklist

- [ ] Logo abaixo do `package hello`, um comentário `// Plano:` com 2 ou 3 linhas dizendo como você vai resolver (pode ser em português).
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  git add .
  git commit -m "drill 001: hello world"
  git push
  ```
