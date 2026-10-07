# Dicas — Drill 002

> **Regra:** antes de abrir uma dica, escreva no `convert.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**.

## CelsiusToFahrenheit

<details>
<summary>Nível 1 — direção</summary>

A fórmula já está no README. Tudo o que entra (`c`) e tudo o que sai é `float64`, então não precisa converter nada aqui. É um bom lugar para usar a sua constante.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Guarde o 32 numa constante com um nome que explique o que ele é. Depois, monte a conta com `c`, os números da fórmula e a constante, e devolva o resultado.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

Em Go, a fórmula se escreve quase igual ao papel: `c*9/5 + nomeDaConstante`. Uma constante sem tipo (`const x = 32`) funciona numa conta com `float64`.

</details>

## MinutesToHours

<details>
<summary>Nível 1 — direção</summary>

Uma hora tem 60 minutos. Releia a primeira armadilha do README: o que acontece quando você divide um `int` por outro `int`?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Antes de dividir, transforme `minutes` em `float64`. Repare que a ordem importa: converter **depois** de dividir não resolve, porque as casas decimais já foram perdidas na divisão.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

Olhe o `AveragePrice` no README: é o mesmo movimento. `float64(minutes)` e só então a divisão por 60.

</details>

## Describe

<details>
<summary>Nível 1 — direção</summary>

`name` já é texto, mas `age` é `int`. Você não pode juntar `int` com `string` usando `+`. Precisa transformar o número em texto primeiro, e `string(age)` **não** faz isso: rode o teste com ele e leia a mensagem.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Existe um pacote da biblioteca padrão só para converter entre texto e outros tipos. O nome dele é a junção de "string" com "conversion". Procure em https://pkg.go.dev/std.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`strconv.Itoa(n)` transforma um `int` em `string` (Itoa = "integer to ASCII"). Outra opção é `fmt.Sprintf("%s is %d years old", name, age)`, que monta o texto inteiro de uma vez.

</details>

## Erros comuns do compilador

| Mensagem | O que significa |
|---|---|
| `cannot use minutes / 60 (value of type int) as float64 value in return statement` | A conta deu `int`, mas a função prometeu devolver `float64`. Converta. |
| `MinutesToHours(90) = 1, want 1.5` | Compilou, mas a divisão foi feita entre dois `int` e as casas decimais sumiram. Converta **antes** de dividir. |
| `invalid operation: ... (mismatched types float64 and int)` | Você misturou `float64` com uma **variável** `int` na mesma conta. Converta uma delas, ou use uma constante sem tipo. |
| `conversion from int to string yields a string of one rune, not a string of digits` | `string(36)` não vira `"36"`: vira o caractere de código 36 (`$`). Para número → texto, use o pacote certo (dica de nível 2). |
| `declared and not used: x` | Você criou uma variável e não usou. Use ou apague. |
| `no new variables on left side of :=` | A variável já existe. Para mudar o valor, use `=` em vez de `:=`. |
