# Review — Drill 022

## Veredito

Entregue na quarta: as quatro funções certas, comentários em tudo e o refazer certo de memória. A única trava foi montar o `Book` com os valores dos parâmetros, e ela é a mesma passagem do drill 005 com outra roupa.

## Processo

- **O `TotalPages` e o `CountByAuthor` saíram sem ajuda,** e você escreveu que "nem sabe explicar como". Eu sei explicar: o `CountByAuthor` é o `WordCount` (020) e o `CharFrequency` (021) com outra chave. Você reconheceu o padrão **contar por chave** sem precisar dizer o nome dele. É o cardápio de padrões virando automático, que é o objetivo dele.
- **No `MarkRead`, você mudou a própria cópia recebida e devolveu,** em vez de montar um `Book` novo campo por campo. Isso mostra que você entendeu a parte nova de hoje: o parâmetro já é uma cópia sua, e mexer nele é seguro.
- **Comentários no pacote, no tipo e em todas as funções,** com o verbo certo (*returns*, *creates*, *represents*). No 021 não tinha nenhum.
- **As "Dificuldades" são suas** e dizem exatamente onde você travou. É com esse registro que eu calibro o próximo drill.

## Sobre o `// Plano:` opcional

Pedido aceito. A partir de agora:

- **Se você já vê a solução,** vai direto para o código.
- **Se travar,** escreva o plano (exemplo à mão, padrão, passos) ou a tabela **antes** de abrir uma dica ou de pedir ajuda. É nesse momento que ele vale o tempo que custa.
- **As "Dificuldades" continuam obrigatórias,** como você mesmo propôs.

Um dado honesto: no 021 também não teve plano, e houve ajuda nas quatro funções. No 022, as partes de pensar usavam um padrão que você já tinha feito duas vezes. Então a regra não é "o plano não serve"; é "o plano serve quando a lógica ainda não está pronta na cabeça".

Fica um aviso para o futuro: no problema de entrevista (na dose de 1h30), o `// Approach:` antes de codar continua. Explicar a ideia antes de escrever o código é justamente o que o entrevistador avalia.

## Erros e melhorias

**[ferramenta] Montar a struct com os valores dos parâmetros.** A sua estratégia estava certa: criar um `Book` e devolver. O que faltou foi a escrita, e aqui a culpa é em parte do README: a tabela mostrou só o literal com valores fixos (`Book{Title: "Dune", ...}`), nunca dentro de uma função. O salto é o mesmo do drill 005: o exemplo tem valores fixos, e no código cada valor fixo vira o nome da variável de onde ele veio. O `"Dune"` do exemplo chega pelo parâmetro `title`, então fica `Title: title`. Da próxima vez, pergunte: **de onde vem cada valor deste literal?**

**[idiomático] A variável do loop leva o nome do item, não do campo que você vai usar.** Você escreveu `for _, p := range books` (p de *pages*?) e `for _, a := range books` (a de *author*?). Funciona, mas o nome engana: `p` e `a` são livros inteiros. Você provavelmente batizou a variável pelo que ia tirar de cada item. Da próxima vez, dê o nome do item (`b` ou `book`), e o campo diz o que você tira dele: `b.Pages`, `b.Author`. É o ponto de atenção "a variável do loop é o item da vez", agora no nome.

**[idiomático] `counts`, no plural, para um map de contagens.** Um map guarda um número por chave (a lição do 021), e o plural deixa isso visível na leitura. O mesmo vale para o `count` do refazer.

## Solução de referência

```go
// Package books keeps track of books with a struct.
package books

// Book is one book on the shelf.
type Book struct {
	Title  string
	Author string
	Pages  int
	Read   bool
}

// NewBook returns an unread book.
// Each value of the literal comes from a parameter; Read is left out, so it is false.
func NewBook(title, author string, pages int) Book {
	return Book{Title: title, Author: author, Pages: pages}
}

// TotalPages returns the sum of the pages of all books.
// Pattern: accumulate, adding one field of each struct.
func TotalPages(books []Book) int {
	total := 0
	for _, b := range books {
		total += b.Pages
	}
	return total
}

// CountByAuthor returns how many books each author has.
// Pattern: count by key, with a field of the struct as the key.
func CountByAuthor(books []Book) map[string]int {
	counts := map[string]int{}
	for _, b := range books {
		counts[b.Author]++
	}
	return counts
}

// MarkRead returns b marked as read; the caller's book does not change.
// b is already a copy, so changing it here is safe.
func MarkRead(b Book) Book {
	b.Read = true
	return b
}
```

Por que assim:

- **`b` nos dois loops:** o mesmo nome para a mesma coisa (um livro), e o campo diz o que se usa dele.
- **`Read` fora do literal do `NewBook`:** o valor zero de `bool` já é `false`. Escrever `Read: false` não está errado, mas não acrescenta nada.

## Refazer de memória (WordCount)

Certo de memória. No 020, o `WordCount` travou na diferença entre `string` e `[]string`; hoje, o `ToLower` e o `Fields` saíram encadeados na ordem certa. O nome `stringToMap` engana: a variável é um `[]string` de palavras, não um map. `words` diz o que ela é.

## Para levar

Uma struct é uma ficha de campos fixos: você lê com ponto e cria com `Campo: valor`, e numa função os valores vêm dos parâmetros. Passada para uma função, ela é copiada, então mexer no parâmetro não mexe no original.
