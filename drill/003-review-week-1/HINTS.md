# Dicas — Drill 003

> **Regra:** antes de abrir uma dica, escreva no `party.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**.

## TotalSlices

<details>
<summary>Nível 1 — direção</summary>

Olhe o `TotalSeeds` no README. É o mesmo formato.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Declare a constante das 8 fatias fora da função, no topo do arquivo. A função só multiplica o parâmetro pela constante.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`const slicesPerPizza = 8` e `return pizzas * slicesPerPizza`.

</details>

## SlicesEach

<details>
<summary>Nível 1 — direção</summary>

Quantas fatias existem? Você já tem uma função que responde isso. Depois é só dividir.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Chame `TotalSlices(pizzas)` e divida o resultado por `people`. Os dois são `int`, então a divisão já descarta o que sobra. Aqui isso é exatamente o que você quer.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

Dá para chamar uma função dentro de uma conta: `TotalSlices(pizzas) / people`.

</details>

## CostEach

<details>
<summary>Nível 1 — direção</summary>

Quanto custam todas as pizzas juntas? Depois, divida entre as pessoas. O resultado tem centavos, então a conta tem que ser feita em `float64`.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

`pricePerPizza` já é `float64`, mas `pizzas` e `people` são `int`. Converta os dois antes de usar na conta.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`float64(pizzas) * pricePerPizza / float64(people)`

</details>

## Summary

<details>
<summary>Nível 1 — direção</summary>

O texto tem três números. Dois deles você já sabe calcular: são as partes 1 e 2. O terceiro é um dos parâmetros.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Use `fmt.Sprintf` com três `%d`, na ordem em que os números aparecem na frase. Passe as chamadas de função direto como argumentos.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`fmt.Sprintf("%d people, %d slices, %d slices each", people, TotalSlices(...), SlicesEach(...))`. Complete os parênteses com os parâmetros certos.

</details>

## Erros comuns do compilador

| Mensagem | O que significa |
|---|---|
| `invalid operation: ... (mismatched types float64 and int)` | Você misturou `float64` com `int` numa conta. Converta o `int` com `float64(...)`. |
| `cannot use ... (value of type float64) as int value in return statement` | A função promete `int`, mas a conta deu `float64`. No `SlicesEach`, a divisão deve ficar toda em `int`. |
| `Summary(2, 5) = "16 people, 5 slices, ..."` | Compilou, mas os números estão fora de ordem. Os argumentos do `Sprintf` entram na mesma ordem dos `%d` na frase. |
| `%!d(string=...)` dentro do texto | Você passou uma string onde o verbo `%d` esperava um inteiro (ou trocou a ordem dos argumentos). |
| `syntax error: unexpected :=, expected =` | Constante usa `=`, não `:=`. |
