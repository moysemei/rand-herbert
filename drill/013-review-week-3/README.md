# Drill 013 — Revisão da semana: três clássicos de entrevista

**Hoje:** sem conceito novo. A semana foi de loops (`for` clássico, `for` como while, condições combinadas e parar cedo), e hoje você junta tudo em três problemas que aparecem em entrevista. Dois deles pedem ajudantes que você reescreve de memória.

## O que é

Revisão do que você já sabe:

- **Cada variável do loop tem um papel:** o contador (o número da vez), o limite (só diz quando parar), o acumulador (nasce antes, cresce dentro) e, às vezes, um valor fixo. (drills 009 e 010)
- **Somar e contar são diferentes:** `total += x` ou `count++`. (drill 010)
- **Percorrer dígitos é "pega o último, tira o último":** `n % 10` e `n /= 10`, até sobrar 0. (drill 011)
- **Remontar um número é "abre espaço e encaixa":** `rev = rev*10 + digit`. (drill 011)
- **Padrão busca:** o "achei" pode sair de dentro do loop; o "não achei" vem depois dele. (drill 012)
- **Função com nome em minúscula** é ajudante interno do pacote. (drill 007)
- **Depois de um loop que encolhe o `n`, ele vale 0.** Se você precisar do original, guarde uma cópia antes. (drill 011)

## Por que é assim

Os três problemas de hoje têm versão no LeetCode: "Count Primes" (204), "Palindrome Number" (9) e "Add Digits" (258). Eles não caem em entrevista porque são difíceis, e sim porque mostram rápido se a pessoa sabe montar um loop, escolher o acumulador certo e reaproveitar uma função. É exatamente o que você treinou esta semana.

Reescrever o `isPrime` e o `sumDigits` **de memória**, sem abrir os drills 011 e 012, é parte do exercício: o que você consegue reescrever sem olhar é o que você realmente aprendeu.

## Como se usa

Exemplo de outro assunto, para lembrar a cópia do original e um loop com `if` dentro:

```go
package collatz

import "fmt"

// Describe returns how many steps n takes to reach 1 in the Collatz sequence:
// if n is even, halve it; if it is odd, triple it and add 1.
func Describe(n int) string {
	original := n // n will change inside the loop
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return fmt.Sprintf("%d takes %d steps", original, steps)
}
```

## Armadilhas

Os "pontos de atenção" que apareceram nos reviews:

- **Comparar com o `n` depois do loop.** Ele já vale 0. Compare com a cópia.
- **Contar quando era para somar (ou o contrário).** Antes de escrever, diga no plano se a função **conta** ou **soma**.
- **O "não achei" dentro do loop.** No `isPrime`, o `return true` vem depois do loop.
- **Os casos pequenos:** 0, 1, 2 e números de um dígito só. Teste cada um na cabeça.
- **Nomes que dizem o que guardam:** `count`, `rev`, `original`. Nada de `f`.
- **Comentários em todas as funções,** inclusive nos ajudantes: `returns` para número ou texto, `reports whether` para `bool`.

## O problema

Crie o arquivo **`puzzles.go`** nesta pasta, com **`package puzzles`**. Em todas as funções, `n` e `limit` são >= 0.

**Ajudantes, reescritos de memória (sem abrir os drills 011 e 012):**
- **`func isPrime(n int) bool`**: o número é primo?
- **`func sumDigits(n int) int`**: a soma dos dígitos.

**As três funções:**

1. **`func CountPrimes(limit int) int`**: quantos primos existem de 1 até `limit`, incluindo o `limit`. Use o `isPrime`.
   - `CountPrimes(10)` → `4` (2, 3, 5, 7)
   - `CountPrimes(2)` → `1`
   - `CountPrimes(1)` → `0`
   - `CountPrimes(30)` → `10`
2. **`func IsPalindromeNumber(n int) bool`**: o número se lê igual de trás para frente?
   - `IsPalindromeNumber(121)` → `true`
   - `IsPalindromeNumber(1221)` → `true`
   - `IsPalindromeNumber(7)` → `true`
   - `IsPalindromeNumber(123)` → `false`
   - `IsPalindromeNumber(10)` → `false`
   - **Sem strings**: só com `%`, `/`, `*` e `+`.
3. **`func DigitalRoot(n int) int`**: some os dígitos de `n`. Se o resultado tiver mais de um dígito, some os dígitos dele, e continue até sobrar um dígito só. Use o `sumDigits`.
   - `DigitalRoot(9875)` → `2` (9875 → 29 → 11 → 2)
   - `DigitalRoot(99)` → `9` (99 → 18 → 9)
   - `DigitalRoot(7)` → `7`
   - `DigitalRoot(0)` → `0`

O `IsPalindromeNumber` e o `DigitalRoot` são as partes de pensar.

**Antes de codar**, no `// Plano:`:
- **Uma linha por função:** conta ou soma? O que muda a cada volta? Quando para?
- **À mão:** o `IsPalindromeNumber(10)` (qual é o número invertido, e com o que você compara?) e o `DigitalRoot(9875)` (quais números aparecem, e o que faz o loop parar?).

Se travar, faça a tabela antes de abrir a dica. Depois as dicas, depois o chat. E anote nas "Dificuldades" o que travou e o que ajudou.

## Checklist

- [ ] `// Plano:` com uma linha por função e as duas contas à mão.
- [ ] Comentário do pacote e de **cada** função, inclusive dos dois ajudantes.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 013: week 3 review"
  git push
  ```
