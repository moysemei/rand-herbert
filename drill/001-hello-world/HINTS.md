# Dicas — Drill 001

> **Regra:** antes de abrir uma dica, escreva no `hello.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**.

## Hello e Greet

<details>
<summary>Nível 1 — direção</summary>

Olhe a função `Label` no exemplo do README. A estrutura é a mesma: recebe (ou não) alguma coisa, monta um texto e devolve.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

`Hello` não recebe nada, então os parênteses ficam vazios: `()`. `Greet` recebe `name string`. As duas devolvem `string`. Monte o texto juntando pedaços com `+`. Confira a vírgula, o espaço e o `!`.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`Greet` junta três pedaços: o texto `"Hello, "` (com o espaço no final), o `name` e o `"!"`.

</details>

## Shout

<details>
<summary>Nível 1 — direção</summary>

O pacote da biblioteca padrão que trabalha com texto se chama `strings`. Abra https://pkg.go.dev/strings e procure (Ctrl+F) por "upper".

</details>

<details>
<summary>Nível 2 — estratégia</summary>

São dois passos: chamar `Greet(name)` para montar o texto e passar esse resultado para a função que deixa tudo em maiúsculas. Não esqueça do `import "strings"` no topo do arquivo.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`strings.ToUpper(s)` devolve `s` em maiúsculas. Você pode passar a chamada de uma função direto como argumento de outra: `f(g(x))`.

</details>

## Erros comuns do compilador

| Mensagem | O que significa |
|---|---|
| `undefined: Hello` | A função ainda não existe, ou o nome está diferente. Maiúscula e minúscula importam: `hello` não é `Hello`. |
| `found packages main (hello.go) and hello (hello_test.go)` | Os dois arquivos precisam declarar o mesmo `package hello`. |
| `"strings" imported and not used` | Você importou o pacote, mas não usou (ou escreveu o nome errado na hora de usar). |
| `more than one character in rune literal` | Você usou aspas simples numa string. Troque por aspas duplas. |
| `syntax error: unexpected semicolon or newline before {` | A `{` tem que ficar na mesma linha do `func`. |
| `Greet("Gopher") = "Hello Gopher!", want "Hello, Gopher!"` | Compilou, mas o texto está diferente. Compare caractere por caractere o que veio (`got`) com o esperado (`want`). |
