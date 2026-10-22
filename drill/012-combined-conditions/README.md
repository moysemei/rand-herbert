# Drill 012 — Condições combinadas e parar o loop na hora certa

**Hoje:** juntar várias regras numa decisão só e sair do loop assim que a resposta aparecer.

## O que é

**Condições combinadas.** Quando duas regras podem valer ao mesmo tempo, o caso que junta as duas é o **mais específico** e vem primeiro. É a regra da escada de `if` do drill 006 e do `switch` do 007: a primeira condição verdadeira vence.

**Três jeitos de mudar o rumo de um loop:**

```go
for i := 1; i <= 10; i++ {
	if i%2 == 0 {
		continue // skip the rest of this round, go to the next i
	}
	if i > 7 {
		break // leave the loop now; the function goes on after the }
	}
	if i == 5 {
		return i // leave the whole function now
	}
}
```

**O padrão "busca".** Para responder "existe algum?", você percorre os candidatos e **para no primeiro** que serve. Se o loop acabar sem achar nenhum, a resposta é "não existe", e ela fica na linha **depois** do loop.

## Por que é assim

- **Parar cedo economiza trabalho.** Para provar que 1000 não é primo, basta achar o 2. Ninguém precisa testar os outros 997 candidatos. Em entrevista, "dá para parar antes?" é uma pergunta que sempre vale fazer.
- **Em Python existe `for ... else`** (o `else` roda se o loop não deu `break`). O Go não tem isso: o lugar do "não achei" é a linha logo depois do loop.
- **Um detalhe que vale ouro:** os divisores vêm em pares. Se 36 = 4 × 9, então quando você testa o 4, já "testou" o 9 também. Por isso, para saber se `n` é primo, basta testar os divisores `d` enquanto `d*d <= n`. Para 97, isso dá 9 testes em vez de 95. Não é obrigatório hoje, mas se você usar, explique no plano por que funciona.

## Como se usa

Exemplos de outro assunto (não são o problema de hoje):

```go
package search

// HasDigit reports whether digit appears in n (n >= 0).
func HasDigit(n, digit int) bool {
	if n == 0 {
		return digit == 0
	}
	for n > 0 {
		if n%10 == digit {
			return true // found it: no need to look at the other digits
		}
		n /= 10
	}
	return false // the loop ended without finding it
}

// SumSkipping returns 1 + 2 + ... + n, skipping the multiples of skip.
func SumSkipping(n, skip int) int {
	total := 0
	for i := 1; i <= n; i++ {
		if i%skip == 0 {
			continue
		}
		total += i
	}
	return total
}
```

## Armadilhas

- **Decidir "não achei" dentro do loop.** No padrão busca, o `return true` (achei) pode ficar dentro do loop, mas o `return false` (não achei) só pode vir **depois** dele. Um `else { return false }` dentro do loop desiste no primeiro candidato que não serve.
- **A ordem das condições.** Se você testa "múltiplo de 3" antes de "múltiplo de 3 e de 5", o 15 nunca chega no caso certo.
- **`string(n)` não é o `str(n)` do Python.** Em Go, `string(7)` não dá `"7"`: dá o caractere de código 7, que é invisível. Número vira texto com `strconv.Itoa(n)` (drill 002). O `go vet`, que roda junto com o `go test`, avisa.
- **Os casos pequenos.** O 0 e o 1 não são primos, e o 2 é. Confira o que o seu loop faz com cada um.
- **"Depois de n" não inclui o `n`.** Se a busca começa no próprio `n`, um `n` primo devolve ele mesmo.

## O problema

Crie o arquivo **`rules.go`** nesta pasta, com **`package rules`**, e implemente:

1. **`func FizzBuzz(n int) string`**: múltiplo de 3 e de 5 → `"FizzBuzz"`; só de 3 → `"Fizz"`; só de 5 → `"Buzz"`; os outros → o próprio número como texto.
   - `FizzBuzz(15)` → `"FizzBuzz"`
   - `FizzBuzz(9)` → `"Fizz"`
   - `FizzBuzz(10)` → `"Buzz"`
   - `FizzBuzz(7)` → `"7"`
2. **`func IsPrime(n int) bool`**: o número é primo? (Primo: maior que 1, e divisível só por 1 e por ele mesmo.)
   - `IsPrime(2)` → `true`
   - `IsPrime(17)` → `true`
   - `IsPrime(21)` → `false`
   - `IsPrime(1)` → `false`
   - `IsPrime(0)` → `false`
3. **`func NextPrime(n int) int`**: o menor primo **maior** que `n`. Use o seu `IsPrime`.
   - `NextPrime(10)` → `11`
   - `NextPrime(13)` → `17`
   - `NextPrime(1)` → `2`

O `IsPrime` é a parte de pensar.

**Antes de codar**, no `// Plano:`:
- **Uma linha por função:** quais são os candidatos, quando o loop para e o que a função devolve se ele acabar sem achar nada.
- **Na mão, `IsPrime(21)` e `IsPrime(7)`:** quais `d` você testa, e em qual deles a resposta aparece?

Se travar, faça a tabela antes de abrir a dica. Depois as dicas, depois o chat.

## Checklist

- [ ] `// Plano:` com uma linha por função e os testes do 21 e do 7 feitos à mão.
- [ ] Comentário do pacote e de cada função: `returns` ou, quando devolve `bool`, `reports whether`.
- [ ] Nesta pasta, `go test -v` mostra `PASS` nos três testes.
- [ ] Na raiz do repo:

  ```
  go fmt ./...
  git add .
  git commit -m "drill 012: combined conditions"
  git push
  ```
