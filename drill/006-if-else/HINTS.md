# Dicas — Drill 006

> **Regra:** antes de abrir uma dica, escreva no `decide.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**.

## Max

<details>
<summary>Nível 1 — direção</summary>

Quantos caminhos existem? Qual pergunta separa um do outro?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Se `a` for maior que `b`, devolva `a`. Senão, devolva `b`. No empate, tanto faz qual você devolve: os dois são o mesmo número.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

Um `if a > b` com `return a` dentro, e depois do `if`, `return b`. Sem `else`: é o retorno antecipado do README.

</details>

## Classify

<details>
<summary>Nível 1 — direção</summary>

São quatro casos. Em que ordem você testa? Faça com -4: se a primeira pergunta for "é par?", o que a função devolve?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Primeiro "é negativo?", depois "é zero?", depois "é par?". Se nenhuma for verdadeira, sobra só um caso possível, e o último `return` não precisa de condição. "É par" você fez no drill 004.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```
if n < 0     -> return "negative"
if n == 0    -> return "zero"
if n%2 == 0  -> return "positive even"
return "positive odd"
```

</details>

## RoundToTen

<details>
<summary>Nível 1 — direção</summary>

Qual dígito decide se o número sobe ou desce? Você aprendeu a pegar esse dígito no drill 005.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

- Pegue o último dígito.
- O múltiplo de 10 de baixo é o número menos o último dígito: 47 - 7 = 40.
- Se o último dígito for 5 ou mais, o resultado é esse múltiplo mais 10. Senão, é o próprio múltiplo.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

```go
last := n % 10
if last >= 5 {
	return n - last + 10
}
return n - last
```

Se você usar `if last := n % 10; last >= 5 {`, o `last` deixa de existir depois do `}`, e o `return n - last` de baixo não compila. Por isso aqui a variável vem antes do `if`.

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `syntax error: unexpected keyword else, expected }` | O `else` foi para a linha de baixo do `}`. Tem que ser `} else {`, na mesma linha. |
| `missing return` | Algum caminho da função chega no fim sem `return`. Confira o que acontece quando nenhum `if` é verdadeiro. |
| `undefined: last` | Você está usando uma variável fora do bloco onde ela foi criada (por exemplo, criada no `if last := ...;` e usada depois do `}`). |
| `non-boolean condition in if statement` | A condição não é `bool`. Compare com `==`, `<`, `>`... |
| `Classify(-4) = "positive even", want "negative"` | O teste de "par" veio antes do teste de "negativo". A ordem importa. |
| `Classify(0) = "positive even", want "zero"` | Zero é par: o teste do zero precisa vir antes do de "par". |
| `RoundToTen(47) = 40, want 50` | Você só arredondou para baixo. Falta o caso de subir. |
| `RoundToTen(45) = 40, want 50` | Terminado em 5 sobe: use `>=`, não `>`. |
