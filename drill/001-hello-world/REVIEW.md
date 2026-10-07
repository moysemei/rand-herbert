# Review — Drill 001

## Veredito

Entregue: os três testes passaram. O `Shout` reaproveitando o `Greet`, que era o desafio, ficou certo de primeira.

## Processo

- **Você escreveu o plano antes de codar.** Foi por ele que deu pra ver onde o erro nasceu: ele dizia "imprimir".
- **Você aplicou a regra das dicas sozinho.** Quando travou, anotou no próprio arquivo o que não estava entendendo, em vez de largar ou pedir a resposta.
- **Você chegou na pergunta certa:** "se eu der `return "Hello, Gopher!"`, como faria pra Ada?". Essa pergunta é o conceito de parâmetro inteiro.
- **Você tirou o `import "fmt"`** assim que parou de usar o `Println`, antes que o compilador reclamasse.
- **Você leu o `got` e o `want` do teste** pra achar o `!` que faltava.

## Erros e melhorias

**[lógica] Imprimir no lugar de devolver.** A primeira versão tinha `return fmt.Println(...)`.
- *Por que você provavelmente pensou assim:* todo "Hello, World" que você já viu imprime na tela, e em Python o `print` aparece desde o primeiro dia.
- *Como pensar:* pergunte "quem vai usar esse valor?". Se for outro código (um teste, outra função, o `Shout`), a função **devolve**. Imprimir é só para mostrar a um humano, e em Go isso fica na borda do programa (o `main`), não nas funções que calculam.
- *O erro "too many return values":* `fmt.Println` devolve dois valores (`int` e `error`). O `return` tentou repassar os dois, mas a função tinha prometido uma `string`.

**[lógica] Valor fixo no lugar do parâmetro.** A segunda versão do `Greet` tinha `return "Gopher"`.
- *Como pensar:* antes de rodar os testes, execute a função na cabeça com **duas entradas diferentes** (`"Gopher"` e `"Ada"`). Se a resposta não mudar quando a entrada muda, você não usou o parâmetro.

**[idiomático] Comentário do pacote.** `// Package hello.` cala o aviso, mas a convenção é uma frase completa dizendo o que o pacote faz: `// Package hello builds greeting messages.`

**[idiomático] Funções exportadas sem comentário.** Em Go, toda função com letra maiúscula leva um comentário acima, começando pelo nome dela: `// Greet returns a greeting for name.` É isso que aparece na documentação (`go doc`, pkg.go.dev).

**[idiomático] Nome do arquivo.** `hello-world.go` compila, mas a convenção para nomes de arquivo é tudo minúsculo, sem hífen: `hello.go` ou `hello_world.go`. Em Go, o nome do arquivo às vezes tem significado (o sufixo `_test.go` é um exemplo), então vale seguir o padrão.

**[idiomático] Um import só.** Com um único pacote, o mais comum é `import "strings"`, sem parênteses. Com parênteses também funciona; é só estilo.

Sobre o plano: ele ainda diz "a segunda para **imprimir** outros nomes". Na próxima vez, descreva o que a função **devolve**.

## Solução de referência

```go
// Package hello builds greeting messages.
package hello

import "strings"

// Hello returns the classic greeting.
func Hello() string {
	return Greet("World") // Hello is just Greet with "World"
}

// Greet returns a greeting for name.
func Greet(name string) string {
	return "Hello, " + name + "!"
}

// Shout returns the greeting for name in upper case.
func Shout(name string) string {
	return strings.ToUpper(Greet(name))
}
```

A lógica do `Greet` e do `Shout` é idêntica à sua. A diferença está nos comentários e num detalhe do `Hello`: ele é só o `Greet` com `"World"`, então pode reaproveitar o `Greet` do mesmo jeito que você fez no `Shout`. Assim, se um dia o formato da saudação mudar, você muda em um lugar só.

Outra forma que você vai ver muito em código Go é `fmt.Sprintf("Hello, %s!", name)`. Ela **monta** a string e devolve, sem imprimir nada. Para juntar dois ou três pedaços, o `+` é simples e serve bem.

## Para levar

Funções devolvem valores, e quem chamou decide o que fazer com eles. O parâmetro é a lacuna que quem chama preenche: teste mentalmente com duas entradas diferentes.
