// Package books provides basic operations with books using structs.
package books

// Book represents a book with its title, author, pages, and reading status.
type Book struct {
	Title  string
	Author string
	Pages  int
	Read   bool
}

// NewBook creates a new book with its reading status set to false.
func NewBook(title, author string, pages int) Book {
	return Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

// TotalPages returns the sum of pages from all books.
func TotalPages(books []Book) int {
	sumPages := 0
	for _, p := range books {
		sumPages += p.Pages
	}

	return sumPages
}

// CountByAuthor returns the number of books written by each author.
func CountByAuthor(books []Book) map[string]int {
	count := map[string]int{}

	for _, a := range books {
		count[a.Author]++
	}

	return count
}

// MarkRead returns a copy of the book marked as read.
func MarkRead(b Book) Book {
	b.Read = true

	return b
}

/*
### Sugestão do aluno e do ChatGPT — Tornar o `// Plano:` opcional

Claude, o aluno gostaria de solicitar uma mudança na metodologia dos próximos drills: **que o `// Plano:` deixe de ser obrigatório e passe a ser opcional**.

Ele entende por que o planejamento foi introduzido e reconhece sua utilidade, especialmente nos exercícios em que tem dificuldade para identificar o algoritmo.

Entretanto, relatou que a obrigação de parar para escrever o plano antes de programar acaba interrompendo seu raciocínio. Quando já consegue visualizar uma solução, prefere colocar a mão no código diretamente.

No Drill 022, por exemplo, conseguiu implementar `TotalPages` e `CountByAuthor` sem precisar estruturar previamente os passos em português. Isso não prejudicou a qualidade das implementações.

**Minha avaliação como mentor:** considero a solicitação razoável. O plano deve funcionar como ferramenta de apoio, não como uma exigência burocrática.

Sugiro a seguinte abordagem:

- **Quando souber como começar:** pode ir diretamente para o código.
- **Quando travar na lógica:** utilizar o plano em português, exemplos concretos ou tabelas de rastreio.
- **Quando conhecer a lógica, mas não a sintaxe:** identificar a operação que falta, sem exigir um planejamento completo.
- **Quando pedir ajuda:** explicar o que tentou e onde travou, mesmo que não tenha escrito um plano formal.

Manteria o `// Dificuldades:` como parte importante da revisão, porque ele registra os obstáculos reais e permite acompanhar a evolução ao longo dos drills.

A proposta não é abandonar o raciocínio antes de programar, mas **permitir que o aluno escolha quando precisa externalizá-lo por escrito**.

O objetivo continua sendo desenvolver autonomia, e reconhecer quando já consegue resolver um problema sem planejamento formal também é um sinal de progresso.
*/

// Dificuldades:
// Agora sou eu escrevendo. A unica dificuldade mesmo foi na primeira func, onde que eu eu
// não tinha entendido como criar um novo Book baseado em uma estrutura definida como  a struct
// O resto das funções saiu bem natural e eu nem sei explicar como, sinceramente.
