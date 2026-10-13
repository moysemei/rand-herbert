# Drill 005 — Quebrando números em dígitos

**Hoje:** reforçar `/` e `%`, agora para separar os dígitos de um número e montar ele de volta. Não tem conceito novo: é o raciocínio do `SecondsToClock` aplicado em outra unidade.

## O que é

Um número decimal é uma soma de dígitos vezes posições:

```
3856 = 3 × 1000 + 8 × 100 + 5 × 10 + 6
```

Separar os dígitos é o mesmo problema do relógio: em vez de 3600 e 60, as "unidades" são 1000, 100 e 10. E as duas perguntas continuam as mesmas:

| Pergunta | Operador | Com 3856 |
|---|---|---|
| Quantos 1000 inteiros cabem? | `/` | `3856 / 1000` = `3` |
| O que sobra depois de tirar os milhares? | `%` | `3856 % 1000` = `856` |
| Qual é o último dígito? | `% 10` | `3856 % 10` = `6` |
| Como fica sem o último dígito? | `/ 10` | `3856 / 10` = `385` |

E a regra do review do 004 vale aqui: **primeiro calcular, depois formatar.** Cada pedaço vira uma variável, e só no fim o `Sprintf` monta o texto.

## Por que é assim

- **`/` e `%` são as duas metades da mesma divisão.** A divisão inteira diz quantas vezes cabe e o resto diz o que sobrou. Juntos, eles reconstroem o número: `n == n/d*d + n%d` vale sempre. Com 3856 e 1000: `3*1000 + 856`.
- **Em Python existe `divmod(3856, 1000)`, que devolve `(3, 856)`.** Em Go não existe: você escreve as duas contas. E onde o Python usa `//` para divisão inteira, em Go o `/` entre dois `int` já descarta a parte decimal.
- **Isso cai em entrevista.** "Reverse Integer", "Palindrome Number" e "Add Digits" são problemas clássicos (LeetCode 7, 9 e 258) feitos com `% 10` e `/ 10`. Hoje é com 3 dígitos fixos; quando chegar o `for`, vai ser com qualquer quantidade.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package measure

import "fmt"

const centimetersPerMeter = 100

// Split turns a length in centimeters into text like "2m 54cm".
func Split(total int) string {
	meters := total / centimetersPerMeter      // how many whole meters fit?
	centimeters := total % centimetersPerMeter // what is left?
	return fmt.Sprintf("%dm %dcm", meters, centimeters)
}

// Join does the opposite: 2 meters and 54 centimeters -> 254.
func Join(meters, centimeters int) int {
	return meters*centimetersPerMeter + centimeters
}
```

Repare: `Split` calcula duas variáveis e só depois chama o `Sprintf`; `Join` remonta o número multiplicando cada parte pela sua posição e somando.

## Armadilhas

- **Copiar a linha e esquecer de trocar.** Cada variável responde uma pergunta diferente. Se em todos os testes dois dígitos saem iguais, duas linhas fazem a mesma conta.
- **Confira o intervalo.** Um dígito fica sempre entre 0 e 9. Se aparecer `72` ou `47` no lugar de um dígito, faltou um passo.
- **`a % b / c` é `(a % b) / c`.** O `%` e o `/` têm a mesma prioridade e são lidos da esquerda para a direita. Na dúvida, use parênteses.
- **Zero à esquerda some.** Um `int` não guarda "021": ele é `21`.
- **Negativos mudam as contas.** Em Python, `-7 // 2` é `-4` (arredonda para baixo); em Go, `-7 / 2` é `-3` (corta a parte decimal). Hoje só tem números positivos.

## O problema

Crie o arquivo **`digits.go`** nesta pasta, com **`package digits`**, e implemente:

1. **`func LastDigit(n int) int`**: o último dígito de `n` (com `n` >= 0).
   - `LastDigit(1234)` → `4`
   - `LastDigit(7)` → `7`
   - `LastDigit(90)` → `0`
2. **`func DescribeDigits(n int) string`**: para `n` de 100 a 999, diz quantas centenas, dezenas e unidades ele tem.
   - `DescribeDigits(472)` → `"hundreds: 4, tens: 7, ones: 2"`
   - `DescribeDigits(905)` → `"hundreds: 9, tens: 0, ones: 5"`
   - `DescribeDigits(100)` → `"hundreds: 1, tens: 0, ones: 0"`
3. **`func ReverseThreeDigits(n int) int`**: para `n` de 100 a 999, devolve o número com os dígitos invertidos.
   - `ReverseThreeDigits(472)` → `274`
   - `ReverseThreeDigits(905)` → `509`
   - `ReverseThreeDigits(120)` → `21` (o "021" vira 21)
   - `ReverseThreeDigits(111)` → `111`
   - **Sem strings**: só com `/`, `%`, `*` e `+`.

Hoje o `10` e o `100` podem ficar direto no código: são as casas decimais e se explicam sozinhos. Constante vale a pena quando o número precisa de um nome para fazer sentido, como o 3600 do drill passado.

O `ReverseThreeDigits` vai repetir as contas do `DescribeDigits`. Por hoje, tudo bem: no tópico de funções você aprende a devolver três valores de uma vez.

**Antes de codar**, faça na mão no `// Plano:`:
- 472: a conta de cada dígito (centenas, dezenas, unidades), com o operador que usou.
- Como montar 274 a partir desses três dígitos.

Se travar, as dicas estão em `HINTS.md`.

## Checklist

- [ ] `// Plano:` com o 472 desmontado e o 274 montado, à mão.
- [ ] Comentário do pacote (`// Package digits ...`) e de cada função, começando pelo nome.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  git add .
  git commit -m "drill 005: digits"
  git push
  ```
