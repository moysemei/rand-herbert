# Drill 004 — Operadores

**Hoje:** fazer contas e perguntas com operadores: resto da divisão, comparações e lógica (`&&`, `||`, `!`).

## O que é

Três famílias de operadores:

| Família | Operadores | Resultado |
|---|---|---|
| Aritméticos | `+` `-` `*` `/` `%` | número |
| Comparação | `==` `!=` `<` `>` `<=` `>=` | `bool` |
| Lógicos | `&&` (e), `\|\|` (ou), `!` (não) | `bool` |

O **`%`** devolve o **resto** da divisão inteira: `17 % 5` é `2`, porque 17 = 5 × 3 + **2**.

Uma comparação já **é** um valor `bool`, então dá para devolver direto, sem `if`:

```go
func IsAdult(age int) bool {
	return age >= 18 // true or false
}
```

## Por que é assim

- **O `%` está por trás de muita lógica:** par ou ímpar (`n % 2`), último dígito (`n % 10`), "a cada 3 itens" (`i % 3 == 0`), relógios e calendários. Junto com a divisão inteira, ele quebra um número em partes: `17 / 5` diz quantas vezes o 5 cabe, `17 % 5` diz o que sobra.
- **Go usa símbolos, não palavras.** Python tem `and`, `or`, `not`; Go tem `&&`, `||`, `!`.
- **Só `bool` é verdadeiro ou falso.** Em Python, `0`, `""` e listas vazias contam como falso. Em Go, não: um `int` nunca vira `bool` sozinho. Você sempre compara explicitamente (`n % 2 == 0`, nunca só `n % 2`).
- **Curto-circuito:** em `a && b`, se `a` já é `false`, o `b` nem é avaliado. Em `a || b`, se `a` já é `true`, o mesmo.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package shop

const itemsPerBox = 12

// FullBoxes returns how many boxes can be completely filled.
func FullBoxes(items int) int {
	return items / itemsPerBox
}

// Leftover returns how many items don't fit in a full box.
func Leftover(items int) int {
	return items % itemsPerBox // 30 items -> 2 full boxes, 6 left over
}

// FreeShipping reports whether an order ships for free:
// orders of 100 or more, or any order from a member.
func FreeShipping(total float64, isMember bool) bool {
	return total >= 100 || isMember
}

// NeedsReview reports whether an order must be checked by hand:
// large orders from someone who is not a member.
func NeedsReview(total float64, isMember bool) bool {
	return total > 1000 && !isMember
}
```

## Armadilhas

- **`=` não é `==`.** `=` atribui um valor; `==` compara. `return n % 2 = 0` é erro de sintaxe.
- **`%` com número negativo segue o sinal da esquerda.** Em Go, `-3 % 2` é `-1`. Em Python é `1`. Quem testa "é ímpar?" com `n % 2 == 1` erra com negativos.
- **`&&` vem antes de `||`.** `a || b && c` é lido como `a || (b && c)`. Quando misturar os dois, use parênteses: além de garantir a ordem, deixa a intenção clara para quem lê.
- **`%` só funciona com inteiros.** Com `float64` dá erro de compilação.

## O problema

Crie o arquivo **`numbers.go`** nesta pasta, com **`package numbers`**, e implemente:

1. **`IsEven(n int) bool`**: o número é par?
   - `IsEven(4)` → `true`
   - `IsEven(7)` → `false`
   - `IsEven(-3)` → `false`
2. **`SecondsToClock(total int) string`**: transforma segundos em horas, minutos e segundos.
   - `SecondsToClock(59)` → `"0h 0m 59s"`
   - `SecondsToClock(3665)` → `"1h 1m 5s"` (3600 s = 1 h; sobram 65 s = 1 min e 5 s)
   - `SecondsToClock(86399)` → `"23h 59m 59s"`
   - Use constantes para os 60 segundos de um minuto e os segundos de uma hora.
3. **`IsLeapYear(year int) bool`**: o ano é bissexto? A regra:
   - é bissexto se for divisível por 4,
   - **exceto** se for divisível por 100,
   - **a não ser que** também seja divisível por 400.
   - `IsLeapYear(2024)` → `true`
   - `IsLeapYear(2023)` → `false`
   - `IsLeapYear(1900)` → `false` (divisível por 100, mas não por 400)
   - `IsLeapYear(2000)` → `true` (divisível por 400)
   - **Sem `if`**: monte a resposta só com comparações e operadores lógicos.

O `SecondsToClock` e o `IsLeapYear` são as partes de pensar. Antes de codar, faça na mão, no `// Plano:`, o caminho de 3665 segundos e o do ano 1900.

Se travar, as dicas estão em `HINTS.md`.

## Checklist

- [ ] `// Plano:` com o cálculo de 3665 segundos e a checagem do ano 1900 feitos à mão.
- [ ] Comentário do pacote e de cada função, em inglês, começando pelo nome.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  git add .
  git commit -m "drill 004: operators"
  git push
  ```
