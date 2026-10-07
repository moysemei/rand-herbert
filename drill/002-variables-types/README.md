# Drill 002 — Variáveis e tipos

**Hoje:** guardar valores em variáveis e converter entre tipos, que em Go é sempre explícito.

## O que é

Uma **variável** é um nome para um valor, e todo valor em Go tem um **tipo**. Os tipos básicos que você vai usar o tempo todo:

| Tipo | O que guarda | Exemplo |
|---|---|---|
| `int` | número inteiro | `42`, `-7` |
| `float64` | número com casas decimais | `3.14`, `0.5` |
| `string` | texto | `"Gopher"` |
| `bool` | verdadeiro ou falso | `true`, `false` |

Três formas de declarar:

```go
var age int = 30     // nome, tipo e valor
var count int        // sem valor: recebe o "zero value" do tipo (aqui, 0)
name := "Ada"        // forma curta: o Go deduz o tipo (string). Só dentro de funções.
```

E **constantes**, valores que nunca mudam:

```go
const daysInWeek = 7
```

## Por que é assim

- **Go nunca converte tipos sozinho.** Em Python, `1 + 1.5` funciona; em Go, somar um `int` com um `float64` é erro de compilação. Você converte na mão: `float64(x)`, `int(y)`. Parece chato, mas evita bugs silenciosos, como perder as casas decimais sem perceber.
- **Toda variável sempre tem um valor válido.** Se você não der um, ela recebe o *zero value*: `0` para números, `""` para string, `false` para bool. Não existe "undefined" ou `None` para esses tipos.
- **`:=` existe para deixar o código curto.** Dentro de funções, quase todo mundo usa `:=`. O `var` aparece quando você quer o zero value ou está fora de uma função.
- **Constante sem tipo se adapta.** `const daysInWeek = 7` funciona tanto numa conta com `int` quanto com `float64`. Uma variável com tipo (`var x int = 7`) não se adapta: ela é `int` e pronto.

## Como se usa

Exemplo de outro assunto (não é o problema de hoje):

```go
package shop

const taxRate = 0.1 // constant: never changes

// PriceWithTax returns the price plus 10% tax.
func PriceWithTax(price float64) float64 {
	tax := price * taxRate // := creates the variable and sets its value
	return price + tax
}

// AveragePrice returns the average price of count items.
func AveragePrice(total float64, count int) float64 {
	return total / float64(count) // count is int: convert before dividing
}
```

Repare que:

- `tax` só existe dentro de `PriceWithTax`. Variável criada dentro da função morre quando a função termina.
- Em `AveragePrice`, `total` é `float64` e `count` é `int`. Para dividir um pelo outro, os dois precisam ser do mesmo tipo, então `count` vira `float64(count)`.
- A conversão `float64(count)` não muda a variável `count`. Ela cria um **novo valor**, do outro tipo.

## Armadilhas

- **Divisão entre dois `int` dá `int`.** `7 / 2` é `3`, não `3.5`: a parte decimal é jogada fora. Para ter decimais, pelo menos um lado precisa ser `float64` **antes** da divisão.
- **`:=` só cria, não reatribui.** Na segunda vez, use `=`: `x := 1` e depois `x = 2`. Escrever `x := 2` de novo no mesmo lugar dá `no new variables on left side of :=`.
- **Variável declarada e não usada é erro**, igual import não usado: `declared and not used`.
- **Converter número em texto não é tão óbvio quanto parece.** Se o compilador reclamar de algo nessa linha, leia a mensagem com calma: ela dá a pista.

## O problema

Crie o arquivo **`convert.go`** nesta pasta, com **`package convert`**, e implemente:

1. **`CelsiusToFahrenheit(c float64) float64`**
   - Fórmula: F = C × 9 / 5 + 32
   - `CelsiusToFahrenheit(100)` → `212`
   - `CelsiusToFahrenheit(-40)` → `-40`
2. **`MinutesToHours(minutes int) float64`**
   - `MinutesToHours(90)` → `1.5`
   - `MinutesToHours(45)` → `0.75`
3. **`Describe(name string, age int) string`**
   - `Describe("Ada", 36)` → `"Ada is 36 years old"`
   - `Describe("Rob", 7)` → `"Rob is 7 years old"`

**Regras do dia** (o teste não confere, o review confere):

- Use pelo menos uma **constante** (`const`) e pelo menos uma variável com **`:=`**.
- Antes de rodar o teste, faça cada função na cabeça com **duas entradas diferentes**. Ficou de lição do drill 001.

Se travar, as dicas estão em `HINTS.md`.

## Checklist

- [ ] Logo abaixo do `package convert`, um comentário `// Plano:` dizendo o que cada função **devolve** e como.
- [ ] Um comentário acima de cada função, começando pelo nome dela (ex.: `// MinutesToHours converts ...`).
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  git add .
  git commit -m "drill 002: variables and types"
  git push
  ```
