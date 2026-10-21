# Drill 011 — for como while

**Hoje:** repetir **enquanto** uma condição for verdadeira, sem saber antes quantas voltas o loop vai dar. O caso de hoje é percorrer os dígitos de um número de qualquer tamanho.

## O que é

Um `for` só com a condição é o `while` dos outros idiomas:

```go
for condition {
	// runs again and again while condition is true
}
```

E um `for` sem nada é um loop infinito, do qual você sai com `break` (ou com `return`):

```go
for {
	// ...
	if done {
		break // leaves the loop right now
	}
}
```

O `continue` pula o resto da volta atual e vai direto para a próxima. Você vai usar `break` e `continue` de verdade no próximo tema (condições combinadas); hoje basta saber que existem.

**Percorrer os dígitos** é o "pega o último, tira o último" do review 005: `n % 10` pega o último dígito e `n / 10` tira ele. Com 3856:

| volta | `n` no começo | `n % 10` (o último) | `n / 10` (sem o último) |
|---|---|---|---|
| 1 | 3856 | 6 | 385 |
| 2 | 385 | 5 | 38 |
| 3 | 38 | 8 | 3 |
| 4 | 3 | 3 | 0 |
| — | 0 | o loop acaba | |

O loop é `for n > 0 { ... }`, e dentro dele o `n` encolhe: `n /= 10` (o mesmo que `n = n / 10`). Ninguém precisa saber antes quantos dígitos o número tem; o loop para sozinho quando sobra 0.

## Por que é assim

- **Uma palavra só.** O Python tem `while n > 0:`; o Go escreve `for n > 0 {`. É o mesmo `for` de sempre, com só a condição.
- **Qual `for` usar:** quando você sabe quantas voltas são (de 1 até `n`, `exp` vezes), use o `for` clássico. Quando depende do que acontece dentro (até o número acabar, até achar algo), use o `for` com condição.
- **Mudar o parâmetro é seguro.** Dentro da função, o `n` é uma cópia: fazer `n /= 10` não muda a variável de quem chamou. (Quando você **quiser** mudar a variável de fora, é com ponteiro, lá no tópico 18.)
- **Loop infinito sem querer** acontece quando nada dentro do loop leva a condição a ficar falsa. O `go test` fica parado. Aperte **Ctrl+C**. (Se você esperar, depois de 10 minutos ele desiste com `panic: test timed out`.)

## Como se usa

Exemplos de outro assunto (não são o problema de hoje):

```go
package savings

// MonthsToSave returns how many months it takes to reach goal
// saving perMonth each month.
func MonthsToSave(goal, perMonth int) int {
	saved := 0
	months := 0
	for saved < goal { // we don't know in advance how many rounds
		saved += perMonth
		months++
	}
	return months
}

// NextLeapYear returns the first leap year after year.
func NextLeapYear(year int) int {
	for {
		year++ // year is a copy: changing it is safe
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			break
		}
	}
	return year
}
```

## Armadilhas

- **Esquecer de mudar o `n`.** Se nada dentro do loop encolhe o `n`, o `n > 0` nunca fica falso. Resultado: loop infinito.
- **O caso do zero.** Com `n = 0`, o `for n > 0` não roda nenhuma vez. Às vezes isso já dá a resposta certa, às vezes não. Pense no zero separadamente.
- **A ordem dentro do loop.** Pegue o último dígito **antes** de tirar ele. Ao contrário, você pula o primeiro.
- **Depois do loop, o `n` vale 0.** Se você precisar do valor original, guarde uma cópia antes.
- **`break` sai do loop; `return` sai da função inteira.**

## O problema

Crie o arquivo **`digitloop.go`** nesta pasta, com **`package digitloop`**. Em todas as funções, `n` >= 0. Use o `for` com condição (o formato while).

1. **`func SumDigits(n int) int`**: a soma dos dígitos.
   - `SumDigits(472)` → `13` (4 + 7 + 2)
   - `SumDigits(1000)` → `1`
   - `SumDigits(9)` → `9`
   - `SumDigits(0)` → `0`
2. **`func CountDigits(n int) int`**: quantos dígitos o número tem.
   - `CountDigits(472)` → `3`
   - `CountDigits(1000)` → `4`
   - `CountDigits(7)` → `1`
   - `CountDigits(0)` → `1` (o zero tem um dígito)
3. **`func ReverseNumber(n int) int`**: o número com os dígitos invertidos. É o `ReverseThreeDigits` do drill 005 para qualquer tamanho, e por isso não dá para ter as variáveis `hundreds`, `tens` e `ones` fixas.
   - `ReverseNumber(472)` → `274`
   - `ReverseNumber(1200)` → `21`
   - `ReverseNumber(123456)` → `654321`
   - `ReverseNumber(5)` → `5`

O `CountDigits(0)` e o `ReverseNumber` são as partes de pensar.

**Antes de codar**, no `// Plano:`:
- **Uma linha por função:** o que muda a cada volta, o que fica fixo, o valor inicial e quando o loop para.
- **Uma tabela do `ReverseNumber(472)`,** com as colunas volta, `n`, último dígito e `rev`. A coluna `rev` tem que terminar em 274.

Se travar, faça a tabela antes de abrir a dica. Depois as dicas, depois o chat.

## Checklist

- [ ] `// Plano:` com uma linha por função e a tabela do `ReverseNumber(472)`.
- [ ] Comentário do pacote e de cada função: `// Name returns ...`, com espaço depois do `//`.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 011: for as while"
  git push
  ```
