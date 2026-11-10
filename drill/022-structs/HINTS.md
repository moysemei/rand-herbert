# Dicas — Drill 022

**Regra:** antes de abrir uma dica, escreva no código (num comentário) o que você tentou e onde travou. Abra um nível de cada vez.

## 1. NewBook

<details><summary>Nível 1 — direção</summary>

Olhe a primeira linha da tabela de operações: "criar com valores".
</details>

<details><summary>Nível 2 — estratégia</summary>

Não há loop nem `if`: a função só monta um `Book` com os três valores recebidos e o devolve. O `Read` você nem precisa escrever, porque o valor zero de `bool` já é `false`.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
devolver Book{ Title: <o título recebido>, Author: <...>, Pages: <...> }
```
</details>

## 2. TotalPages

<details><summary>Nível 1 — direção</summary>

Padrão do cardápio: **acumular**. Só muda o que você soma.
</details>

<details><summary>Nível 2 — estratégia</summary>

Um total começando em 0; para cada livro da lista, some as **páginas** daquele livro (não o livro inteiro); no fim, devolva o total.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
total = 0
para cada livro b da lista:
    total += o campo de páginas de b
devolver total
```
</details>

## 3. CountByAuthor

<details><summary>Nível 1 — direção</summary>

Você já escreveu essa função duas vezes, com outros nomes: `WordCount` (020) e `CharFrequency` (021). O que muda é **o que vira a chave**.
</details>

<details><summary>Nível 2 — estratégia</summary>

Padrão **contar por chave**. Map vazio (com `{}`, não com `var`); para cada livro, some 1 na chave que é o **autor daquele livro**; devolva o map. A chave é o valor de um campo do livro.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
counts = map de string para int, vazio
para cada livro b da lista:
    counts[ <o autor de b> ] += 1
devolver counts
```
</details>

## 4. MarkRead

<details><summary>Nível 1 — direção</summary>

Releia o parágrafo "Struct é copiada quando passa para uma função", no "O que é". Isso que parece um problema é exatamente o que você precisa aqui.
</details>

<details><summary>Nível 2 — estratégia</summary>

O `b` que a função recebe já é uma cópia, só sua. Mexer nele não toca o original. Então: mude o campo que diz se foi lido e devolva essa cópia. Não precisa montar um `Book` novo campo por campo.
</details>

<details><summary>Nível 3 — quase a resposta</summary>

```
no b recebido, colocar o campo de lido em true
devolver b
```
</details>

## Erros comuns do compilador (e do teste)

| Mensagem | O que significa |
|---|---|
| `undefined: Book` | O tipo `Book` não está no seu arquivo, ou está com outro nome. Copie a definição do README. |
| `invalid operation: total += b (mismatched types int and Book)` | Você somou o livro inteiro. Some um **campo** dele (com ponto). |
| `b.pages undefined (type Book has no field or method pages, but does have field Pages)` | Campo com letra minúscula. Os nomes têm que ser exatamente os do tipo. |
| `counts.Author undefined (type map[string]int has no field or method Author)` | Ponto num map. Map é com colchete; o ponto é para o campo do livro. |
| `too few values in struct literal of type Book` | Criou o `Book` sem os nomes dos campos e faltou algum. Use `Campo: valor`. |
| `panic: assignment to entry in nil map` | O map foi criado com `var` (é `nil`). Use `map[string]int{}`. |
| `CountByAuthor(...) = map[Dune:1 Dune Messiah:1 Emma:1], want map[Austen:1 Herbert:2]` | A chave foi o título, não o autor. |
| `MarkRead(dune) = {... Read:false}, want {... Read:true}` | O livro devolvido não foi marcado: talvez você tenha devolvido um `Book` novo sem o `Read`, ou mudado o campo de outra variável. |
| `after MarkRead, the original dune changed` | Algo mudou o original. Com struct passada por valor, isso não deveria acontecer; confira se a função recebe `Book` (e não `*Book`). |
