# Dicas — Drill 005

> **Regra:** antes de abrir uma dica, escreva no `digits.go`, num comentário, o que você tentou e onde acha que está o problema. Depois abra **só o próximo nível**.

## LastDigit

<details>
<summary>Nível 1 — direção</summary>

Qual operador responde "o que sobra"? E sobra da divisão por quanto?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

O último dígito é o resto da divisão por 10: tudo que é múltiplo de 10 vai embora na divisão e só a unidade sobra.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`n % 10`

</details>

## DescribeDigits

<details>
<summary>Nível 1 — direção</summary>

Faça na mão com 472, como você fez com o 3665. Centenas: quantos 100 cabem? Unidades: o que sobra da divisão por 10? Dezenas: são os "minutos" do relógio, o pedaço do meio.

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Três variáveis, uma pergunta cada:

- **Centenas:** divisão inteira por 100.
- **Unidades:** resto da divisão por 10.
- **Dezenas:** primeiro tire as centenas (resto da divisão por 100, que dá 72), depois veja quantos 10 cabem no que sobrou.

Depois, um `fmt.Sprintf` com três `%d`, com as variáveis na mesma ordem dos slots.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`tens := n % 100 / 10`. Molde do texto: `"hundreds: %d, tens: %d, ones: %d"`.

</details>

## ReverseThreeDigits

<details>
<summary>Nível 1 — direção</summary>

472 = 4 × 100 + 7 × 10 + 2. Escreva o 274 do mesmo jeito. Qual dígito mudou de posição e qual ficou no lugar?

</details>

<details>
<summary>Nível 2 — estratégia</summary>

Separe os três dígitos com as mesmas contas do `DescribeDigits`. Depois multiplique cada dígito pela sua **nova** posição e some: as unidades vão para as centenas, as centenas vão para as unidades, e as dezenas ficam onde estão.

</details>

<details>
<summary>Nível 3 — quase a resposta</summary>

`ones*100 + tens*10 + hundreds`

</details>

## Erros comuns

| Mensagem | O que significa |
|---|---|
| `syntax error: non-declaration statement outside function body` | Faltou a palavra no começo da linha, quase sempre o `func`. No nvim, a mesma coisa aparece como `expected declaration, found ...`. |
| `declared and not used: tens` | Você criou a variável e não usou. Em Go isso é erro, não aviso. Confira se ela entrou no `Sprintf` ou na conta. |
| `"fmt" imported and not used` | Importou o `fmt` e não usou. Ele só é necessário no `DescribeDigits`. |
| `fmt.Sprintf call needs 2 args but has 3 args` | O número de `%d` no molde não bate com o número de valores depois dele. |
| `cannot use fmt.Sprintf(...) (value of type string) as int value in return statement` | No `ReverseThreeDigits` você montou um texto, mas a função devolve `int`. Monte o número com `*` e `+`. |
| `DescribeDigits(472) = "hundreds: 4, tens: 47, ones: 2"` | As dezenas foram calculadas sem tirar as centenas antes. |
| `DescribeDigits(472) = "hundreds: 4, tens: 72, ones: 2"` | Você tirou as centenas, mas não dividiu por 10 o que sobrou. |
| `ReverseThreeDigits(472) = 13` | Você somou os dígitos sem multiplicar cada um pela sua posição. |
